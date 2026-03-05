// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates a minimal Query/Retrieve SCU.
// It performs C-ECHO + C-FIND, and optionally C-MOVE or C-GET.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/client"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

var (
	host             = flag.String("host", "localhost", "SCP host")
	port             = flag.Int("port", 11113, "SCP port")
	callingAE        = flag.String("calling-ae", "QRSCU", "Calling AE Title")
	calledAE         = flag.String("called-ae", "QRSCP", "Called AE Title")
	queryLevel       = flag.String("level", "STUDY", "Query level: PATIENT|STUDY|SERIES|IMAGE")
	patientName      = flag.String("patient-name", "", "Optional PatientName match key (supports wildcard, e.g. DOE^*)")
	patientID        = flag.String("patient-id", "", "Optional PatientID match key")
	studyDate        = flag.String("study-date", "", "Optional StudyDate (YYYYMMDD or range YYYYMMDD-YYYYMMDD)")
	studyUID         = flag.String("study-uid", "", "Optional StudyInstanceUID match key")
	retrieveMode     = flag.String("retrieve", "none", "Retrieve mode: none|move|get")
	retrieveStudyUID = flag.String("retrieve-study-uid", "", "StudyInstanceUID for retrieve (defaults to first C-FIND result)")
	moveDestination  = flag.String("move-destination", "QRDEST", "Move destination AE (used when -retrieve=move)")
	timeout          = flag.Duration("timeout", 30*time.Second, "Overall timeout")
)

func main() {
	flag.Parse()

	level, err := parseQueryLevel(*queryLevel)
	if err != nil {
		log.Fatal(err)
	}

	retrieve := normalizeRetrieveMode(*retrieveMode)
	if retrieve != "none" && retrieve != "move" && retrieve != "get" {
		log.Fatalf("invalid -retrieve value %q, expected none|move|get", *retrieveMode)
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	c := client.New(
		client.WithCallingAE(strings.TrimSpace(*callingAE)),
		client.WithCalledAE(strings.TrimSpace(*calledAE)),
		client.WithConnectTimeout(10*time.Second),
		client.WithRequestTimeout(*timeout),
	)

	addPresentationContexts(c)

	fmt.Printf("=== DICOM Query/Retrieve SCU Example ===\n")
	fmt.Printf("Target:      %s:%d\n", *host, *port)
	fmt.Printf("Calling AE:  %s\n", *callingAE)
	fmt.Printf("Called AE:   %s\n", *calledAE)
	fmt.Printf("Level:       %s\n", level)
	fmt.Printf("Retrieve:    %s\n", retrieve)
	fmt.Println()

	if err := c.Connect(ctx, *host, *port); err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	defer func() {
		if err := c.Close(); err != nil {
			log.Printf("close warning: %v", err)
		}
	}()

	if err := c.CEcho(ctx); err != nil {
		log.Fatalf("C-ECHO failed: %v", err)
	}
	fmt.Println("C-ECHO succeeded")

	query := buildQueryDataset(level)
	results, err := c.CFind(ctx, level, query)
	if err != nil {
		log.Fatalf("C-FIND failed: %v", err)
	}

	printFindResults(level, results)
	if len(results) == 0 || retrieve == "none" {
		return
	}

	targetStudyUID := strings.TrimSpace(*retrieveStudyUID)
	if targetStudyUID == "" {
		targetStudyUID, _ = results[0].GetString(tag.StudyInstanceUID)
	}
	if targetStudyUID == "" {
		log.Fatal("unable to determine StudyInstanceUID for retrieve")
	}

	identifier := dataset.NewWithElements([]element.Element{
		element.NewString(tag.StudyInstanceUID, vr.UI, []string{targetStudyUID}),
	})

	switch retrieve {
	case "move":
		fmt.Printf("\nStarting C-MOVE for StudyInstanceUID=%s destination=%s\n", targetStudyUID, *moveDestination)
		err = c.CMove(ctx, dimse.QueryRetrieveLevelStudy, strings.TrimSpace(*moveDestination), identifier,
			func(remaining, completed, failed, warning uint16) bool {
				fmt.Printf("  C-MOVE progress: remaining=%d completed=%d failed=%d warning=%d\n", remaining, completed, failed, warning)
				return true
			})
		if err != nil {
			log.Fatalf("C-MOVE failed: %v", err)
		}
		fmt.Println("C-MOVE completed")
	case "get":
		fmt.Printf("\nStarting C-GET for StudyInstanceUID=%s\n", targetStudyUID)
		err = c.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier,
			func(remaining, completed, failed, warning uint16) bool {
				fmt.Printf("  C-GET progress:  remaining=%d completed=%d failed=%d warning=%d\n", remaining, completed, failed, warning)
				return true
			})
		if err != nil {
			log.Fatalf("C-GET failed: %v", err)
		}
		fmt.Println("C-GET completed")
	}
}

