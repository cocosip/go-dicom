// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package client

import (
	"context"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// CEcho sends a C-ECHO request and waits for the response.
// This is used to verify connectivity and association status.
//
// Returns an error if:
//   - Not connected
//   - Send fails
//   - Response status is not success (0x0000)
//
// Example:
//
//	if err := client.CEcho(ctx); err != nil {
//	    return fmt.Errorf("C-ECHO failed: %w", err)
//	}
func (c *Client) CEcho(ctx context.Context) error {
	if !c.connected {
		return fmt.Errorf("client not connected")
	}

	// Create C-ECHO request
	req := dimse.NewCEchoRequest()

	// Send and wait for response
	resp, err := c.service.SendCEcho(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send C-ECHO: %w", err)
	}

	// Check status
	if !resp.Status().IsSuccess() {
		return fmt.Errorf("C-ECHO failed with status: %s", resp.Status())
	}

	return nil
}

// CStore sends a C-STORE request to store a DICOM dataset on the server.
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - ds: The dataset to store (must contain SOPClassUID and SOPInstanceUID)
//
// Returns an error if:
//   - Not connected
//   - Dataset is invalid (missing required UIDs)
//   - Send fails
//   - Response status is not success (0x0000)
//
// Example:
//
//	if err := client.CStore(ctx, dataset); err != nil {
//	    return fmt.Errorf("C-STORE failed: %w", err)
//	}
func (c *Client) CStore(ctx context.Context, ds *dataset.Dataset) error {
	if !c.connected {
		return fmt.Errorf("client not connected")
	}

	if ds == nil {
		return fmt.Errorf("dataset is nil")
	}

	// Create C-STORE request
	req, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		return fmt.Errorf("failed to create C-STORE request: %w", err)
	}

	// Send and wait for response
	resp, err := c.service.SendCStore(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send C-STORE: %w", err)
	}

	// Check status
	if !resp.Status().IsSuccess() {
		return fmt.Errorf("C-STORE failed with status: %s", resp.Status())
	}

	return nil
}

// CFind sends a C-FIND request and collects all matching results.
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - level: Query/Retrieve level (PATIENT, STUDY, SERIES, IMAGE)
//   - query: Query dataset with search criteria
//
// Returns:
//   - Slice of matching datasets
//   - Error if the operation fails
//
// The method will collect all Pending responses (0xFF00) and return
// when it receives a Success (0x0000) or Failure status.
//
// Example:
//
//	query := dataset.New()
//	query.AddString(tag.PatientName, "DOE^JOHN*")
//	query.AddString(tag.StudyDate, "20240101-20241231")
//
//	results, err := client.CFind(ctx, dimse.QueryRetrieveLevelStudy, query)
//	if err != nil {
//	    return fmt.Errorf("C-FIND failed: %w", err)
//	}
//
//	for _, result := range results {
//	    // Process each result
//	}
func (c *Client) CFind(ctx context.Context, level dimse.QueryRetrieveLevel, query *dataset.Dataset) ([]*dataset.Dataset, error) {
	if !c.connected {
		return nil, fmt.Errorf("client not connected")
	}

	if query == nil {
		return nil, fmt.Errorf("query is nil")
	}

	// Create C-FIND request
	req := dimse.NewCFindRequest(level, query)

	// Send request
	respCh, err := c.service.SendCFind(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to send C-FIND: %w", err)
	}

	// Collect all responses
	var results []*dataset.Dataset

	for {
		select {
		case resp, ok := <-respCh:
			if !ok {
				// Channel closed, no more responses
				return results, nil
			}

			statusCode := resp.StatusCode()

			// Pending status (0xFF00) - more results coming
			if statusCode == 0xFF00 {
				// Extract identifier dataset
				if resp.HasIdentifier() {
					identifier := resp.DataDataset()
					if identifier != nil {
						results = append(results, identifier)
					}
				}
				continue
			}

			// Success status (0x0000) - final response
			if statusCode == 0x0000 {
				return results, nil
			}

			// Any other status is an error
			return results, fmt.Errorf("C-FIND failed with status: 0x%04X", statusCode)

		case <-ctx.Done():
			return results, ctx.Err()
		}
	}
}

