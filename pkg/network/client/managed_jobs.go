// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

var (
	managedVerificationSOPClassUID   = "1.2.840.10008.1.1"
	managedExplicitVRLittleEndianUID = "1.2.840.10008.1.2.1"
	managedImplicitVRLittleEndianUID = "1.2.840.10008.1.2"
)

type managedCEchoJob struct{ complete func(error) }

// NewCEchoJob creates a managed C-ECHO operation. complete may be nil.
func NewCEchoJob(complete func(error)) Job {
	return &managedCEchoJob{complete: complete}
}

func (*managedCEchoJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return []PresentationContextSpec{{
		AbstractSyntax:   managedVerificationSOPClassUID,
		TransferSyntaxes: []string{managedExplicitVRLittleEndianUID, managedImplicitVRLittleEndianUID},
	}}, nil
}

func (*managedCEchoJob) Execute(ctx context.Context, client *Client) error { return client.CEcho(ctx) }

func (j *managedCEchoJob) Complete(err error) {
	if j.complete != nil {
		j.complete(err)
	}
}

type managedCStoreJob struct {
	dataset  *dataset.Dataset
	contexts []PresentationContextSpec
	complete func(error)
}

// NewCStoreJob creates a managed C-STORE operation from an independent deep
// clone of source. The source transfer syntax is proposed first when known.
func NewCStoreJob(source *dataset.Dataset, complete func(error)) (Job, error) {
	if source == nil {
		return nil, fmt.Errorf("C-STORE dataset is nil")
	}
	cloned, err := source.DeepCloneChecked()
	if err != nil {
		return nil, fmt.Errorf("clone C-STORE dataset: %w", err)
	}
	sopClassUID, ok := cloned.GetString(tag.SOPClassUID)
	if !ok || sopClassUID == "" {
		return nil, fmt.Errorf("C-STORE dataset has no SOP Class UID")
	}
	transferSyntaxes := []string{}
	if sourceSyntax := cloned.InternalTransferSyntax(); sourceSyntax != nil {
		transferSyntaxes = append(transferSyntaxes, sourceSyntax.UID().UID())
	}
	transferSyntaxes = mergeTransferSyntaxes(transferSyntaxes, []string{managedExplicitVRLittleEndianUID, managedImplicitVRLittleEndianUID})
	return &managedCStoreJob{
		dataset: cloned,
		contexts: []PresentationContextSpec{{
			AbstractSyntax:   sopClassUID,
			TransferSyntaxes: transferSyntaxes,
		}},
		complete: complete,
	}, nil
}

func (j *managedCStoreJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedCStoreJob) Execute(ctx context.Context, client *Client) error {
	return client.CStore(ctx, j.dataset)
}

func (j *managedCStoreJob) Complete(err error) {
	if j.complete != nil {
		j.complete(err)
	}
}

type managedCFindJob struct {
	level    dimse.QueryRetrieveLevel
	query    *dataset.Dataset
	contexts []PresentationContextSpec
	complete func([]*dataset.Dataset, error)
	results  []*dataset.Dataset
}

// NewCFindJob creates a managed C-FIND operation from an independent deep
// clone of query. The query/retrieve information model selects its SOP Class.
func NewCFindJob(
	level dimse.QueryRetrieveLevel,
	query *dataset.Dataset,
	complete func([]*dataset.Dataset, error),
) (Job, error) {
	cloned, err := cloneManagedDataset(query, "C-FIND query")
	if err != nil {
		return nil, err
	}
	req := dimse.NewCFindRequest(level, cloned)
	return &managedCFindJob{
		level: level, query: cloned, contexts: managedContexts(req.AffectedSOPClassUID()), complete: complete,
	}, nil
}

func (j *managedCFindJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedCFindJob) Execute(ctx context.Context, client *Client) error {
	results, err := client.CFind(ctx, j.level, j.query)
	j.results = results
	return err
}

func (j *managedCFindJob) Complete(err error) {
	if j.complete != nil {
		j.complete(j.results, err)
	}
}

type managedCMoveJob struct {
	level           dimse.QueryRetrieveLevel
	moveDestination string
	identifier      *dataset.Dataset
	callback        CMoveCallback
	contexts        []PresentationContextSpec
	complete        func(error)
}

