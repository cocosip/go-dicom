// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package service

import (
	"bytes"
	"context"
	"io"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/imaging/codec"
	"github.com/cocosip/go-dicom/pkg/imaging/imagetypes"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
	"github.com/cocosip/go-dicom/pkg/network/association"
	"github.com/cocosip/go-dicom/pkg/network/dimse"
	"github.com/cocosip/go-dicom/pkg/network/pdu"
)

func TestEncodeDIMSEMessage(t *testing.T) {
	// Create a C-ECHO request
	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	// Encode the message
	commandData, datasetData, err := EncodeDIMSEMessage(req, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("EncodeDIMSEMessage failed: %v", err)
	}

	// Command data should not be empty
	if len(commandData) == 0 {
		t.Error("Command data is empty")
	}

	// C-ECHO should not have dataset
	if datasetData != nil {
		t.Error("C-ECHO should not have dataset")
	}
}

func TestEncodeDIMSEMessageWithDataset(t *testing.T) {
	t.Skip("Skipping test - CreateTestDataset helper not yet implemented")
	// TODO: Implement this test once we have a test dataset creation helper
	// Create a C-STORE request with dataset
	// ds := dimse.CreateTestDataset() // Helper to create test dataset
	// req, err := dimse.NewCStoreRequest(ds)
	// if err != nil {
	// 	t.Skip("Skipping test - NewCStoreRequest requires valid dataset")
	// 	return
	// }
	//
	// req.SetMessageID(1)
	//
	// // Encode the message
	// commandData, datasetData, err := EncodeDIMSEMessage(req, transfer.ExplicitVRLittleEndian)
	// if err != nil {
	// 	t.Fatalf("EncodeDIMSEMessage failed: %v", err)
	// }
	//
	// // Both command and dataset should be present
	// if len(commandData) == 0 {
	// 	t.Error("Command data is empty")
	// }
	//
	// if len(datasetData) == 0 {
	// 	t.Error("Dataset data is empty for C-STORE")
	// }
}

func TestServiceStart(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service
	service := NewService(client, nil)
	defer func() { _ = service.Close() }()

	// Start service
	err := service.Start()
	if err != nil {
		t.Fatalf("Start failed: %v", err)
	}

	// Give goroutines time to start
	time.Sleep(50 * time.Millisecond)

	// Close service
	_ = service.Close()

	// Give goroutines time to stop
	time.Sleep(50 * time.Millisecond)
}

func TestSendMessage(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create an association with a presentation context for C-ECHO
	assoc := createTestAssociation()

	// Create service with association
	service := NewService(client, assoc, WithAssociationRequestor(true))
	defer func() { _ = service.Close() }()

	// Create a send request
	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID failed: %v", err)
	}

	sendReq := &sendRequest{
		message:  req,
		resultCh: make(chan error, 1),
	}

	// Start a goroutine to read from server side
	go func() {
		// Just read and discard data to prevent blocking
		buf := make([]byte, 4096)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	// Send message
	err := service.sendMessage(sendReq)
	if err != nil {
		t.Fatalf("sendMessage failed: %v", err)
	}
}

func TestSendMessageUsesNegotiatedRemoteMaxPDULength(t *testing.T) {
	assoc := createTestAssociation()
	assoc.MaxPDULength = 256

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true), WithMaxPDULength(4096))

	ds := dataset.NewWithTransferSyntax(transfer.ExplicitVRLittleEndian)
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	_ = ds.Add(element.NewOtherByte(tag.PixelData, bytes.Repeat([]byte{0x7f}, 2048)))

	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}
	if err := msg.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}

	sendReq := &sendRequest{
		message:  msg,
		resultCh: make(chan error, 1),
	}

	if err := service.sendMessage(sendReq); err != nil {
		t.Fatalf("sendMessage() error = %v", err)
	}
	if len(conn.writes) < 2 {
		t.Fatalf("sendMessage() wrote %d PDU(s), want fragmentation", len(conn.writes))
	}
	for i, write := range conn.writes {
		if len(write) > int(assoc.MaxPDULength) {
			t.Fatalf("PDU %d length = %d, exceeds negotiated remote maximum %d", i, len(write), assoc.MaxPDULength)
		}
	}
}

