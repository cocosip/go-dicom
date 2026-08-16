// Copyright (c) 2025 go-dicom authors.
// Licensed under the Microsoft Public License (MS-PL).

package printing

import (
	"context"
	"fmt"
	"sort"

	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

const printActionTypeID uint16 = 1

// DIMSEService is the DIMSE-N request surface required by the print client.
type DIMSEService interface {
	SendNCreate(context.Context, *dimse.NCreateRequest) (*dimse.NCreateResponse, error)
	SendNSet(context.Context, *dimse.NSetRequest) (*dimse.NSetResponse, error)
	SendNAction(context.Context, *dimse.NActionRequest) (*dimse.NActionResponse, error)
	SendNDelete(context.Context, *dimse.NDeleteRequest) (*dimse.NDeleteResponse, error)
}

// Client executes the Basic Print Management DIMSE-N workflow.
type Client struct {
	service DIMSEService
}

// NewClient creates a Basic Print Management client over an established DIMSE service.
func NewClient(service DIMSEService) *Client {
	return &Client{service: service}
}

// Print creates the print hierarchy, sets each Image Box, and prints the Film Session.
// A failure stops the workflow at the first unsuccessful operation; no remote rollback is implied.
func (c *Client) Print(ctx context.Context, session *FilmSession) error {
	if c == nil || c.service == nil {
		return fmt.Errorf("printing: DIMSE service is required")
	}
	if ctx == nil {
		return fmt.Errorf("printing: context is required")
	}
	if err := validatePrintSession(session); err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	sessionDataset, err := session.ToDataset()
	if err != nil {
		return fmt.Errorf("printing: encode Film Session %q: %w", session.SOPInstanceUID, err)
	}
	sessionCreate := dimse.NewNCreateRequest(session.SOPClassUID, session.SOPInstanceUID, sessionDataset)
	sessionResponse, err := c.service.SendNCreate(ctx, sessionCreate)
	if err != nil {
		return fmt.Errorf("printing: N-CREATE Film Session %q: %w", session.SOPInstanceUID, err)
	}
	if err := requireSuccessfulResponse("N-CREATE Film Session", session.SOPInstanceUID, sessionResponse); err != nil {
		return err
	}

	for _, lut := range session.PresentationLUTs {
		if err := ctx.Err(); err != nil {
			return err
		}
		lutDataset, err := lut.ToDataset()
		if err != nil {
			return fmt.Errorf("printing: encode Presentation LUT %q: %w", lut.SOPInstanceUID, err)
		}
		request := dimse.NewNCreateRequest(presentationLUTSOPClassUID, lut.SOPInstanceUID, lutDataset)
		response, err := c.service.SendNCreate(ctx, request)
		if err != nil {
			return fmt.Errorf("printing: N-CREATE Presentation LUT %q: %w", lut.SOPInstanceUID, err)
		}
		if err := requireSuccessfulResponse("N-CREATE Presentation LUT", lut.SOPInstanceUID, response); err != nil {
			return err
		}
	}

	for _, filmBox := range session.BasicFilmBoxes {
		if err := ctx.Err(); err != nil {
			return err
		}
		filmBoxDataset, err := filmBox.ToDataset()
		if err != nil {
			return fmt.Errorf("printing: encode Film Box %q: %w", filmBox.SOPInstanceUID, err)
		}
		filmBoxDataset.Remove(tag.ReferencedImageBoxSequence)
		createRequest := dimse.NewNCreateRequest(basicFilmBoxSOPClassUID, filmBox.SOPInstanceUID, filmBoxDataset)
		createResponse, err := c.service.SendNCreate(ctx, createRequest)
		if err != nil {
			return fmt.Errorf("printing: N-CREATE Film Box %q: %w", filmBox.SOPInstanceUID, err)
		}
		if err := requireSuccessfulResponse("N-CREATE Film Box", filmBox.SOPInstanceUID, createResponse); err != nil {
			return err
		}
		remoteReferences, err := imageBoxReferences(createResponse, len(filmBox.BasicImageBoxes))
		if err != nil {
			return fmt.Errorf("printing: N-CREATE Film Box %q: %w", filmBox.SOPInstanceUID, err)
		}
		localImageBoxes := append([]*ImageBox(nil), filmBox.BasicImageBoxes...)
		sort.SliceStable(localImageBoxes, func(left, right int) bool {
			return localImageBoxes[left].ImageBoxPosition < localImageBoxes[right].ImageBoxPosition
		})

		for index, imageBox := range localImageBoxes {
			if err := ctx.Err(); err != nil {
				return err
			}
			remoteReference := remoteReferences[index]
			if remoteReference.SOPClassUID != imageBox.SOPClassUID {
				return fmt.Errorf("printing: referenced Image Box %d SOP Class UID %q does not match local Image Box %q",
					index, remoteReference.SOPClassUID, imageBox.SOPClassUID)
			}
			imageDataset, err := imageBox.ToDataset()
			if err != nil {
				return fmt.Errorf("printing: encode Image Box %q: %w", imageBox.SOPInstanceUID, err)
			}
			setRequest := dimse.NewNSetRequest(remoteReference.SOPClassUID, remoteReference.SOPInstanceUID, imageDataset)
			setResponse, err := c.service.SendNSet(ctx, setRequest)
			if err != nil {
				return fmt.Errorf("printing: N-SET Image Box %q: %w", remoteReference.SOPInstanceUID, err)
			}
			if err := requireSuccessfulResponse("N-SET Image Box", remoteReference.SOPInstanceUID, setResponse); err != nil {
				return err
			}
		}
	}

	if err := ctx.Err(); err != nil {
		return err
	}
	actionRequest := dimse.NewNActionRequest(session.SOPClassUID, session.SOPInstanceUID, printActionTypeID, nil)
	actionResponse, err := c.service.SendNAction(ctx, actionRequest)
	if err != nil {
		return fmt.Errorf("printing: N-ACTION Film Session %q: %w", session.SOPInstanceUID, err)
	}
	return requireSuccessfulResponse("N-ACTION Film Session", session.SOPInstanceUID, actionResponse)
}

func imageBoxReferences(response *dimse.NCreateResponse, expected int) ([]SOPReference, error) {
	if response == nil || response.DataDataset() == nil {
		return nil, fmt.Errorf("referenced Image Box Sequence is missing from the response Attribute List")
	}
	references, err := readReferenceSequence(response.DataDataset(), tag.ReferencedImageBoxSequence)
	if err != nil {
		return nil, fmt.Errorf("read referenced Image Box Sequence: %w", err)
	}
	if len(references) != expected {
		return nil, fmt.Errorf("referenced Image Box Sequence has %d items, want %d", len(references), expected)
	}
	seen := make(map[string]struct{}, len(references))
	for index, reference := range references {
		if !uid.IsValid(reference.SOPClassUID) || !uid.IsValid(reference.SOPInstanceUID) {
			return nil, fmt.Errorf("referenced Image Box item %d has an invalid SOP UID", index)
		}
		if _, exists := seen[reference.SOPInstanceUID]; exists {
			return nil, fmt.Errorf("referenced Image Box Sequence contains duplicate SOP Instance UID %q", reference.SOPInstanceUID)
		}
		seen[reference.SOPInstanceUID] = struct{}{}
	}
	return references, nil
}

func validatePrintSession(session *FilmSession) error {
	if session == nil {
		return fmt.Errorf("printing: FilmSession is required")
	}
	if !session.IsValid() {
		return fmt.Errorf("printing: FilmSession %q is not valid", session.SOPInstanceUID)
	}
	seenUIDs := make(map[string]string)
	registerUID := func(kind, value string) error {
		if value == "" {
			return fmt.Errorf("printing: %s has no SOP Instance UID", kind)
		}
		if previous, exists := seenUIDs[value]; exists {
			return fmt.Errorf("printing: %s and %s share SOP Instance UID %q", previous, kind, value)
		}
		seenUIDs[value] = kind
		return nil
	}
	if session.SOPClassUID == "" {
		return fmt.Errorf("printing: FilmSession has no SOP Class UID")
	}
	if err := registerUID("FilmSession", session.SOPInstanceUID); err != nil {
		return err
	}
	for index, lut := range session.PresentationLUTs {
		if lut == nil || !lut.IsValid() {
			return fmt.Errorf("printing: PresentationLUT at index %d is not valid", index)
		}
		if err := registerUID("PresentationLUT", lut.SOPInstanceUID); err != nil {
			return err
		}
	}
	for boxIndex, filmBox := range session.BasicFilmBoxes {
		if filmBox == nil {
			return fmt.Errorf("printing: FilmBox at index %d is nil", boxIndex)
		}
		if err := registerUID("FilmBox", filmBox.SOPInstanceUID); err != nil {
			return err
		}
		seenPositions := make(map[uint16]struct{}, len(filmBox.BasicImageBoxes))
		for imageIndex, imageBox := range filmBox.BasicImageBoxes {
			if imageBox == nil {
				return fmt.Errorf("printing: ImageBox at FilmBox %q index %d is nil", filmBox.SOPInstanceUID, imageIndex)
			}
			if int(imageBox.ImageBoxPosition) > len(filmBox.BasicImageBoxes) {
				return fmt.Errorf("printing: FilmBox %q has Image Box position %d outside 1..%d",
					filmBox.SOPInstanceUID, imageBox.ImageBoxPosition, len(filmBox.BasicImageBoxes))
			}
			if _, exists := seenPositions[imageBox.ImageBoxPosition]; exists {
				return fmt.Errorf("printing: FilmBox %q has duplicate Image Box position %d", filmBox.SOPInstanceUID, imageBox.ImageBoxPosition)
			}
			seenPositions[imageBox.ImageBoxPosition] = struct{}{}
			if err := registerUID("ImageBox", imageBox.SOPInstanceUID); err != nil {
				return err
			}
		}
	}
	return nil
}

func requireSuccessfulResponse(operation, sopInstanceUID string, response dimse.Response) error {
	if response == nil {
		return fmt.Errorf("printing: %s %q returned no response", operation, sopInstanceUID)
	}
	responseStatus := response.Status()
	if responseStatus.IsSuccess() || responseStatus.IsWarning() {
		return nil
	}
	return fmt.Errorf("printing: %s %q failed: %s", operation, sopInstanceUID, responseStatus)
}
