// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package element_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

// DeepClone must preserve a BulkDataURI buffer's reference; the old
// implementation flattened every buffer through Data() and turned a URI-only
// buffer into an empty in-memory buffer.
func TestDeepClonePreservesBulkDataURIBuffer(t *testing.T) {
	const uri = "https://example.com/bulk/b1"
	source := element.NewOtherByteFromBuffer(tag.PixelData, buffer.NewBulkDataURI(uri))

	clonedElem := element.DeepClone(source)
	clone, ok := clonedElem.(*element.OtherByte)
	if !ok {
		t.Fatalf("DeepClone() type = %T, want *element.OtherByte", clonedElem)
	}
	cloned, ok := clone.Buffer().(*buffer.BulkDataURIByteBuffer)
	if !ok {
		t.Fatalf("cloned buffer = %T, want *buffer.BulkDataURIByteBuffer", clone.Buffer())
	}
	if cloned.BulkDataURI() != uri {
		t.Fatalf("cloned BulkDataURI = %q, want %q", cloned.BulkDataURI(), uri)
	}
}
