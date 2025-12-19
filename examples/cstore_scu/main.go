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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/parser"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/client"
)

// Configuration - modify these values as needed
const (
	host          = "localhost"
	port          = 11112
	callingAE     = "NETPUSH"
	calledAE      = "NETGATE"
	dicomFile     = ""      // Single DICOM file to send (leave empty to use dicomDir)
	dicomDir      = "D:\\1" // Directory containing DICOM files to send
	timeout       = 30 * time.Second
	verifyOnly    = false // Only verify connection (C-ECHO)
	printMetadata = false // Print file metadata before sending
)

func main() {
	// Validate parameters
	if dicomFile == "" && dicomDir == "" {
		fmt.Println("Error: Either dicomFile or dicomDir must be specified in the source code")
		fmt.Println("Press Enter to exit...")
		fmt.Scanln()
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create DICOM client
	fmt.Printf("=== DICOM C-STORE SCU Example ===\n")
	fmt.Printf("Calling AE: %s\n", callingAE)
	fmt.Printf("Called AE:  %s\n", calledAE)
	fmt.Printf("Server:     %s:%d\n", host, port)
	fmt.Println()

	// Initialize client with configuration
	c := client.New(
		client.WithCallingAE(callingAE),
		client.WithCalledAE(calledAE),
		client.WithConnectTimeout(10*time.Second),
		client.WithRequestTimeout(timeout),
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
	fmt.Printf("Connecting to %s:%d...\n", host, port)
	if err := c.Connect(ctx, host, port); err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		fmt.Println("Press Enter to exit...")
		fmt.Scanln()
		return
	}
	defer func() {
		if err := c.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()

	fmt.Println("Connected successfully!")
	assoc := c.GetAssociation()
	fmt.Printf("Association established with %d presentation contexts\n", len(assoc.PresentationContexts))

	// Print accepted presentation contexts for debugging
	fmt.Println("Accepted Presentation Contexts:")
	for _, pc := range assoc.PresentationContexts {
		if pc.Result == 0 && pc.AcceptedTransferSyntax != nil { // Accepted
			fmt.Printf("  ID %d: %s with %s\n", pc.ID, pc.AbstractSyntax, pc.AcceptedTransferSyntax.UID().Name())
		}
	}
	fmt.Println()

	// Perform C-ECHO to verify connection
	fmt.Println("Performing C-ECHO verification...")
	if err := c.CEcho(ctx); err != nil {
		fmt.Printf("C-ECHO failed: %v\n", err)
		fmt.Println("Press Enter to exit...")
		fmt.Scanln()
		return
	}
	fmt.Println("C-ECHO successful!")
	fmt.Println()

	// If verify-only mode, exit here
	if verifyOnly {
		fmt.Println("Verification complete. Exiting.")
		return
	}

	// Collect files to send
	var filesToSend []string
	if dicomFile != "" {
		filesToSend = append(filesToSend, dicomFile)
	}
	if dicomDir != "" {
		files, err := findDICOMFiles(dicomDir)
		if err != nil {
			fmt.Printf("Failed to scan directory: %v\n", err)
			fmt.Println("Press Enter to exit...")
			fmt.Scanln()
			return
		}
		filesToSend = append(filesToSend, files...)
	}

	if len(filesToSend) == 0 {
		fmt.Println("No DICOM files to send.")
		fmt.Println("Press Enter to exit...")
		fmt.Scanln()
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
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}

// sendDICOMFile reads and sends a single DICOM file
func sendDICOMFile(ctx context.Context, c *client.Client, filePath string) error {
	// Parse the DICOM file
	result, err := parser.ParseFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	// Print metadata if requested
	if printMetadata {
		printFileMetadata(result, filePath)
	}

	// Verify required UIDs are present and ensure they're in the dataset
	sopClassUID, ok := result.Dataset.GetString(tag.SOPClassUID)
	if !ok {
		// Try to get it from FileMetaInformation if available
		if result.FileMetaInformation != nil {
			if metaClass, metaOk := result.FileMetaInformation.MediaStorageSOPClassUID(); metaOk {
				sopClassUID = metaClass
				// Add it to the dataset for C-STORE
				elem := element.NewString(tag.SOPClassUID, vr.UI, []string{sopClassUID})
				if err := result.Dataset.AddOrUpdate(elem); err != nil {
					return fmt.Errorf("failed to add SOPClassUID to dataset: %w", err)
				}
				ok = true
			}
		}
		if !ok {
			return fmt.Errorf("missing SOPClassUID in dataset and FileMetaInformation")
		}
	}

	sopInstanceUID, ok := result.Dataset.GetString(tag.SOPInstanceUID)
	if !ok {
		// Try to get it from FileMetaInformation if available
		if result.FileMetaInformation != nil {
			if metaInstance, metaOk := result.FileMetaInformation.MediaStorageSOPInstanceUID(); metaOk {
				sopInstanceUID = metaInstance
				// Add it to the dataset for C-STORE
				elem := element.NewString(tag.SOPInstanceUID, vr.UI, []string{sopInstanceUID})
				if err := result.Dataset.AddOrUpdate(elem); err != nil {
					return fmt.Errorf("failed to add SOPInstanceUID to dataset: %w", err)
				}
				ok = true
			}
		}
		if !ok {
			return fmt.Errorf("missing SOPInstanceUID in dataset and FileMetaInformation")
		}
	}

	// Remove Group Length elements before sending (compatibility issue with some PACS)
	// Group Length elements have tag format (GGGG,0000)
	elementsToRemove := []*tag.Tag{}
	for _, elem := range result.Dataset.Elements() {
		if elem.Tag().Element() == 0x0000 {
			elementsToRemove = append(elementsToRemove, elem.Tag())
		}
	}
	for _, t := range elementsToRemove {
		result.Dataset.Remove(t)
	}

	// Send using C-STORE
	if err := c.CStore(ctx, result.Dataset); err != nil {
		return fmt.Errorf("C-STORE failed: %w", err)
	}

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