func addPresentationContexts(c *client.Client) {
	transferSyntaxes := []string{
		uid.ImplicitVRLittleEndian.UID(),
		uid.ExplicitVRLittleEndian.UID(),
	}

	c.AddPresentationContext(uid.Verification.UID(), transferSyntaxes...)

	c.AddPresentationContext(uid.PatientRootQueryRetrieveInformationModelFind.UID(), transferSyntaxes...)
	c.AddPresentationContext(uid.PatientRootQueryRetrieveInformationModelMove.UID(), transferSyntaxes...)
	c.AddPresentationContext(uid.PatientRootQueryRetrieveInformationModelGet.UID(), transferSyntaxes...)

	c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelFind.UID(), transferSyntaxes...)
	c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelMove.UID(), transferSyntaxes...)
	c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelGet.UID(), transferSyntaxes...)
}

func buildQueryDataset(level dimse.QueryRetrieveLevel) *dataset.Dataset {
	ds := dataset.NewWithElements([]element.Element{
		element.NewString(tag.QueryRetrieveLevel, vr.CS, []string{string(level)}),
		element.NewString(tag.PatientName, vr.PN, []string{strings.TrimSpace(*patientName)}),
		element.NewString(tag.PatientID, vr.LO, []string{strings.TrimSpace(*patientID)}),
		element.NewString(tag.StudyDate, vr.DA, []string{strings.TrimSpace(*studyDate)}),
		element.NewString(tag.StudyInstanceUID, vr.UI, []string{strings.TrimSpace(*studyUID)}),
		element.NewString(tag.AccessionNumber, vr.SH, []string{""}),
		element.NewString(tag.Modality, vr.CS, []string{""}),
		element.NewString(tag.SeriesInstanceUID, vr.UI, []string{""}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{""}),
		element.NewString(tag.RetrieveAETitle, vr.AE, []string{""}),
	})

	return ds
}

func printFindResults(level dimse.QueryRetrieveLevel, results []*dataset.Dataset) {
	fmt.Printf("\nC-FIND completed: %d result(s) at %s level\n", len(results), level)
	for i, ds := range results {
		patientNameValue, _ := ds.GetString(tag.PatientName)
		patientIDValue, _ := ds.GetString(tag.PatientID)
		studyDateValue, _ := ds.GetString(tag.StudyDate)
		studyUIDValue, _ := ds.GetString(tag.StudyInstanceUID)
		seriesUIDValue, _ := ds.GetString(tag.SeriesInstanceUID)
		sopUIDValue, _ := ds.GetString(tag.SOPInstanceUID)
		modalityValue, _ := ds.GetString(tag.Modality)

		fmt.Printf("[%d] Patient=%s (%s) StudyDate=%s Modality=%s\n",
			i+1,
			nonEmpty(patientNameValue, "-"),
			nonEmpty(patientIDValue, "-"),
			nonEmpty(studyDateValue, "-"),
			nonEmpty(modalityValue, "-"),
		)
		if studyUIDValue != "" {
			fmt.Printf("    StudyInstanceUID:  %s\n", studyUIDValue)
		}
		if seriesUIDValue != "" {
			fmt.Printf("    SeriesInstanceUID: %s\n", seriesUIDValue)
		}
		if sopUIDValue != "" {
			fmt.Printf("    SOPInstanceUID:    %s\n", sopUIDValue)
		}
	}
}

func parseQueryLevel(raw string) (dimse.QueryRetrieveLevel, error) {
	switch strings.ToUpper(strings.TrimSpace(raw)) {
	case "PATIENT":
		return dimse.QueryRetrieveLevelPatient, nil
	case "STUDY":
		return dimse.QueryRetrieveLevelStudy, nil
	case "SERIES":
		return dimse.QueryRetrieveLevelSeries, nil
	case "IMAGE":
		return dimse.QueryRetrieveLevelImage, nil
	default:
		return "", fmt.Errorf("invalid -level %q, expected PATIENT|STUDY|SERIES|IMAGE", raw)
	}
}

func normalizeRetrieveMode(raw string) string {
	return strings.ToLower(strings.TrimSpace(raw))
}

func nonEmpty(v, fallback string) string {
	if strings.TrimSpace(v) == "" {
		return fallback
	}
	return v
}