func TestSendMessagePrefersCStoreSourceTransferSyntax(t *testing.T) {
	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRBigEndian)
	addAcceptedPresentationContext(t, assoc, 5, transfer.ExplicitVRLittleEndian)

	originalPixelData := []byte{0x34, 0x12}
	ds := newCStorePixelDataset(t, transfer.ExplicitVRLittleEndian, originalPixelData)
	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	if err := service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)}); err != nil {
		t.Fatalf("sendMessage() error = %v", err)
	}

	pdvs := readRecordedPDVs(t, conn)
	if len(pdvs) == 0 {
		t.Fatal("sendMessage() wrote no PDVs")
	}
	for i, pdv := range pdvs {
		if pdv.PresentationContextID != 5 {
			t.Fatalf("PDV %d presentation context ID = %d, want 5 for the source transfer syntax", i, pdv.PresentationContextID)
		}
	}

	pixelData, ok := ds.Get(tag.PixelData)
	if !ok {
		t.Fatal("source dataset lost PixelData")
	}
	if got := pixelData.(*element.OtherWord).GetData(); !bytes.Equal(got, originalPixelData) {
		t.Fatalf("source PixelData = %v, want unchanged %v", got, originalPixelData)
	}
}

func TestSendMessageTranscodesCStoreToAcceptedTransferSyntax(t *testing.T) {
	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRBigEndian)

	originalPixelData := []byte{0x34, 0x12}
	ds := newCStorePixelDataset(t, transfer.ExplicitVRLittleEndian, originalPixelData)
	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	if err := service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)}); err != nil {
		t.Fatalf("sendMessage() error = %v", err)
	}

	commandData, datasetData := recordedDIMSEData(t, conn)
	_, sentDataset, err := DecodeDIMSEMessage(commandData, datasetData, transfer.ExplicitVRBigEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}
	sentPixelData, ok := sentDataset.Get(tag.PixelData)
	if !ok {
		t.Fatal("sent dataset has no PixelData")
	}
	if got := sentPixelData.(*element.OtherWord).GetData(); !bytes.Equal(got, []byte{0x12, 0x34}) {
		t.Fatalf("sent PixelData = %v, want Big Endian bytes [18 52]", got)
	}

	if ds.InternalTransferSyntax() != transfer.ExplicitVRLittleEndian {
		t.Fatalf("source transfer syntax = %v, want unchanged Explicit VR Little Endian", ds.InternalTransferSyntax())
	}
	sourcePixelData, _ := ds.Get(tag.PixelData)
	if got := sourcePixelData.(*element.OtherWord).GetData(); !bytes.Equal(got, originalPixelData) {
		t.Fatalf("source PixelData = %v, want unchanged %v", got, originalPixelData)
	}
}

func TestSendMessageRejectsCStoreWithoutUsableTransferSyntax(t *testing.T) {
	registry := codec.GetGlobalRegistry()
	previousCodec, hadPreviousCodec := registry.GetCodec(transfer.JPEG2000Lossless)
	registry.UnregisterCodec(transfer.JPEG2000Lossless)
	t.Cleanup(func() {
		if hadPreviousCodec {
			registry.RegisterCodec(transfer.JPEG2000Lossless, previousCodec)
		}
	})

	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRLittleEndian)

	ds := newCStorePixelDataset(t, transfer.JPEG2000Lossless, []byte{0x00, 0x00})
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0xff, 0x4f, 0xff, 0x51}))
	if err := ds.AddOrUpdate(fragments); err != nil {
		t.Fatalf("Dataset.AddOrUpdate(PixelData) error = %v", err)
	}
	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	err = service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)})
	if err == nil {
		t.Fatal("sendMessage() error = nil, want untranscodable transfer syntax error")
	}
	for _, want := range []string{
		testCTImageStorageUID,
		transfer.JPEG2000Lossless.UID().UID(),
		transfer.ExplicitVRLittleEndian.UID().UID(),
	} {
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("sendMessage() error = %q, want it to contain %q", err, want)
		}
	}
	if len(conn.writes) != 0 {
		t.Fatalf("sendMessage() wrote %d network chunks before rejecting the dataset", len(conn.writes))
	}
}