// NewCMoveJob creates a managed C-MOVE operation from an independent deep
// clone of identifier.
func NewCMoveJob(
	level dimse.QueryRetrieveLevel,
	moveDestination string,
	identifier *dataset.Dataset,
	callback CMoveCallback,
	complete func(error),
) (Job, error) {
	if moveDestination == "" {
		return nil, fmt.Errorf("C-MOVE destination is empty")
	}
	cloned, err := cloneManagedDataset(identifier, "C-MOVE identifier")
	if err != nil {
		return nil, err
	}
	req := dimse.NewCMoveRequest(level, moveDestination, cloned)
	return &managedCMoveJob{
		level: level, moveDestination: moveDestination, identifier: cloned, callback: callback,
		contexts: managedContexts(req.AffectedSOPClassUID()), complete: complete,
	}, nil
}

func (j *managedCMoveJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedCMoveJob) Execute(ctx context.Context, client *Client) error {
	return client.CMove(ctx, j.level, j.moveDestination, j.identifier, j.callback)
}

func (j *managedCMoveJob) Complete(err error) {
	if j.complete != nil {
		j.complete(err)
	}
}

type managedCGetJob struct {
	level      dimse.QueryRetrieveLevel
	identifier *dataset.Dataset
	callback   CGetCallback
	contexts   []PresentationContextSpec
	complete   func(error)
}

// NewCGetJob creates a managed C-GET operation from an independent deep clone
// of identifier.
func NewCGetJob(
	level dimse.QueryRetrieveLevel,
	identifier *dataset.Dataset,
	callback CGetCallback,
	complete func(error),
) (Job, error) {
	cloned, err := cloneManagedDataset(identifier, "C-GET identifier")
	if err != nil {
		return nil, err
	}
	req := dimse.NewCGetRequest(level, cloned)
	return &managedCGetJob{
		level: level, identifier: cloned, callback: callback, contexts: managedContexts(req.AffectedSOPClassUID()), complete: complete,
	}, nil
}

func (j *managedCGetJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedCGetJob) Execute(ctx context.Context, client *Client) error {
	return client.CGet(ctx, j.level, j.identifier, j.callback)
}

func (j *managedCGetJob) Complete(err error) {
	if j.complete != nil {
		j.complete(err)
	}
}

type managedNActionJob struct {
	requestedSOPClassUID    string
	requestedSOPInstanceUID string
	actionTypeID            uint16
	actionInformation       *dataset.Dataset
	contexts                []PresentationContextSpec
	complete                func(*dimse.NActionResponse, error)
	response                *dimse.NActionResponse
}

// NewNActionJob creates a managed N-ACTION operation. actionInformation is
// deep-cloned before the job is queued.
func NewNActionJob(
	requestedSOPClassUID string,
	requestedSOPInstanceUID string,
	actionTypeID uint16,
	actionInformation *dataset.Dataset,
	complete func(*dimse.NActionResponse, error),
) (Job, error) {
	if requestedSOPClassUID == "" {
		return nil, fmt.Errorf("N-ACTION SOP Class UID is empty")
	}
	cloned, err := cloneOptionalManagedDataset(actionInformation, "N-ACTION action information")
	if err != nil {
		return nil, err
	}
	return &managedNActionJob{
		requestedSOPClassUID: requestedSOPClassUID, requestedSOPInstanceUID: requestedSOPInstanceUID,
		actionTypeID: actionTypeID, actionInformation: cloned, contexts: managedContexts(requestedSOPClassUID), complete: complete,
	}, nil
}

func (j *managedNActionJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedNActionJob) Execute(ctx context.Context, client *Client) error {
	response, err := client.NAction(ctx, dimse.NewNActionRequest(
		j.requestedSOPClassUID, j.requestedSOPInstanceUID, j.actionTypeID, j.actionInformation,
	))
	j.response = response
	return err
}

func (j *managedNActionJob) Complete(err error) {
	if j.complete != nil {
		j.complete(j.response, err)
	}
}

type managedNCreateJob struct {
	affectedSOPClassUID    string
	affectedSOPInstanceUID string
	attributeList          *dataset.Dataset
	contexts               []PresentationContextSpec
	complete               func(*dimse.NCreateResponse, error)
	response               *dimse.NCreateResponse
}

