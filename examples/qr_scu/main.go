// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates a complete Query/Retrieve SCU workflow.
// It performs:
// 1. C-ECHO verification
// 2. C-FIND query to search for studies/series/images
// 3. C-MOVE or C-GET to retrieve DICOM instances
//
// For C-GET, C-STORE sub-operations are received back over the same association.
// No separate SCP port is required.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/client"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/server"
	"github.com/cocosip/go-dicom/pkg/network/service"
	"github.com/cocosip/go-dicom/pkg/network/status"
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
	storePort        = flag.Int("store-port", 11114, "Local Storage SCP port for receiving C-MOVE images")
	outputDir        = flag.String("output-dir", "./received_qr", "Directory to save received DICOM files")
	timeout          = flag.Duration("timeout", 180*time.Second, "Overall timeout")
	verbose          = flag.Bool("verbose", false, "Enable verbose logging")
)

const (
	retrieveModeGet  = "get"
	retrieveModeMove = "move"
)

var receivedCount atomic.Int32

func main() {
	flag.Parse()

	level, err := parseQueryLevel(*queryLevel)
	if err != nil {
		log.Fatal(err)
	}

	retrieve := normalizeRetrieveMode(*retrieveMode)
	if retrieve != "none" && retrieve != retrieveModeMove && retrieve != retrieveModeGet {
		log.Fatalf("invalid -retrieve value %q, expected none|move|get", *retrieveMode)
	}

	// Create output directory for C-GET / C-MOVE results
	if retrieve == retrieveModeGet || retrieve == retrieveModeMove {
		if err := os.MkdirAll(*outputDir, 0755); err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	printBanner(level, retrieve)

	// Step 1: C-ECHO
	fmt.Println("\n[Step 1/4] Performing C-ECHO verification...")
	if err := performCEcho(ctx); err != nil {
		log.Fatalf("C-ECHO failed: %v", err)
	}
	fmt.Println("✓ C-ECHO succeeded")

	// Step 2: C-FIND
	fmt.Printf("\n[Step 2/4] Performing C-FIND query (level=%s)...\n", level)
	results, err := performCFind(ctx, level)
	if err != nil {
		log.Fatalf("C-FIND failed: %v", err)
	}
	fmt.Printf("✓ C-FIND completed: found %d result(s)\n", len(results))
	printFindResults(level, results)

	if len(results) == 0 {
		fmt.Println("\nNo results found. Exiting.")
		return
	}

	if retrieve == "none" {
		fmt.Println("\nRetrieve mode is 'none'. Exiting.")
		return
	}

	// Step 3: Select target for retrieve
	targetStudyUID := strings.TrimSpace(*retrieveStudyUID)
	if targetStudyUID == "" {
		targetStudyUID, _ = results[0].GetString(tag.StudyInstanceUID)
	}
	if targetStudyUID == "" {
		log.Fatal("unable to determine StudyInstanceUID for retrieve")
	}

	// Step 4: Perform retrieve
	switch retrieve {
	case retrieveModeMove:
		fmt.Printf("\n[Step 3/4] Starting embedded Storage SCP on port %d...\n", *storePort)
		stopSCP, err := startEmbeddedStoreSCP(ctx)
		if err != nil {
			log.Fatalf("Failed to start embedded Storage SCP: %v", err)
		}
		fmt.Printf("  ✓ Storage SCP listening on port %d (AE: %s)\n", *storePort, *moveDestination)
		fmt.Printf("  Configure the remote QR SCP to route AE \"%s\" -> this host:%d\n", *moveDestination, *storePort)
		fmt.Printf("  Output Dir: %s\n", *outputDir)

		fmt.Printf("\n[Step 4/4] Performing C-MOVE...\n")
		fmt.Printf("  Target Study:     %s\n", targetStudyUID)
		fmt.Printf("  Move Destination: %s\n", *moveDestination)
		if err := performCMove(ctx, targetStudyUID); err != nil {
			stopSCP()
			log.Fatalf("C-MOVE failed: %v", err)
		}
		// Allow any in-flight C-STORE sub-operations to finish before stopping.
		time.Sleep(2 * time.Second)
		stopSCP()
		fmt.Println("✓ C-MOVE completed")
		fmt.Printf("✓ Received %d DICOM instance(s)\n", receivedCount.Load())
	case retrieveModeGet:
		fmt.Printf("\n[Step 3/4] Performing C-GET...\n")
		fmt.Printf("  Target Study: %s\n", targetStudyUID)
		fmt.Printf("  Output Dir:   %s\n", *outputDir)
		fmt.Println("  (C-STORE sub-operations received on the same association)")
		if err := performCGet(ctx, targetStudyUID); err != nil {
			log.Fatalf("C-GET failed: %v", err)
		}
		fmt.Println("✓ C-GET completed")
		fmt.Printf("\n✓ Received %d DICOM instance(s)\n", receivedCount.Load())
	}

	fmt.Println("\n=== Query/Retrieve workflow completed successfully ===")
}

func printBanner(level dimse.QueryRetrieveLevel, retrieve string) {
	fmt.Printf("=== DICOM Query/Retrieve SCU - Complete Workflow ===\n")
	fmt.Printf("Target:          %s:%d\n", *host, *port)
	fmt.Printf("Calling AE:      %s\n", *callingAE)
	fmt.Printf("Called AE:       %s\n", *calledAE)
	fmt.Printf("Query Level:     %s\n", level)
	fmt.Printf("Retrieve Mode:   %s\n", retrieve)
	switch retrieve {
	case retrieveModeMove:
		fmt.Printf("Move Dest AE:    %s\n", *moveDestination)
		fmt.Printf("Store Port:      %d\n", *storePort)
		fmt.Printf("Output Dir:      %s\n", *outputDir)
	case retrieveModeGet:
		fmt.Printf("Output Dir:      %s\n", *outputDir)
	}
	fmt.Println()
}

// startEmbeddedStoreSCP starts a local Storage SCP in the background that receives
// images pushed by the remote QR SCP during a C-MOVE operation. It returns a stop
// function that shuts the server down gracefully.
func startEmbeddedStoreSCP(parentCtx context.Context) (stop func(), err error) {
	srv := server.New(
		server.WithPort(*storePort),
		server.WithAssociationTimeout(30*time.Second),
		server.WithRequestTimeout(120*time.Second),
	)

	srv.SetCStoreHandler(handleCStoreSubOperation)
	srv.SetAssociationNegotiatorFunc(acceptAllContexts)

	ctx, cancel := context.WithCancel(parentCtx)

	ready := make(chan error, 1)
	go func() {
		// Signal that the goroutine has started; the OS bind happens inside
		// ListenAndServe, so we give it a short grace period.
		ready <- nil
		if serveErr := srv.ListenAndServe(ctx); serveErr != nil && ctx.Err() == nil {
			log.Printf("Storage SCP error: %v", serveErr)
		}
	}()

	if startErr := <-ready; startErr != nil {
		cancel()
		return nil, startErr
	}

	// Brief pause so the TCP listener is fully bound before the caller proceeds.
	time.Sleep(200 * time.Millisecond)

	stop = func() {
		cancel()
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		_ = srv.Shutdown(shutdownCtx)
	}

	return stop, nil
}

// acceptAllContexts is an association negotiator that accepts every proposed
// presentation context using its first offered transfer syntax.
func acceptAllContexts(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
	for _, pc := range assoc.PresentationContexts {
		if len(pc.ProposedTransferSyntaxes) > 0 {
			pc.Accept(pc.ProposedTransferSyntaxes[0])
		}
	}
	return responder.SendAccept(ctx, assoc)
}

func performCEcho(ctx context.Context) error {
	c := client.New(
		client.WithCallingAE(strings.TrimSpace(*callingAE)),
		client.WithCalledAE(strings.TrimSpace(*calledAE)),
		client.WithConnectTimeout(10*time.Second),
		client.WithRequestTimeout(*timeout),
	)
	addPresentationContexts(c)

	if err := c.Connect(ctx, *host, *port); err != nil {
		return fmt.Errorf("connect failed: %w", err)
	}
	defer func() { _ = c.Close() }()

	return c.CEcho(ctx)
}

func performCFind(ctx context.Context, level dimse.QueryRetrieveLevel) ([]*dataset.Dataset, error) {
	c := client.New(
		client.WithCallingAE(strings.TrimSpace(*callingAE)),
		client.WithCalledAE(strings.TrimSpace(*calledAE)),
		client.WithConnectTimeout(10*time.Second),
		client.WithRequestTimeout(*timeout),
	)
	addPresentationContexts(c)

	if err := c.Connect(ctx, *host, *port); err != nil {
		return nil, fmt.Errorf("connect failed: %w", err)
	}
	defer func() { _ = c.Close() }()

	query := buildQueryDataset(level)
	return c.CFind(ctx, level, query)
}

func performCMove(ctx context.Context, studyUID string) error {
	c := client.New(
		client.WithCallingAE(strings.TrimSpace(*callingAE)),
		client.WithCalledAE(strings.TrimSpace(*calledAE)),
		client.WithConnectTimeout(10*time.Second),
		client.WithRequestTimeout(*timeout),
	)
	addPresentationContexts(c)

	if err := c.Connect(ctx, *host, *port); err != nil {
		return fmt.Errorf("connect failed: %w", err)
	}
	defer func() { _ = c.Close() }()

	identifier := dataset.NewWithElements([]element.Element{
		element.NewString(tag.StudyInstanceUID, vr.UI, []string{studyUID}),
	})

	return c.CMove(ctx, dimse.QueryRetrieveLevelStudy, strings.TrimSpace(*moveDestination), identifier,
		func(remaining, completed, failed, warning uint16) bool {
			fmt.Printf("  Progress: remaining=%d completed=%d failed=%d warning=%d\n",
				remaining, completed, failed, warning)
			return true
		})
}

func performCGet(ctx context.Context, studyUID string) error {
	c := client.New(
		client.WithCallingAE(strings.TrimSpace(*callingAE)),
		client.WithCalledAE(strings.TrimSpace(*calledAE)),
		client.WithConnectTimeout(10*time.Second),
		client.WithRequestTimeout(*timeout),
		client.WithCStoreHandler(handleCStoreSubOperation),
	)
	addPresentationContexts(c)

	if err := c.Connect(ctx, *host, *port); err != nil {
		return fmt.Errorf("connect failed: %w", err)
	}
	defer func() { _ = c.Close() }()

	identifier := dataset.NewWithElements([]element.Element{
		element.NewString(tag.StudyInstanceUID, vr.UI, []string{studyUID}),
	})

	return c.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier,
		func(remaining, completed, failed, warning uint16) bool {
			fmt.Printf("  Progress: remaining=%d completed=%d failed=%d warning=%d\n",
				remaining, completed, failed, warning)
			return true
		})
}