func TestSendMessageUsesRegisteredCodecForCStoreTranscoding(t *testing.T) {
	registry := codec.GetGlobalRegistry()
	previousCodec, hadPreviousCodec := registry.GetCodec(transfer.JPEG2000Lossless)
	registry.RegisterCodec(transfer.JPEG2000Lossless, passthroughCodec{syntax: transfer.JPEG2000Lossless})
	t.Cleanup(func() {
		if hadPreviousCodec {
			registry.RegisterCodec(transfer.JPEG2000Lossless, previousCodec)
		} else {
			registry.UnregisterCodec(transfer.JPEG2000Lossless)
		}
	})

	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRLittleEndian)

	ds := newCStorePixelDataset(t, transfer.JPEG2000Lossless, []byte{0x00, 0x00})
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0x34, 0x12}))
	if err := ds.AddOrUpdate(fragments); err != nil {
		t.Fatalf("Dataset.AddOrUpdate(PixelData) error = %v", err)
	}
	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	if err := service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)}); err != nil {
		t.Fatalf("sendMessage() error = %v", err)
	}

	commandData, datasetData := recordedDIMSEData(t, conn)
	_, sentDataset, err := DecodeDIMSEMessage(commandData, datasetData, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}
	sentPixelData, ok := sentDataset.Get(tag.PixelData)
	if !ok {
		t.Fatal("sent dataset has no PixelData")
	}
	if got := sentPixelData.(*element.OtherWord).GetData(); !bytes.Equal(got, []byte{0x34, 0x12}) {
		t.Fatalf("sent PixelData = %v, want decoded bytes [52 18]", got)
	}
	if _, ok := ds.Get(tag.PixelData); !ok {
		t.Fatal("source dataset lost PixelData")
	}
}

func TestSendMessageConvertsDecodedCStorePixelsToAcceptedBigEndian(t *testing.T) {
	registry := codec.GetGlobalRegistry()
	previousCodec, hadPreviousCodec := registry.GetCodec(transfer.JPEG2000Lossless)
	registry.RegisterCodec(transfer.JPEG2000Lossless, passthroughCodec{syntax: transfer.JPEG2000Lossless})
	t.Cleanup(func() {
		if hadPreviousCodec {
			registry.RegisterCodec(transfer.JPEG2000Lossless, previousCodec)
		} else {
			registry.UnregisterCodec(transfer.JPEG2000Lossless)
		}
	})

	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRBigEndian)
	ds := newCStorePixelDataset(t, transfer.JPEG2000Lossless, []byte{0x00, 0x00})
	fragments := element.NewOtherByteFragment(tag.PixelData)
	fragments.AddFragment(buffer.NewMemory([]byte{0x34, 0x12}))
	if err := ds.AddOrUpdate(fragments); err != nil {
		t.Fatalf("Dataset.AddOrUpdate(PixelData) error = %v", err)
	}
	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	if err := service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)}); err != nil {
		t.Fatalf("sendMessage() error = %v", err)
	}

	commandData, datasetData := recordedDIMSEData(t, conn)
	_, sentDataset, err := DecodeDIMSEMessage(commandData, datasetData, transfer.ExplicitVRBigEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}
	sentPixelData, ok := sentDataset.Get(tag.PixelData)
	if !ok {
		t.Fatal("sent dataset has no PixelData")
	}
	if got := sentPixelData.(*element.OtherWord).GetData(); !bytes.Equal(got, []byte{0x12, 0x34}) {
		t.Fatalf("sent PixelData = %v, want Big Endian bytes [18 52]", got)
	}
}

type passthroughCodec struct {
	syntax *transfer.Syntax
}

func (c passthroughCodec) Name() string {
	return "passthrough"
}

func (c passthroughCodec) TransferSyntax() *transfer.Syntax {
	return c.syntax
}

func (passthroughCodec) GetDefaultParameters() codec.Parameters {
	return codec.NewBaseParameters()
}

func (passthroughCodec) Encode(oldPixelData, newPixelData imagetypes.PixelData, _ codec.Parameters) error {
	return copyPixelFrames(oldPixelData, newPixelData)
}

func (passthroughCodec) Decode(oldPixelData, newPixelData imagetypes.PixelData, _ codec.Parameters) error {
	return copyPixelFrames(oldPixelData, newPixelData)
}