// CFindWithCallback sends a C-FIND request and calls the callback for each result.
// This is useful for processing results as they arrive without collecting them all in memory.
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - level: Query/Retrieve level (PATIENT, STUDY, SERIES, IMAGE)
//   - query: Query dataset with search criteria
//   - callback: Function called for each result. Return false to stop receiving results.
//
// Example:
//
//	err := client.CFindWithCallback(ctx, dimse.QueryRetrieveLevelStudy, query,
//	    func(result *dataset.Dataset) bool {
//	        // Process result
//	        fmt.Printf("Found study: %s\n", result.GetString(tag.StudyInstanceUID))
//	        return true // Continue receiving
//	    })
func (c *Client) CFindWithCallback(ctx context.Context, level dimse.QueryRetrieveLevel, query *dataset.Dataset, callback func(*dataset.Dataset) bool) error {
	if !c.connected {
		return fmt.Errorf("client not connected")
	}

	if query == nil {
		return fmt.Errorf("query is nil")
	}

	if callback == nil {
		return fmt.Errorf("callback is nil")
	}

	// Create C-FIND request
	req := dimse.NewCFindRequest(level, query)

	// Send request - returns channel of responses
	respCh, err := c.service.SendCFind(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send C-FIND: %w", err)
	}

	// Process responses
	for {
		select {
		case resp, ok := <-respCh:
			if !ok {
				// Channel closed
				return nil
			}

			statusCode := resp.StatusCode()

			// Pending status (0xFF00) - more results coming
			if statusCode == 0xFF00 {
				if resp.HasIdentifier() {
					identifier := resp.DataDataset()
					// Call callback, stop if it returns false
					if identifier != nil && !callback(identifier) {
						return nil
					}
				}
				continue
			}

			// Success status (0x0000) - final response
			if statusCode == 0x0000 {
				return nil
			}

			// Any other status is an error
			return fmt.Errorf("C-FIND failed with status: 0x%04X", statusCode)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// CStoreWithPriority stores a dataset with a specific priority.
//
// Priority values:
//   - dimse.PriorityLow (0x0002)
//   - dimse.PriorityMedium (0x0000) - default
//   - dimse.PriorityHigh (0x0001)
func (c *Client) CStoreWithPriority(ctx context.Context, ds *dataset.Dataset, priority dimse.Priority) error {
	if !c.connected {
		return fmt.Errorf("client not connected")
	}

	if ds == nil {
		return fmt.Errorf("dataset is nil")
	}

	// Create C-STORE request
	req, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		return fmt.Errorf("failed to create C-STORE request: %w", err)
	}

	// Set priority
	req.SetPriority(uint16(priority))

	// Send and wait for response
	resp, err := c.service.SendCStore(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send C-STORE: %w", err)
	}

	// Check status
	if !resp.Status().IsSuccess() {
		return fmt.Errorf("C-STORE failed with status: %s", resp.Status())
	}

	return nil
}

// CFindPatientRoot is a convenience method for Patient Root Query/Retrieve.
func (c *Client) CFindPatientRoot(ctx context.Context, level dimse.QueryRetrieveLevel, query *dataset.Dataset) ([]*dataset.Dataset, error) {
	// TODO: Set SOP Class UID to Patient Root Query/Retrieve (1.2.840.10008.5.1.4.1.2.1.1)
	return c.CFind(ctx, level, query)
}

// CFindStudyRoot is a convenience method for Study Root Query/Retrieve.
func (c *Client) CFindStudyRoot(ctx context.Context, level dimse.QueryRetrieveLevel, query *dataset.Dataset) ([]*dataset.Dataset, error) {
	// TODO: Set SOP Class UID to Study Root Query/Retrieve (1.2.840.10008.5.1.4.1.2.2.1)
	return c.CFind(ctx, level, query)
}

// CStoreMultiple stores multiple datasets in sequence.
// This is a convenience method that calls CStore for each dataset.
// It stops at the first error and returns the number of successfully stored datasets.
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - datasets: Slice of datasets to store
//
// Returns:
//   - Number of successfully stored datasets
//   - Error if any store operation fails
//
// Example:
//
//	count, err := client.CStoreMultiple(ctx, datasets)
//	if err != nil {
//	    fmt.Printf("Stored %d out of %d datasets before error: %v\n", count, len(datasets), err)
//	}
func (c *Client) CStoreMultiple(ctx context.Context, datasets []*dataset.Dataset) (int, error) {
	if !c.connected {
		return 0, fmt.Errorf("client not connected")
	}

	if len(datasets) == 0 {
		return 0, nil
	}

	for i, ds := range datasets {
		if err := c.CStore(ctx, ds); err != nil {
			return i, fmt.Errorf("failed to store dataset %d: %w", i, err)
		}

		// Check context cancellation between stores
		select {
		case <-ctx.Done():
			return i + 1, ctx.Err()
		default:
		}
	}

	return len(datasets), nil
}

// Ping sends a C-ECHO request to verify the connection is still alive.
// This is an alias for CEcho with a more intuitive name.
func (c *Client) Ping(ctx context.Context) error {
	return c.CEcho(ctx)
}

// CMoveCallback is called for each C-MOVE response.
// Return false to stop receiving updates.
type CMoveCallback func(remaining, completed, failed, warning uint16) bool

// CMove requests that the server move DICOM instances to a specified destination.
//
// The level parameter specifies the query/retrieve level (Patient, Study, Series, or Image).
// The moveDestination specifies the AE title where instances should be moved to.
// The identifier contains the matching keys that identify which instances to move.
// The callback is called for each pending response with sub-operation counts.
//
// C-MOVE is typically used when the client wants instances sent to a third-party destination.
// If you want instances sent directly to the requesting client, use CGet instead.
//
// Example:
//
//	identifier := dataset.New()
//	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
//	err := client.CMove(ctx, dimse.QueryRetrieveLevelStudy, "DEST_AE", identifier,
//	    func(remaining, completed, failed, warning uint16) bool {
//	        fmt.Printf("Move progress: %d remaining, %d completed, %d failed, %d warning\n",
//	            remaining, completed, failed, warning)
//	        return true // Continue receiving updates
//	    })
func (c *Client) CMove(ctx context.Context, level dimse.QueryRetrieveLevel, moveDestination string,
	identifier *dataset.Dataset, callback CMoveCallback) error {

	if !c.connected {
		return fmt.Errorf("client not connected")
	}

	if identifier == nil {
		return fmt.Errorf("identifier is nil")
	}

	if moveDestination == "" {
		return fmt.Errorf("move destination is empty")
	}

	// Create C-MOVE request
	req := dimse.NewCMoveRequest(level, moveDestination, identifier)

	// Send request - returns channel of responses
	respCh, err := c.service.SendCMove(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send C-MOVE: %w", err)
	}

	// Process responses
	for {
		select {
		case resp, ok := <-respCh:
			if !ok {
				// Channel closed
				return nil
			}

			statusCode := resp.StatusCode()

			// Pending status (0xFF00) - sub-operations in progress
			if statusCode == 0xFF00 {
				if callback != nil && resp.HasSubOperationCounts() {
					// Call callback, stop if it returns false
					if !callback(
						resp.NumberOfRemainingSubOperations(),
						resp.NumberOfCompletedSubOperations(),
						resp.NumberOfFailedSubOperations(),
						resp.NumberOfWarningSubOperations(),
					) {
						return nil
					}
				}
				continue
			}

			// Success status (0x0000) - all sub-operations completed
			if statusCode == 0x0000 {
				return nil
			}

			// Warning status (0xB000) - completed with warnings
			if statusCode == 0xB000 {
				return nil
			}

			// Any other status is an error
			return fmt.Errorf("C-MOVE failed with status: 0x%04X", statusCode)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// CGetCallback is called for each C-GET response with sub-operation progress.
// Return false to stop receiving updates.
type CGetCallback func(remaining, completed, failed, warning uint16) bool

// CGet requests that the server retrieve DICOM instances and send them back to the client.
//
// The level parameter specifies the query/retrieve level (Patient, Study, Series, or Image).
// The identifier contains the matching keys that identify which instances to retrieve.
// The callback is called for each pending response with sub-operation counts.
//
// C-GET retrieves instances directly to the requesting client over the same association.
// The client must have a C-STORE handler configured to receive the instances.
//
// Note: C-GET combines the query and retrieval in a single operation. The server will
// send C-STORE sub-operations for each matching instance.
//
// Example:
//
//	identifier := dataset.New()
//	identifier.Add(element.NewString(tag.StudyInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
//	err := client.CGet(ctx, dimse.QueryRetrieveLevelStudy, identifier,
//	    func(remaining, completed, failed, warning uint16) bool {
//	        fmt.Printf("Get progress: %d remaining, %d completed, %d failed, %d warning\n",
//	            remaining, completed, failed, warning)
//	        return true // Continue receiving updates
//	    })
func (c *Client) CGet(ctx context.Context, level dimse.QueryRetrieveLevel,
	identifier *dataset.Dataset, callback CGetCallback) error {

	if !c.connected {
		return fmt.Errorf("client not connected")
	}

	if identifier == nil {
		return fmt.Errorf("identifier is nil")
	}

	// Create C-GET request
	req := dimse.NewCGetRequest(level, identifier)

	// Send request - returns channel of responses
	respCh, err := c.service.SendCGet(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to send C-GET: %w", err)
	}

	// Process responses
	for {
		select {
		case resp, ok := <-respCh:
			if !ok {
				// Channel closed
				return nil
			}

			statusCode := resp.StatusCode()

			// Pending status (0xFF00) - sub-operations in progress
			if statusCode == 0xFF00 {
				if callback != nil && resp.HasSubOperationCounts() {
					// Call callback, stop if it returns false
					if !callback(
						resp.NumberOfRemainingSubOperations(),
						resp.NumberOfCompletedSubOperations(),
						resp.NumberOfFailedSubOperations(),
						resp.NumberOfWarningSubOperations(),
					) {
						return nil
					}
				}
				continue
			}

			// Success status (0x0000) - all sub-operations completed
			if statusCode == 0x0000 {
				return nil
			}

			// Warning status (0xB000) - completed with warnings
			if statusCode == 0xB000 {
				return nil
			}

			// Any other status is an error
			return fmt.Errorf("C-GET failed with status: 0x%04X", statusCode)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
