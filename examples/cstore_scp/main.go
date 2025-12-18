// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to receive DICOM files using C-STORE (SCP).
// This example shows how to:
// - Start a DICOM server
// - Accept C-STORE requests
// - Save received DICOM files to disk
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/writer"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/server"
	"github.com/cocosip/go-dicom/pkg/network/service"
)

var (
	// Command line flags
	port          = flag.Int("port", 11112, "Port to listen on")
	storageDir    = flag.String("storage", "./received_dicom", "Directory to store received DICOM files")
	maxConn       = flag.Int("max-conn", 10, "Maximum concurrent connections (0 = unlimited)")
	verbose       = flag.Bool("verbose", false, "Enable verbose logging")
	organizeByAE  = flag.Bool("organize-by-ae", false, "Organize files by calling AE title")
	organizeByDate = flag.Bool("organize-by-date", false, "Organize files by study date")
)

// Statistics tracks server metrics
type Statistics struct {
	mu                sync.Mutex
	totalReceived     int
	totalFailed       int
	bytesReceived     int64
	connectionsServed int
	startTime         time.Time
}

var stats = &Statistics{
	startTime: time.Now(),
}

func main() {
	flag.Parse()

	// Create storage directory if it doesn't exist
	if err := os.MkdirAll(*storageDir, 0755); err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}

	// Create DICOM server
	fmt.Printf("=== DICOM C-STORE SCP Example ===\n")
	fmt.Printf("Port:          %d\n", *port)
	fmt.Printf("Storage Dir:   %s\n", *storageDir)
	fmt.Printf("Max Conn:      %d\n", *maxConn)
	fmt.Printf("Verbose:       %v\n", *verbose)
	fmt.Println()

	srv := server.New(
		server.WithPort(*port),
		server.WithMaxConnections(*maxConn),
		server.WithAssociationTimeout(30*time.Second),
		server.WithRequestTimeout(60*time.Second),
	)

	// Set up C-ECHO handler (verification)
	srv.SetCEchoHandler(handleCEcho)

	// Set up C-STORE handler (storage)
	srv.SetCStoreHandler(handleCStore)

	// Set up association negotiator
	srv.SetAssociationNegotiatorFunc(handleAssociationNegotiation)

	// Set up connection lifecycle handlers
	srv.SetConnectionLifecycleHandlerFuncs(
		handleAbort,
		handleConnectionClosed,
	)

	// Start server in goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverErr := make(chan error, 1)
	go func() {
		fmt.Printf("Starting DICOM SCP server on port %d...\n", *port)
		fmt.Println("Press Ctrl+C to stop")
		fmt.Println()
		serverErr <- srv.ListenAndServe(ctx)
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		fmt.Printf("\nReceived signal: %v\n", sig)
		fmt.Println("Shutting down gracefully...")

		// Create shutdown context with timeout
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}

	case err := <-serverErr:
		if err != nil && err != context.Canceled {
			log.Fatalf("Server error: %v", err)
		}
	}

	// Print statistics
	printStatistics()
	fmt.Println("Server stopped.")
}

// handleCEcho handles C-ECHO verification requests
func handleCEcho(_ context.Context, req *dimse.CEchoRequest) (*dimse.CEchoResponse, error) {
	if *verbose {
		log.Printf("C-ECHO request received: %s", req)
	}
	return dimse.NewCEchoResponseFromRequest(req, 0x0000), nil
}

// handleCStore handles C-STORE storage requests
func handleCStore(_ context.Context, req *dimse.CStoreRequest) (*dimse.CStoreResponse, error) {
	// Get the dataset
	ds := req.DataDataset()
	if ds == nil {
		log.Println("ERROR: C-STORE request has no dataset")
		stats.incrementFailed()
		return dimse.NewCStoreResponseFromRequest(req, 0xC000), // Failed - unable to process
			fmt.Errorf("no dataset in C-STORE request")
	}

	// Extract metadata
	sopClassUID := req.AffectedSOPClassUID()
	sopInstanceUID := req.AffectedSOPInstanceUID()

	// Get additional metadata for logging
	patientName, _ := ds.GetString(tag.PatientName)
	studyDate, _ := ds.GetString(tag.StudyDate)
	modality, _ := ds.GetString(tag.Modality)

	log.Printf("C-STORE received:")
	log.Printf("  SOP Class:     %s", sopClassUID)
	log.Printf("  SOP Instance:  %s", sopInstanceUID)
	if patientName != "" {
		log.Printf("  Patient:       %s", patientName)
	}
	if studyDate != "" {
		log.Printf("  Study Date:    %s", studyDate)
	}
	if modality != "" {
		log.Printf("  Modality:      %s", modality)
	}

	// Determine file path
	filePath := generateFilePath(ds, sopInstanceUID)

	// Create directory if needed
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("  ERROR: Failed to create directory: %v", err)
		stats.incrementFailed()
		return dimse.NewCStoreResponseFromRequest(req, 0xC000), err
	}

	// Save the DICOM file
	if err := saveDICOMFile(filePath, ds); err != nil {
		log.Printf("  ERROR: Failed to save file: %v", err)
		stats.incrementFailed()
		return dimse.NewCStoreResponseFromRequest(req, 0xC000), err
	}

	// Get file size for statistics
	if info, err := os.Stat(filePath); err == nil {
		stats.addReceived(info.Size())
	}

	log.Printf("  Saved to:      %s", filePath)
	log.Println("  Status:        SUCCESS")
	log.Println()

	// Return success response
	return dimse.NewCStoreResponseFromRequest(req, 0x0000), nil
}