func copyPixelFrames(oldPixelData, newPixelData imagetypes.PixelData) error {
	for i := 0; i < oldPixelData.FrameCount(); i++ {
		frame, err := oldPixelData.GetFrame(i)
		if err != nil {
			return err
		}
		if err := newPixelData.AddFrame(frame); err != nil {
			return err
		}
	}
	return nil
}

func TestSendMessageRejectsCStorePixelDataWithoutSourceTransferSyntax(t *testing.T) {
	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRLittleEndian)

	ds := newCStorePixelDataset(t, transfer.ExplicitVRLittleEndian, []byte{0x34, 0x12})
	ds.SetInternalTransferSyntax(nil)
	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	err = service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)})
	if err == nil {
		t.Fatal("sendMessage() error = nil, want missing source transfer syntax error")
	}
	if !strings.Contains(err.Error(), "source transfer syntax is unknown") {
		t.Fatalf("sendMessage() error = %q, want unknown source transfer syntax context", err)
	}
	if len(conn.writes) != 0 {
		t.Fatalf("sendMessage() wrote %d network chunks before rejecting the dataset", len(conn.writes))
	}
}

func TestSendMessageDoesNotRequireCodecWithoutPixelData(t *testing.T) {
	registry := codec.GetGlobalRegistry()
	previousCodec, hadPreviousCodec := registry.GetCodec(transfer.JPEG2000Lossless)
	registry.UnregisterCodec(transfer.JPEG2000Lossless)
	t.Cleanup(func() {
		if hadPreviousCodec {
			registry.RegisterCodec(transfer.JPEG2000Lossless, previousCodec)
		}
	})

	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRLittleEndian)
	ds := dataset.NewWithTransferSyntax(transfer.JPEG2000Lossless)
	_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
	_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testStudyInstanceUID}))
	_ = ds.Add(element.NewString(tag.PatientName, vr.PN, []string{"NoPixel^Object"}))
	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	if err := service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)}); err != nil {
		t.Fatalf("sendMessage() error = %v, want direct encoding without a pixel codec", err)
	}

	commandData, datasetData := recordedDIMSEData(t, conn)
	_, sentDataset, err := DecodeDIMSEMessage(commandData, datasetData, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}
	if got, ok := sentDataset.GetString(tag.PatientName); !ok || got != "NoPixel^Object" {
		t.Fatalf("sent PatientName = %q, %v; want NoPixel^Object, true", got, ok)
	}
}

func TestSendMessageValidatesExplicitCStoreContextWithoutPixelData(t *testing.T) {
	const verificationSOPClassUID = "1.2.840.10008.1.1"

	tests := []struct {
		name    string
		context *association.PresentationContext
		wantErr string
	}{
		{
			name:    "rejected context",
			context: association.NewPresentationContext(3, testCTImageStorageUID, transfer.ExplicitVRLittleEndian),
			wantErr: "is not accepted",
		},
		{
			name: "wrong SOP Class",
			context: func() *association.PresentationContext {
				pc := association.NewPresentationContext(3, verificationSOPClassUID, transfer.ExplicitVRLittleEndian)
				pc.Accept(transfer.ExplicitVRLittleEndian)
				return pc
			}(),
			wantErr: "is for SOP Class",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
			if err := assoc.AddPresentationContext(tt.context); err != nil {
				t.Fatalf("AddPresentationContext() error = %v", err)
			}
			ds := dataset.New()
			_ = ds.Add(element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}))
			_ = ds.Add(element.NewString(tag.SOPInstanceUID, vr.UI, []string{testStudyInstanceUID}))
			msg, err := dimse.NewCStoreRequest(ds)
			if err != nil {
				t.Fatalf("NewCStoreRequest() error = %v", err)
			}
			msg.SetPresentationContextID(3)

			conn := &recordingConn{}
			service := NewService(conn, assoc, WithAssociationRequestor(true))
			err = service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)})
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("sendMessage() error = %v, want error containing %q", err, tt.wantErr)
			}
			if len(conn.writes) != 0 {
				t.Fatalf("sendMessage() wrote %d network chunks before rejecting the context", len(conn.writes))
			}
		})
	}
}

