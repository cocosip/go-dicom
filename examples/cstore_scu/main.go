// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

// Package main demonstrates how to send DICOM files using C-STORE (SCU).
// This example shows how to:
// - Connect to a DICOM server
// - Read DICOM files
// - Send files using C-STORE operation
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/network/client"
)

var (
	// Command line flags
	host         = flag.String("host", "localhost", "DICOM server hostname or IP")
	port         = flag.Int("port", 11112, "DICOM server port")
	callingAE    = flag.String("calling-ae", "STORE_SCU", "Calling AE Title")
	calledAE     = flag.String("called-ae", "STORE_SCP", "Called AE Title")
	dicomFile    = flag.String("file", "", "DICOM file to send (required)")
	dicomDir     = flag.String("dir", "", "Directory containing DICOM files to send")
	timeout      = flag.Duration("timeout", 30*time.Second, "Operation timeout")
	verifyOnly   = flag.Bool("verify", false, "Only verify connection (C-ECHO)")
	printMetadata = flag.Bool("metadata", false, "Print file metadata before sending")
)

func main() {
	flag.Parse()

	// Validate parameters
	if *dicomFile == "" && *dicomDir == "" {
		fmt.Println("Error: Either -file or -dir must be specified")
		flag.Usage()
		os.Exit(1)
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	// Create DICOM client
	fmt.Printf("=== DICOM C-STORE SCU Example ===\n")
	fmt.Printf("Calling AE: %s\n", *callingAE)
	fmt.Printf("Called AE:  %s\n", *calledAE)
	fmt.Printf("Server:     %s:%d\n", *host, *port)
	fmt.Println()

	// Initialize client with configuration
	c := client.New(
		client.WithCallingAE(*callingAE),
		client.WithCalledAE(*calledAE),
		client.WithConnectTimeout(10*time.Second),
		client.WithRequestTimeout(*timeout),
	)

	// Add presentation contexts for various storage SOP classes
	// Verification SOP Class (for C-ECHO)
	c.AddPresentationContext(
		uid.Verification.UID(),
		uid.ImplicitVRLittleEndian.UID(),
		uid.ExplicitVRLittleEndian.UID(),
	)

	// CT Image Storage
	c.AddPresentationContext(
		uid.CTImageStorage.UID(),
		uid.ImplicitVRLittleEndian.UID(),
		uid.ExplicitVRLittleEndian.UID(),
		uid.JPEGLosslessSV1.UID(),
	)

	// MR Image Storage
	c.AddPresentationContext(
		"1.2.840.10008.5.1.4.1.1.4", // MR Image Storage
		uid.ImplicitVRLittleEndian.UID(),
		uid.ExplicitVRLittleEndian.UID(),
	)

	// Secondary Capture Image Storage (commonly used)
	c.AddPresentationContext(
		"1.2.840.10008.5.1.4.1.1.7", // Secondary Capture Image Storage
		uid.ImplicitVRLittleEndian.UID(),
		uid.ExplicitVRLittleEndian.UID(),
	)

	// Connect to server
	fmt.Printf("Connecting to %s:%d...\n", *host, *port)
	if err := c.Connect(ctx, *host, *port); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func() {
		if err := c.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()

	fmt.Println("Connected successfully!")
	fmt.Printf("Association established with %d presentation contexts\n", len(c.GetAssociation().PresentationContexts))
	fmt.Println()

	// Perform C-ECHO to verify connection
	fmt.Println("Performing C-ECHO verification...")
	if err := c.CEcho(ctx); err != nil {
		log.Fatalf("C-ECHO failed: %v", err)
	}
	fmt.Println("C-ECHO successful!")
	fmt.Println()

	// If verify-only mode, exit here
	if *verifyOnly {
		fmt.Println("Verification complete. Exiting.")
		return
	}

	// Collect files to send
	var filesToSend []string
	if *dicomFile != "" {
		filesToSend = append(filesToSend, *dicomFile)
	}
	if *dicomDir != "" {
		files, err := findDICOMFiles(*dicomDir)
		if err != nil {
			log.Fatalf("Failed to scan directory: %v", err)
		}
		filesToSend = append(filesToSend, files...)
	}

	if len(filesToSend) == 0 {
		fmt.Println("No DICOM files to send.")
		return
	}

	// Send files
	fmt.Printf("Sending %d file(s)...\n\n", len(filesToSend))
	successCount := 0
	failureCount := 0

	for i, filePath := range filesToSend {
		fmt.Printf("[%d/%d] Processing: %s\n", i+1, len(filesToSend), filepath.Base(filePath))

		if err := sendDICOMFile(ctx, c, filePath); err != nil {
			fmt.Printf("  ERROR: %v\n", err)
			failureCount++
		} else {
			fmt.Println("  SUCCESS")
			successCount++
		}
		fmt.Println()
	}

	// Summary
	fmt.Println("=== Summary ===")
	fmt.Printf("Total files: %d\n", len(filesToSend))
	fmt.Printf("Successful:  %d\n", successCount)
	fmt.Printf("Failed:      %d\n", failureCount)
}

// sendDICOMFile reads and sends a single DICOM file
func sendDICOMFile(ctx context.Context, c *client.Client, filePath string) error {
	// Parse the DICOM file
	result, err := parser.ParseFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	// Print metadata if requested
	if *printMetadata {
		printFileMetadata(result, filePath)
	}

	// Verify required UIDs are present
	sopClassUID, ok := result.Dataset.GetString(tag.SOPClassUID)
	if !ok {
		return fmt.Errorf("missing SOPClassUID in dataset")
	}

	sopInstanceUID, ok := result.Dataset.GetString(tag.SOPInstanceUID)
	if !ok {
		return fmt.Errorf("missing SOPInstanceUID in dataset")
	}

	fmt.Printf("  SOP Class:    %s\n", sopClassUID)
	fmt.Printf("  SOP Instance: %s\n", sopInstanceUID)

	// Send using C-STORE
	startTime := time.Now()
	if err := c.CStore(ctx, result.Dataset); err != nil {
		return fmt.Errorf("C-STORE failed: %w", err)
	}
	duration := time.Since(startTime)

	fmt.Printf("  Sent in:      %v\n", duration)

	return nil
}

// findDICOMFiles recursively finds all files in a directory
// Returns all files (assumes they are DICOM files - validation happens during parse)
func findDICOMFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Add all files - we'll validate they're DICOM during parsing
		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// printFileMetadata prints basic metadata about a DICOM file
func printFileMetadata(result *parser.ParseResult, filePath string) {
	fmt.Println("  === File Metadata ===")
	fmt.Printf("  File:         %s\n", filepath.Base(filePath))
	fmt.Printf("  Transfer Syntax: %s\n", result.TransferSyntax.UID().Name())

	if patientName, ok := result.Dataset.GetString(tag.PatientName); ok {
		fmt.Printf("  Patient Name: %s\n", patientName)
	}
	if patientID, ok := result.Dataset.GetString(tag.PatientID); ok {
		fmt.Printf("  Patient ID:   %s\n", patientID)
	}
	if studyDesc, ok := result.Dataset.GetString(tag.StudyDescription); ok {
		fmt.Printf("  Study Desc:   %s\n", studyDesc)
	}
	if modality, ok := result.Dataset.GetString(tag.Modality); ok {
		fmt.Printf("  Modality:     %s\n", modality)
	}
	fmt.Println("  ====================")
}
