// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package codec

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/transfer"
	"github.com/cocosip/go-dicom/pkg/dicom/uid"
)

const (
	registryFirstCodecName  = "first"
	registrySecondCodecName = "second"
)

type registryTestCodec struct {
	name   string
	syntax *transfer.Syntax
}

func (c registryTestCodec) Name() string { return c.name }

func (c registryTestCodec) TransferSyntax() *transfer.Syntax { return c.syntax }

func (registryTestCodec) DefaultParameters() Parameters { return NoParameters{} }

func (registryTestCodec) Encode(context.Context, FrameSource, FrameSink, Parameters) error {
	return nil
}

func (registryTestCodec) Decode(context.Context, FrameSource, FrameSink, Parameters) error {
	return nil
}

func TestGlobalRegistryDoesNotRegisterRLE(t *testing.T) {
	const rleLosslessUID = "1.2.840.10008.1.2.5"

	for _, registered := range GlobalRegistry().List() {
		if registered.TransferSyntax().UID().UID() == rleLosslessUID {
			t.Fatal("global registry must not register an RLE codec")
		}
	}
}

func TestRegistryRegisterRejectsInvalidAndDuplicateCodecs(t *testing.T) {
	registry := NewRegistry()
	if err := registry.Register(nil); !errors.Is(err, ErrNilCodec) {
		t.Fatalf("Register(nil) error = %v, want ErrNilCodec", err)
	}
	if err := registry.Register(registryTestCodec{name: "nil-syntax"}); !errors.Is(err, ErrNilTransferSyntax) {
		t.Fatalf("Register(nil syntax) error = %v, want ErrNilTransferSyntax", err)
	}

	first := registryTestCodec{name: registryFirstCodecName, syntax: transfer.JPEG2000Lossless}
	if err := registry.Register(first); err != nil {
		t.Fatalf("Register(first) error = %v", err)
	}
	if err := registry.Register(registryTestCodec{name: "duplicate", syntax: transfer.JPEG2000Lossless}); !errors.Is(err, ErrCodecAlreadyRegistered) {
		t.Fatalf("Register(duplicate) error = %v, want ErrCodecAlreadyRegistered", err)
	}
	if got, _ := registry.Lookup(transfer.JPEG2000Lossless); got.Name() != registryFirstCodecName {
		t.Fatalf("Lookup() after duplicate = %q, want first", got.Name())
	}
}

func TestRegistryReplaceAndUnregisterReturnPreviousCodec(t *testing.T) {
	registry := NewRegistry()
	first := registryTestCodec{name: registryFirstCodecName, syntax: transfer.JPEG2000Lossless}
	second := registryTestCodec{name: registrySecondCodecName, syntax: transfer.JPEG2000Lossless}
	if err := registry.Register(first); err != nil {
		t.Fatal(err)
	}

	previous, err := registry.Replace(second)
	if err != nil {
		t.Fatalf("Replace() error = %v", err)
	}
	if previous == nil || previous.Name() != registryFirstCodecName {
		t.Fatalf("Replace() previous = %#v, want first codec", previous)
	}
	removed, found := registry.Unregister(transfer.JPEG2000Lossless)
	if !found || removed == nil || removed.Name() != registrySecondCodecName {
		t.Fatalf("Unregister() = (%#v, %v), want second codec and true", removed, found)
	}
	if _, found := registry.Unregister(nil); found {
		t.Fatal("Unregister(nil) found a codec")
	}
}

func TestRegistryListIsSortedSnapshot(t *testing.T) {
	registry := NewRegistry()
	for _, registered := range []Codec{
		registryTestCodec{name: "rle", syntax: transfer.RLELossless},
		registryTestCodec{name: "jpeg", syntax: transfer.JPEGBaseline8Bit},
	} {
		if err := registry.Register(registered); err != nil {
			t.Fatal(err)
		}
	}

	listed := registry.List()
	uids := make([]string, len(listed))
	for index, registered := range listed {
		uids[index] = registered.TransferSyntax().UID().UID()
	}
	want := []string{
		transfer.ImplicitVRLittleEndian.UID().UID(),
		transfer.ExplicitVRLittleEndian.UID().UID(),
		transfer.ExplicitVRBigEndian.UID().UID(),
		transfer.JPEGBaseline8Bit.UID().UID(),
		transfer.RLELossless.UID().UID(),
	}
	if !reflect.DeepEqual(uids, want) {
		t.Fatalf("List() UIDs = %v, want sorted %v", uids, want)
	}
	listed[0] = nil
	if got := registry.List()[0]; got == nil {
		t.Fatal("mutating List result changed registry contents")
	}
}

func TestRegistryInstancesAreIsolated(t *testing.T) {
	first := NewRegistry()
	second := NewRegistry()
	if err := first.Register(registryTestCodec{name: "jpeg2000", syntax: transfer.JPEG2000Lossless}); err != nil {
		t.Fatal(err)
	}
	if _, found := second.Lookup(transfer.JPEG2000Lossless); found {
		t.Fatal("registration leaked into another Registry")
	}
	if first == GlobalRegistry() || second == GlobalRegistry() {
		t.Fatal("isolated Registry reused GlobalRegistry")
	}
}

func TestRegistrySupportsConcurrentRuntimeMutation(t *testing.T) {
	registry := NewRegistry()
	const workers = 32
	errorsCh := make(chan error, workers)
	var wg sync.WaitGroup
	for index := range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			uidValue := fmt.Sprintf("1.2.826.0.1.3680043.10.854.900.%d", index)
			syntax := transfer.New(uid.New(uidValue, "Runtime Test", uid.TypeTransferSyntax, false))
			first := registryTestCodec{name: registryFirstCodecName, syntax: syntax}
			if err := registry.Register(first); err != nil {
				errorsCh <- err
				return
			}
			if found, ok := registry.Lookup(syntax); !ok || found.Name() != registryFirstCodecName {
				errorsCh <- fmt.Errorf("Lookup(%s) after Register = (%#v, %v)", uidValue, found, ok)
				return
			}
			second := registryTestCodec{name: registrySecondCodecName, syntax: syntax}
			previous, err := registry.Replace(second)
			if err != nil || previous == nil || previous.Name() != registryFirstCodecName {
				errorsCh <- fmt.Errorf("Replace(%s) = (%#v, %v)", uidValue, previous, err)
				return
			}
			_ = registry.List()
			removed, ok := registry.Unregister(syntax)
			if !ok || removed == nil || removed.Name() != registrySecondCodecName {
				errorsCh <- fmt.Errorf("Unregister(%s) = (%#v, %v)", uidValue, removed, ok)
			}
		}()
	}
	wg.Wait()
	close(errorsCh)
	for err := range errorsCh {
		t.Error(err)
	}
}