func TestDecodedCStoreDatasetCanBeForwarded(t *testing.T) {
	source := newCStorePixelDataset(t, transfer.ExplicitVRLittleEndian, []byte{0x34, 0x12})
	sourceMessage, err := dimse.NewCStoreRequest(source)
	if err != nil {
		t.Fatalf("NewCStoreRequest(source) error = %v", err)
	}
	commandData, datasetData, err := EncodeDIMSEMessage(sourceMessage, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("EncodeDIMSEMessage() error = %v", err)
	}
	_, received, err := DecodeDIMSEMessage(commandData, datasetData, transfer.ExplicitVRLittleEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}
	if received.InternalTransferSyntax() != transfer.ExplicitVRLittleEndian {
		t.Fatalf("received transfer syntax = %v, want Explicit VR Little Endian", received.InternalTransferSyntax())
	}

	forwardMessage, err := dimse.NewCStoreRequest(received)
	if err != nil {
		t.Fatalf("NewCStoreRequest(received) error = %v", err)
	}
	assoc := association.NewAssociation("FORWARD-SCU", "DESTINATION-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRLittleEndian)
	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	if err := service.sendMessage(&sendRequest{message: forwardMessage, resultCh: make(chan error, 1)}); err != nil {
		t.Fatalf("forward sendMessage() error = %v", err)
	}
	if len(conn.writes) == 0 {
		t.Fatal("forward sendMessage() wrote no network data")
	}
}

func TestSendMessageHonorsExplicitCStorePresentationContext(t *testing.T) {
	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")
	addAcceptedPresentationContext(t, assoc, 3, transfer.ExplicitVRBigEndian)
	addAcceptedPresentationContext(t, assoc, 5, transfer.ExplicitVRLittleEndian)

	ds := newCStorePixelDataset(t, transfer.ExplicitVRLittleEndian, []byte{0x34, 0x12})
	msg, err := dimse.NewCStoreRequest(ds)
	if err != nil {
		t.Fatalf("NewCStoreRequest() error = %v", err)
	}
	msg.SetPresentationContextID(3)

	conn := &recordingConn{}
	service := NewService(conn, assoc, WithAssociationRequestor(true))
	if err := service.sendMessage(&sendRequest{message: msg, resultCh: make(chan error, 1)}); err != nil {
		t.Fatalf("sendMessage() error = %v", err)
	}

	commandData, datasetData := recordedDIMSEData(t, conn)
	_, sentDataset, err := DecodeDIMSEMessage(commandData, datasetData, transfer.ExplicitVRBigEndian)
	if err != nil {
		t.Fatalf("DecodeDIMSEMessage() error = %v", err)
	}
	sentPixelData, _ := sentDataset.Get(tag.PixelData)
	if got := sentPixelData.(*element.OtherWord).GetData(); !bytes.Equal(got, []byte{0x12, 0x34}) {
		t.Fatalf("sent PixelData = %v, want Big Endian bytes [18 52]", got)
	}
}

func addAcceptedPresentationContext(t *testing.T, assoc *association.Association, id byte, ts *transfer.Syntax) {
	t.Helper()
	pc := association.NewPresentationContext(id, testCTImageStorageUID, ts)
	pc.Accept(ts)
	if err := assoc.AddPresentationContext(pc); err != nil {
		t.Fatalf("AddPresentationContext() error = %v", err)
	}
}

func newCStorePixelDataset(t *testing.T, ts *transfer.Syntax, pixelData []byte) *dataset.Dataset {
	t.Helper()
	ds := dataset.NewWithTransferSyntax(ts)
	for _, elem := range []element.Element{
		element.NewString(tag.SOPClassUID, vr.UI, []string{testCTImageStorageUID}),
		element.NewString(tag.SOPInstanceUID, vr.UI, []string{testStudyInstanceUID}),
		element.NewUnsignedShort(tag.Rows, []uint16{1}),
		element.NewUnsignedShort(tag.Columns, []uint16{1}),
		element.NewUnsignedShort(tag.BitsAllocated, []uint16{16}),
		element.NewUnsignedShort(tag.BitsStored, []uint16{16}),
		element.NewUnsignedShort(tag.HighBit, []uint16{15}),
		element.NewUnsignedShort(tag.SamplesPerPixel, []uint16{1}),
		element.NewUnsignedShort(tag.PixelRepresentation, []uint16{0}),
		element.NewString(tag.PhotometricInterpretation, vr.CS, []string{"MONOCHROME2"}),
		element.NewOtherWord(tag.PixelData, append([]byte(nil), pixelData...)),
	} {
		if err := ds.Add(elem); err != nil {
			t.Fatalf("Dataset.Add(%s) error = %v", elem.Tag(), err)
		}
	}
	return ds
}

func readRecordedPDVs(t *testing.T, conn *recordingConn) []*PDV {
	t.Helper()
	var stream bytes.Buffer
	for _, write := range conn.writes {
		_, _ = stream.Write(write)
	}

	var result []*PDV
	for i := 0; stream.Len() > 0; i++ {
		raw, err := pdu.ReadPDU(&stream)
		if err != nil {
			t.Fatalf("read recorded PDU %d: %v", i, err)
		}
		pdvs, err := ParsePDataTFPDU(raw)
		if err != nil {
			t.Fatalf("parse recorded PDU %d: %v", i, err)
		}
		result = append(result, pdvs...)
	}
	return result
}

func recordedDIMSEData(t *testing.T, conn *recordingConn) ([]byte, []byte) {
	t.Helper()
	const wantContextID byte = 3
	var commandData []byte
	var datasetData []byte
	for i, pdv := range readRecordedPDVs(t, conn) {
		if pdv.PresentationContextID != wantContextID {
			t.Fatalf("PDV %d presentation context ID = %d, want %d", i, pdv.PresentationContextID, wantContextID)
		}
		if pdv.IsCommand {
			commandData = append(commandData, pdv.Data...)
		} else {
			datasetData = append(datasetData, pdv.Data...)
		}
	}
	if len(commandData) == 0 || len(datasetData) == 0 {
		t.Fatalf("recorded DIMSE data has command=%d bytes, dataset=%d bytes", len(commandData), len(datasetData))
	}
	return commandData, datasetData
}

func TestSendMessageReturnsErrorWhenMaxPDUTooSmall(t *testing.T) {
	assoc := createTestAssociation()
	assoc.MaxPDULength = 11

	service := NewService(&recordingConn{}, assoc, WithAssociationRequestor(true))

	req := dimse.NewCEchoRequest()
	if err := req.SetMessageID(1); err != nil {
		t.Fatalf("SetMessageID() error = %v", err)
	}
	sendReq := &sendRequest{
		message:  req,
		resultCh: make(chan error, 1),
	}

	err := service.sendMessage(sendReq)
	if err == nil {
		t.Fatal("sendMessage() error = nil, want max PDU length error")
	}
}

type recordingConn struct {
	writes [][]byte
}

func (c *recordingConn) Read(_ []byte) (int, error) {
	return 0, io.EOF
}

func (c *recordingConn) Write(p []byte) (int, error) {
	c.writes = append(c.writes, append([]byte(nil), p...))
	return len(p), nil
}

func (c *recordingConn) Close() error {
	return nil
}

func (c *recordingConn) LocalAddr() net.Addr {
	return dummyAddr("local")
}

func (c *recordingConn) RemoteAddr() net.Addr {
	return dummyAddr("remote")
}

func (c *recordingConn) SetDeadline(_ time.Time) error {
	return nil
}

func (c *recordingConn) SetReadDeadline(_ time.Time) error {
	return nil
}

func (c *recordingConn) SetWriteDeadline(_ time.Time) error {
	return nil
}

type dummyAddr string

func (a dummyAddr) Network() string {
	return string(a)
}

func (a dummyAddr) String() string {
	return string(a)
}

// createTestAssociation creates a test association with common presentation contexts
func createTestAssociation() *association.Association {
	assoc := association.NewAssociation("TEST-SCU", "TEST-SCP")

	// Add presentation context for C-ECHO (Verification SOP Class)
	pc := association.NewPresentationContext(1, "1.2.840.10008.1.1", transfer.ExplicitVRLittleEndian)
	pc.AcceptedTransferSyntax = transfer.ExplicitVRLittleEndian
	pc.Result = association.ResultAcceptance
	_ = assoc.AddPresentationContext(pc)

	// Add presentation context for CT Image Storage (for C-STORE tests)
	pc2 := association.NewPresentationContext(3, testCTImageStorageUID, transfer.ExplicitVRLittleEndian)
	pc2.AcceptedTransferSyntax = transfer.ExplicitVRLittleEndian
	pc2.Result = association.ResultAcceptance
	_ = assoc.AddPresentationContext(pc2)

	assoc.SetEstablished(true)
	return assoc
}

func TestSendLoop_ContextCancellation(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service
	service := NewService(client, nil)
	defer func() { _ = service.Close() }()

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start send loop in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.sendLoop(ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

	// Cancel context
	cancel()

	// Wait for loop to exit
	select {
	case err := <-errCh:
		if err != context.Canceled {
			t.Errorf("Expected context.Canceled, got %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("sendLoop did not exit after context cancellation")
	}
}

func TestGroupPDVsIntoPDUs(t *testing.T) {
	assoc := createTestAssociation()
	service := NewService(nil, assoc, WithMaxPDULength(1000))

	tests := []struct {
		name           string
		pdvDataSizes   []int
		maxPDULength   uint32
		expectedGroups int
	}{
		{
			name:           "Single small PDV",
			pdvDataSizes:   []int{100},
			maxPDULength:   1000,
			expectedGroups: 1,
		},
		{
			name:           "Multiple small PDVs fit in one PDU",
			pdvDataSizes:   []int{100, 100, 100},
			maxPDULength:   1000,
			expectedGroups: 1,
		},
		{
			name:           "PDVs require multiple PDUs",
			pdvDataSizes:   []int{400, 400, 400},
			maxPDULength:   1000,
			expectedGroups: 2, // First: 6 (PDU) + 406 (PDV1) + 406 (PDV2) = 818; Second: 6 + 406 = 412
		},
		{
			name:           "Some PDVs can be packed",
			pdvDataSizes:   []int{300, 300, 100, 100},
			maxPDULength:   1000,
			expectedGroups: 1, // All fit: 6 (PDU) + 306*2 + 106*2 = 830
		},
		{
			name:           "Empty PDV list",
			pdvDataSizes:   []int{},
			maxPDULength:   1000,
			expectedGroups: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create PDVs with specified data sizes
			var pdvs []*PDV
			for _, size := range tt.pdvDataSizes {
				pdv := &PDV{
					PresentationContextID: 1,
					IsCommand:             false,
					IsLastFragment:        false,
					Data:                  make([]byte, size),
				}
				pdvs = append(pdvs, pdv)
			}

			// Update service maxPDULength for this test
			service.config.maxPDULength = tt.maxPDULength

			// Group PDVs
			groups := service.groupPDVsIntoPDUs(pdvs)

			if len(groups) != tt.expectedGroups {
				t.Errorf("Expected %d groups, got %d", tt.expectedGroups, len(groups))
			}

			// Verify each group doesn't exceed max PDU length
			const pduHeaderSize = 6
			for i, group := range groups {
				totalSize := pduHeaderSize
				for _, pdv := range group {
					pdvSize := 4 + 1 + 1 + len(pdv.Data) // length + context ID + header + data
					totalSize += pdvSize
				}
				if totalSize > int(tt.maxPDULength) {
					t.Errorf("Group %d exceeds max PDU length: %d > %d", i, totalSize, tt.maxPDULength)
				}
			}

			// Verify all PDVs are included
			totalPDVs := 0
			for _, group := range groups {
				totalPDVs += len(group)
			}
			if totalPDVs != len(pdvs) {
				t.Errorf("Expected %d PDVs total, got %d", len(pdvs), totalPDVs)
			}
		})
	}
}

func TestSendLoop_ServiceClose(t *testing.T) {
	// Create a pipe connection
	server, client := net.Pipe()
	defer func() { _ = server.Close() }()
	defer func() { _ = client.Close() }()

	// Create service
	service := NewService(client, nil)

	// Start send loop in goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- service.sendLoop(service.ctx)
	}()

	// Give loop time to start
	time.Sleep(10 * time.Millisecond)

	// Close service
	_ = service.Close()

	// Wait for loop to exit
	select {
	case err := <-errCh:
		if err != nil && err != context.Canceled {
			t.Errorf("sendLoop returned unexpected error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("sendLoop did not exit after service close")
	}
}
