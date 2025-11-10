// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/client"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// ExampleClient_CEcho demonstrates how to verify connectivity with C-ECHO.
func ExampleClient_CEcho() {
	// Create a new client with custom configuration
	c := client.New(
		client.WithCallingAE("MY_SCU"),
		client.WithCalledAE("REMOTE_SCP"),
		client.WithConnectTimeout(10*time.Second),
	)

	// Add presentation context for Verification SOP Class
	c.AddPresentationContext(
		"1.2.840.10008.1.1",   // Verification SOP Class
		"1.2.840.10008.1.2",   // Implicit VR Little Endian
		"1.2.840.10008.1.2.1", // Explicit VR Little Endian
	)

	// Connect to the DICOM server
	ctx := context.Background()
	if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer c.Close()

	// Send C-ECHO to verify connection
	if err := c.CEcho(ctx); err != nil {
		log.Fatalf("C-ECHO failed: %v", err)
	}

	fmt.Println("Connection verified successfully")
}

// ExampleClient_CStore demonstrates how to store a DICOM dataset.
func ExampleClient_CStore() {
	c := client.New(
		client.WithCallingAE("MY_SCU"),
		client.WithCalledAE("PACS_SCP"),
	)

	// Add presentation context for CT Image Storage
	c.AddPresentationContext(
		"1.2.840.10008.5.1.4.1.1.2", // CT Image Storage
		"1.2.840.10008.1.2",         // Implicit VR Little Endian
		"1.2.840.10008.1.2.1",       // Explicit VR Little Endian
	)

	ctx := context.Background()
	if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer c.Close()

	// Create a dataset to store
	ds := dataset.New()
	ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
	ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5.6.7.8.9"}))
	ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))
	ds.Add(element.NewString(tag.PatientID, vr.LO, []string{"12345"}))

	// Store the dataset
	if err := c.CStore(ctx, ds); err != nil {
		log.Fatalf("C-STORE failed: %v", err)
	}

	fmt.Println("Dataset stored successfully")
}

// ExampleClient_CFind demonstrates how to query a DICOM server.
func ExampleClient_CFind() {
	c := client.New(
		client.WithCallingAE("MY_SCU"),
		client.WithCalledAE("PACS_SCP"),
	)

	// Add presentation context for Study Root Query/Retrieve
	c.AddPresentationContext(
		"1.2.840.10008.5.1.4.1.2.2.1", // Study Root Query/Retrieve - FIND
		"1.2.840.10008.1.2",           // Implicit VR Little Endian
	)

	ctx := context.Background()
	if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer c.Close()

	// Create query dataset
	query := dataset.New()
	query.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN*"}))       // Wildcard search
	query.Add(element.NewString(tag.StudyDate, vr.DA, []string{"20240101-20241231"})) // Date range
	query.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{""}))           // Return value

	// Perform C-FIND at Study level
	results, err := c.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)
	if err != nil {
		log.Fatalf("C-FIND failed: %v", err)
	}

	fmt.Printf("Found %d studies\n", len(results))
	for i, result := range results {
		studyUID, _ := result.GetString(tag.StudyInstanceUID)
		patientName, _ := result.GetString(tag.PatientName)
		fmt.Printf("Study %d: %s - %s\n", i+1, studyUID, patientName)
	}
}

// ExampleClient_CFindWithCallback demonstrates streaming C-FIND results with a callback.
func ExampleClient_CFindWithCallback() {
	c := client.New(
		client.WithCallingAE("MY_SCU"),
		client.WithCalledAE("PACS_SCP"),
	)

	c.AddPresentationContext(
		"1.2.840.10008.5.1.4.1.2.2.1", // Study Root Query/Retrieve - FIND
		"1.2.840.10008.1.2",
	)

	ctx := context.Background()
	if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer c.Close()

	query := dataset.New()
	query.Add(element.NewString(tag.PatientName, vr.PN, []string{"*"})) // All patients

	// Process results as they arrive
	count := 0
	err := c.CFindWithCallback(ctx, dimse.QueryRetrieveLevelPatient, query,
		func(result *dataset.Dataset) bool {
			count++
			patientName, _ := result.GetString(tag.PatientName)
			patientID, _ := result.GetString(tag.PatientID)
			fmt.Printf("Patient %d: %s (%s)\n", count, patientName, patientID)

			// Return true to continue receiving results, false to stop
			return count < 10 // Stop after 10 results
		})

	if err != nil {
		log.Fatalf("C-FIND failed: %v", err)
	}

	fmt.Printf("Processed %d patients\n", count)
}

// ExampleClient_CStoreMultiple demonstrates storing multiple datasets efficiently.
func ExampleClient_CStoreMultiple() {
	c := client.New(
		client.WithCallingAE("MY_SCU"),
		client.WithCalledAE("PACS_SCP"),
	)

	c.AddPresentationContext(
		"1.2.840.10008.5.1.4.1.1.2", // CT Image Storage
		"1.2.840.10008.1.2",
	)

	ctx := context.Background()
	if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer c.Close()

	// Create multiple datasets (e.g., from a CT series)
	datasets := make([]*dataset.Dataset, 100)
	for i := 0; i < 100; i++ {
		ds := dataset.New()
		ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
		instanceUID := fmt.Sprintf("1.2.840.113619.2.1.%d", i)
		ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{instanceUID}))
		ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"DOE^JOHN"}))
		datasets[i] = ds
	}

	// Store all datasets
	count, err := c.CStoreMultiple(ctx, datasets)
	if err != nil {
		log.Fatalf("C-STORE failed after %d datasets: %v", count, err)
	}

	fmt.Printf("Successfully stored %d datasets\n", count)
}

// ExampleClient_Ping demonstrates using Ping to verify connection health.
func ExampleClient_Ping() {
	c := client.New(
		client.WithCallingAE("MY_SCU"),
		client.WithCalledAE("PACS_SCP"),
	)

	c.AddPresentationContext(
		"1.2.840.10008.1.1", // Verification SOP Class
		"1.2.840.10008.1.2",
	)

	ctx := context.Background()
	if err := c.Connect(ctx, "192.168.1.100", 104); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer c.Close()

	// Use Ping to verify connection is still alive
	if err := c.Ping(ctx); err != nil {
		log.Fatalf("Connection is down: %v", err)
	}

	fmt.Println("Connection is healthy")
}

// ExampleDial demonstrates the convenience Dial function.
func ExampleDial() {
	// Dial is a convenience function that creates a client,
	// adds default presentation contexts, and connects in one call
	ctx := context.Background()
	c, err := client.Dial(ctx, "192.168.1.100", 104,
		client.WithCallingAE("MY_SCU"),
		client.WithCalledAE("PACS_SCP"),
		client.WithConnectTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer c.Close()

	// Client is ready to use
	if err := c.Ping(ctx); err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	fmt.Println("Connected and verified successfully")
}