// handleAssociationNegotiation handles association negotiation
func handleAssociationNegotiation(ctx context.Context, assoc *association.Association, responder service.AssociationResponder) error {
	stats.incrementConnections()

	if *verbose {
		log.Printf("Association request from: %s (calling AE: %s)",
			assoc.CallingAE, assoc.CallingAE)
		log.Printf("Proposed presentation contexts: %d", len(assoc.PresentationContexts))
	}

	// Accept all presentation contexts with their first proposed transfer syntax
	// In a real application, you might want to validate:
	// - AE titles against a whitelist
	// - Specific SOP Classes
	// - Transfer syntaxes your implementation supports
	acceptedCount := 0
	for _, pc := range assoc.PresentationContexts {
		if len(pc.ProposedTransferSyntaxes) > 0 {
			// Accept with the first transfer syntax
			pc.Accept(pc.ProposedTransferSyntaxes[0])
			acceptedCount++

			if *verbose {
				log.Printf("  Accepted PC ID %d: %s with TS: %s",
					pc.ID, pc.AbstractSyntax, pc.ProposedTransferSyntaxes[0])
			}
		}
	}

	if acceptedCount == 0 {
		log.Println("WARNING: No presentation contexts accepted")
		return responder.SendReject(ctx, 1, 1, 3) // Result: rejected-permanent, Source: service-user, Reason: no-reason-given
	}

	log.Printf("Accepted %d/%d presentation contexts from %s\n",
		acceptedCount, len(assoc.PresentationContexts), assoc.CallingAE)

	// Send accept response
	return responder.SendAccept(ctx, assoc)
}

// handleAbort handles association abort events
func handleAbort(_ context.Context, source, reason byte) {
	log.Printf("Association aborted - Source: 0x%02X, Reason: 0x%02X", source, reason)
}

// handleConnectionClosed handles connection close events
func handleConnectionClosed(_ context.Context, err error) {
	if err != nil && *verbose {
		log.Printf("Connection closed with error: %v", err)
	} else if *verbose {
		log.Println("Connection closed normally")
	}
}

// generateFilePath generates a file path for storing the DICOM file
func generateFilePath(ds *dataset.Dataset, sopInstanceUID string) string {
	basePath := *storageDir

	// Organize by calling AE if requested
	if *organizeByAE {
		// Note: In a real handler, you'd get the calling AE from the association
		// For now, we'll use a default subdirectory
		basePath = filepath.Join(basePath, "default_ae")
	}

	// Organize by study date if requested
	if *organizeByDate {
		if studyDate, ok := ds.GetString(tag.StudyDate); ok && studyDate != "" {
			// Format: YYYYMMDD -> YYYY/MM/DD
			if len(studyDate) >= 8 {
				year := studyDate[0:4]
				month := studyDate[4:6]
				day := studyDate[6:8]
				basePath = filepath.Join(basePath, year, month, day)
			}
		}
	}

	// Use SOP Instance UID as filename
	filename := fmt.Sprintf("%s.dcm", sopInstanceUID)
	return filepath.Join(basePath, filename)
}

// saveDICOMFile saves a dataset to a DICOM file
func saveDICOMFile(filePath string, ds *dataset.Dataset) error {
	// Use Explicit VR Little Endian as default transfer syntax
	if err := writer.WriteFile(filePath, ds, writer.WithTransferSyntax(transfer.ExplicitVRLittleEndian)); err != nil {
		return fmt.Errorf("failed to write dataset: %w", err)
	}

	return nil
}

// Statistics methods
func (s *Statistics) addReceived(bytes int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalReceived++
	s.bytesReceived += bytes
}

func (s *Statistics) incrementFailed() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalFailed++
}

func (s *Statistics) incrementConnections() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.connectionsServed++
}

func printStatistics() {
	stats.mu.Lock()
	defer stats.mu.Unlock()

	uptime := time.Since(stats.startTime)

	fmt.Println("\n=== Server Statistics ===")
	fmt.Printf("Uptime:           %v\n", uptime.Round(time.Second))
	fmt.Printf("Connections:      %d\n", stats.connectionsServed)
	fmt.Printf("Files Received:   %d\n", stats.totalReceived)
	fmt.Printf("Files Failed:     %d\n", stats.totalFailed)
	fmt.Printf("Bytes Received:   %d (%.2f MB)\n",
		stats.bytesReceived,
		float64(stats.bytesReceived)/(1024*1024))

	if stats.totalReceived > 0 {
		avgSize := float64(stats.bytesReceived) / float64(stats.totalReceived)
		fmt.Printf("Avg File Size:    %.2f KB\n", avgSize/1024)
	}
}
