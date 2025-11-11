// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package dimse_test

import (
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
)

// Example showing how to use MessageIDGenerator for automatic MessageID assignment
func ExampleMessageIDGenerator() {
	// Create a MessageID generator (typically one per Association/Client)
	idGen := dimse.NewMessageIDGenerator()

	// Create requests with messageID = 0 (unassigned)
	echoReq := dimse.NewCEchoRequest()

	// Automatically assign MessageID before sending
	msgID, _ := idGen.AssignMessageID(echoReq)
	fmt.Printf("C-ECHO MessageID: %d\n", msgID)

	// Create another request
    ds := dataset.New()
    _ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
    _ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5"}))

	storeReq, _ := dimse.NewCStoreRequest(ds)
	msgID, _ = idGen.AssignMessageID(storeReq)
	fmt.Printf("C-STORE MessageID: %d\n", msgID)

	// Create a third request
	query := dataset.New()
	findReq := dimse.NewCFindRequestPatientRoot(dimse.QueryRetrieveLevelPatient, query)
	msgID, _ = idGen.AssignMessageID(findReq)
	fmt.Printf("C-FIND MessageID: %d\n", msgID)

	// Output:
	// C-ECHO MessageID: 1
	// C-STORE MessageID: 2
	// C-FIND MessageID: 3
}

// Example showing manual MessageID specification (backward compatibility)
func ExampleMessageIDGenerator_manual() {
	// You can still specify MessageID manually if needed
    echoReq := dimse.NewCEchoRequest()
    _ = echoReq.SetMessageID(100)
	fmt.Printf("Manual MessageID: %d\n", echoReq.MessageID())

	// Generator respects manually set MessageID
	idGen := dimse.NewMessageIDGenerator()
	msgID, _ := idGen.AssignMessageID(echoReq)
	fmt.Printf("MessageID after assignment: %d\n", msgID)

	// Output:
	// Manual MessageID: 100
	// MessageID after assignment: 100
}

// Example showing how Client/Association would use MessageIDGenerator
func ExampleMessageIDGenerator_clientUsage() {
	// Simulated Client structure
	type Client struct {
		messageIDGen *dimse.MessageIDGenerator
	}

	client := &Client{
		messageIDGen: dimse.NewMessageIDGenerator(),
	}

	// Send method automatically assigns MessageID
	send := func(c *Client, msg dimse.Message) error {
		// Assign MessageID if not set
		msgID, err := c.messageIDGen.AssignMessageID(msg)
		if err != nil {
			return err
		}

		fmt.Printf("Sending message with MessageID: %d\n", msgID)
		// ... actual sending logic ...
		return nil
	}

	// User creates requests without worrying about MessageID
	echoReq := dimse.NewCEchoRequest()
    _ = send(client, echoReq)

    ds := dataset.New()
    _ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{"1.2.840.10008.5.1.4.1.1.2"}))
    _ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{"1.2.3.4.5"}))
    storeReq, _ := dimse.NewCStoreRequest(ds)
    _ = send(client, storeReq)

	// Output:
	// Sending message with MessageID: 1
	// Sending message with MessageID: 2
}