// NewNCreateJob creates a managed N-CREATE operation. attributeList is deep
// cloned before the job is queued.
func NewNCreateJob(
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	attributeList *dataset.Dataset,
	complete func(*dimse.NCreateResponse, error),
) (Job, error) {
	if err := validateManagedSOPClass("N-CREATE", affectedSOPClassUID); err != nil {
		return nil, err
	}
	cloned, err := cloneOptionalManagedDataset(attributeList, "N-CREATE attribute list")
	if err != nil {
		return nil, err
	}
	return &managedNCreateJob{
		affectedSOPClassUID: affectedSOPClassUID, affectedSOPInstanceUID: affectedSOPInstanceUID,
		attributeList: cloned, contexts: managedContexts(affectedSOPClassUID), complete: complete,
	}, nil
}

func (j *managedNCreateJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedNCreateJob) Execute(ctx context.Context, client *Client) error {
	response, err := client.NCreate(ctx, dimse.NewNCreateRequest(
		j.affectedSOPClassUID, j.affectedSOPInstanceUID, j.attributeList,
	))
	j.response = response
	return err
}

func (j *managedNCreateJob) Complete(err error) {
	if j.complete != nil {
		j.complete(j.response, err)
	}
}

type managedNGetJob struct {
	requestedSOPClassUID    string
	requestedSOPInstanceUID string
	attributeIdentifierList []*tag.Tag
	contexts                []PresentationContextSpec
	complete                func(*dimse.NGetResponse, error)
	response                *dimse.NGetResponse
}

// NewNGetJob creates a managed N-GET operation from copied attribute tags.
func NewNGetJob(
	requestedSOPClassUID string,
	requestedSOPInstanceUID string,
	attributeIdentifierList []*tag.Tag,
	complete func(*dimse.NGetResponse, error),
) (Job, error) {
	if err := validateManagedSOPClass("N-GET", requestedSOPClassUID); err != nil {
		return nil, err
	}
	cloned, err := cloneManagedTags(attributeIdentifierList)
	if err != nil {
		return nil, err
	}
	return &managedNGetJob{
		requestedSOPClassUID: requestedSOPClassUID, requestedSOPInstanceUID: requestedSOPInstanceUID,
		attributeIdentifierList: cloned, contexts: managedContexts(requestedSOPClassUID), complete: complete,
	}, nil
}

func (j *managedNGetJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedNGetJob) Execute(ctx context.Context, client *Client) error {
	response, err := client.NGet(ctx, dimse.NewNGetRequest(
		j.requestedSOPClassUID, j.requestedSOPInstanceUID, j.attributeIdentifierList,
	))
	j.response = response
	return err
}

func (j *managedNGetJob) Complete(err error) {
	if j.complete != nil {
		j.complete(j.response, err)
	}
}

type managedNSetJob struct {
	requestedSOPClassUID    string
	requestedSOPInstanceUID string
	modificationList        *dataset.Dataset
	contexts                []PresentationContextSpec
	complete                func(*dimse.NSetResponse, error)
	response                *dimse.NSetResponse
}

// NewNSetJob creates a managed N-SET operation. modificationList is deep
// cloned before the job is queued.
func NewNSetJob(
	requestedSOPClassUID string,
	requestedSOPInstanceUID string,
	modificationList *dataset.Dataset,
	complete func(*dimse.NSetResponse, error),
) (Job, error) {
	if err := validateManagedSOPClass("N-SET", requestedSOPClassUID); err != nil {
		return nil, err
	}
	cloned, err := cloneOptionalManagedDataset(modificationList, "N-SET modification list")
	if err != nil {
		return nil, err
	}
	return &managedNSetJob{
		requestedSOPClassUID: requestedSOPClassUID, requestedSOPInstanceUID: requestedSOPInstanceUID,
		modificationList: cloned, contexts: managedContexts(requestedSOPClassUID), complete: complete,
	}, nil
}

func (j *managedNSetJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedNSetJob) Execute(ctx context.Context, client *Client) error {
	response, err := client.NSet(ctx, dimse.NewNSetRequest(
		j.requestedSOPClassUID, j.requestedSOPInstanceUID, j.modificationList,
	))
	j.response = response
	return err
}

func (j *managedNSetJob) Complete(err error) {
	if j.complete != nil {
		j.complete(j.response, err)
	}
}

type managedNDeleteJob struct {
	requestedSOPClassUID    string
	requestedSOPInstanceUID string
	contexts                []PresentationContextSpec
	complete                func(*dimse.NDeleteResponse, error)
	response                *dimse.NDeleteResponse
}