// handleCStoreSubOperation handles a C-STORE request sent by the SCP back over the
// same association during a C-GET operation. It writes the received dataset to disk.
func handleCStoreSubOperation(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
	ds := req.DataDataset()
	if ds == nil {
		return dimse.NewCStoreResponseFromRequest(req, status.CStoreErrorCannotUnderstand), fmt.Errorf("no dataset in C-STORE request")
	}

	sopInstanceUID := ds.TryGetString(tag.SOPInstanceUID)
	if sopInstanceUID == "" {
		sopInstanceUID = fmt.Sprintf("unknown_%d", time.Now().UnixNano())
	}

	filename := fmt.Sprintf("%s.dcm", sopInstanceUID)
	filePath := filepath.Join(*outputDir, filename)

	if *verbose {
		log.Printf("Receiving C-STORE: SOP=%s", sopInstanceUID)
	}

	if err := writer.WriteFile(filePath, ds, writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		log.Printf("Failed to write file: %v", err)
		return dimse.NewCStoreResponseFromRequest(req, status.CStoreErrorCannotUnderstand), err
	}

	receivedCount.Add(1)
	fmt.Printf("  Saved: %s\n", filename)

	return dimse.NewCStoreResponseFromRequest(req, status.Success), nil
}

func addPresentationContexts(c *client.Client) {
	transferSyntaxes := []string{
		uid.ImplicitVRLittleEndian.UID(),
		uid.ExplicitVRLittleEndian.UID(),
	}

	// Verification
	c.AddPresentationContext(uid.Verification.UID(), transferSyntaxes...)

	// Query/Retrieve - Patient Root
	c.AddPresentationContext(uid.PatientRootQueryRetrieveInformationModelFind.UID(), transferSyntaxes...)
	c.AddPresentationContext(uid.PatientRootQueryRetrieveInformationModelMove.UID(), transferSyntaxes...)
	c.AddPresentationContext(uid.PatientRootQueryRetrieveInformationModelGet.UID(), transferSyntaxes...)

	// Query/Retrieve - Study Root
	c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelFind.UID(), transferSyntaxes...)
	c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelMove.UID(), transferSyntaxes...)
	c.AddPresentationContext(uid.StudyRootQueryRetrieveInformationModelGet.UID(), transferSyntaxes...)

	// Storage (for receiving C-STORE during C-GET)
	storageClasses := []string{
		uid.CTImageStorage.UID(),
		uid.MRImageStorage.UID(),
		uid.SecondaryCaptureImageStorage.UID(),
		uid.XRayAngiographicImageStorage.UID(),
		uid.UltrasoundImageStorage.UID(),
	}
	for _, sopClass := range storageClasses {
		c.AddPresentationContext(sopClass, transferSyntaxes...)
	}
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
			nonEmpty(patientNameValue),
			nonEmpty(patientIDValue),
			nonEmpty(studyDateValue),
			nonEmpty(modalityValue),
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

func nonEmpty(v string) string {
	if strings.TrimSpace(v) == "" {
		return "-"
	}
	return v
}