// NewNDeleteJob creates a managed N-DELETE operation.
func NewNDeleteJob(
	requestedSOPClassUID string,
	requestedSOPInstanceUID string,
	complete func(*dimse.NDeleteResponse, error),
) (Job, error) {
	if err := validateManagedSOPClass("N-DELETE", requestedSOPClassUID); err != nil {
		return nil, err
	}
	return &managedNDeleteJob{
		requestedSOPClassUID: requestedSOPClassUID, requestedSOPInstanceUID: requestedSOPInstanceUID,
		contexts: managedContexts(requestedSOPClassUID), complete: complete,
	}, nil
}

func (j *managedNDeleteJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedNDeleteJob) Execute(ctx context.Context, client *Client) error {
	response, err := client.NDelete(ctx, dimse.NewNDeleteRequest(j.requestedSOPClassUID, j.requestedSOPInstanceUID))
	j.response = response
	return err
}

func (j *managedNDeleteJob) Complete(err error) {
	if j.complete != nil {
		j.complete(j.response, err)
	}
}

type managedNEventReportJob struct {
	affectedSOPClassUID    string
	affectedSOPInstanceUID string
	eventTypeID            uint16
	eventInformation       *dataset.Dataset
	contexts               []PresentationContextSpec
	complete               func(*dimse.NEventReportResponse, error)
	response               *dimse.NEventReportResponse
}

// NewNEventReportJob creates a managed N-EVENT-REPORT operation.
// eventInformation is deep-cloned before the job is queued.
func NewNEventReportJob(
	affectedSOPClassUID string,
	affectedSOPInstanceUID string,
	eventTypeID uint16,
	eventInformation *dataset.Dataset,
	complete func(*dimse.NEventReportResponse, error),
) (Job, error) {
	if err := validateManagedSOPClass("N-EVENT-REPORT", affectedSOPClassUID); err != nil {
		return nil, err
	}
	cloned, err := cloneOptionalManagedDataset(eventInformation, "N-EVENT-REPORT event information")
	if err != nil {
		return nil, err
	}
	return &managedNEventReportJob{
		affectedSOPClassUID: affectedSOPClassUID, affectedSOPInstanceUID: affectedSOPInstanceUID,
		eventTypeID: eventTypeID, eventInformation: cloned, contexts: managedContexts(affectedSOPClassUID), complete: complete,
	}, nil
}

func (j *managedNEventReportJob) PresentationContexts() ([]PresentationContextSpec, error) {
	return clonePresentationContextSpecs(j.contexts)
}

func (j *managedNEventReportJob) Execute(ctx context.Context, client *Client) error {
	response, err := client.NEventReport(ctx, dimse.NewNEventReportRequest(
		j.affectedSOPClassUID, j.affectedSOPInstanceUID, j.eventTypeID, j.eventInformation,
	))
	j.response = response
	return err
}

func (j *managedNEventReportJob) Complete(err error) {
	if j.complete != nil {
		j.complete(j.response, err)
	}
}

func managedContexts(abstractSyntax string) []PresentationContextSpec {
	return []PresentationContextSpec{{
		AbstractSyntax:   abstractSyntax,
		TransferSyntaxes: []string{managedExplicitVRLittleEndianUID, managedImplicitVRLittleEndianUID},
	}}
}

func cloneManagedDataset(source *dataset.Dataset, name string) (*dataset.Dataset, error) {
	if source == nil {
		return nil, fmt.Errorf("%s is nil", name)
	}
	return cloneOptionalManagedDataset(source, name)
}

func cloneOptionalManagedDataset(source *dataset.Dataset, name string) (*dataset.Dataset, error) {
	if source == nil {
		return nil, nil
	}
	cloned, err := source.DeepCloneChecked()
	if err != nil {
		return nil, fmt.Errorf("clone %s: %w", name, err)
	}
	return cloned, nil
}

func validateManagedSOPClass(operation, sopClassUID string) error {
	if sopClassUID == "" {
		return fmt.Errorf("%s SOP Class UID is empty", operation)
	}
	return nil
}

func cloneManagedTags(source []*tag.Tag) ([]*tag.Tag, error) {
	if source == nil {
		return nil, nil
	}
	cloned := make([]*tag.Tag, len(source))
	for index, sourceTag := range source {
		if sourceTag == nil {
			return nil, fmt.Errorf("N-GET attribute tag %d is nil", index)
		}
		copied := *sourceTag
		cloned[index] = &copied
	}
	return cloned, nil
}
