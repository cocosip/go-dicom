// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package serialization

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
	"github.com/cocosip/go-dicom/pkg/io/buffer"
)

const testBulkDataURI = "http://example.com/dicom/pixeldata"

func TestToJSON_BulkDataURI(t *testing.T) {
	// Create dataset with BulkDataURI element
	ds := dataset.New()

	// Create a BulkDataUriByteBuffer
	uri := testBulkDataURI
	buf := buffer.NewBulkDataURI(uri)

	// Create OW element with BulkDataURI
	pixelDataTag := tag.New(0x7FE0, 0x0010) // PixelData
	elem := element.NewOtherWordFromBuffer(pixelDataTag, buf)
	ds.Add(elem)

	// Serialize to JSON
	jsonData, err := ToJSON(ds)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}

	// Parse JSON to verify structure
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Verify PixelData tag exists
	pixelDataKey := "7FE00010"
	if _, ok := result[pixelDataKey]; !ok {
		t.Errorf("Expected tag %s not found in JSON", pixelDataKey)
		return
	}

	// Verify BulkDataURI field
	tagData, ok := result[pixelDataKey].(map[string]interface{})
	if !ok {
		t.Errorf("Tag data is not a map")
		return
	}

	if tagData["vr"] != "OW" {
		t.Errorf("VR = %v, want OW", tagData["vr"])
	}

	if tagData["BulkDataURI"] != uri {
		t.Errorf("BulkDataURI = %v, want %s", tagData["BulkDataURI"], uri)
	}

	// Should not have InlineBinary
	if _, ok := tagData["InlineBinary"]; ok {
		t.Error("BulkDataURI element should not have InlineBinary field")
	}
}

func TestToJSON_BulkDataURI_WithData(t *testing.T) {
	// Create dataset with BulkDataURI element that has data loaded
	ds := dataset.New()

	uri := testBulkDataURI
	data := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
	buf := buffer.NewBulkDataURIWithData(uri, data)

	pixelDataTag := tag.New(0x7FE0, 0x0010)
	elem := element.NewOtherByteFromBuffer(pixelDataTag, buf)
	ds.Add(elem)

	// Serialize to JSON - should still use BulkDataURI, not InlineBinary
	jsonData, err := ToJSON(ds)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	tagData := result["7FE00010"].(map[string]interface{})
	if tagData["BulkDataURI"] != uri {
		t.Errorf("BulkDataURI = %v, want %s", tagData["BulkDataURI"], uri)
	}

	// Should not have InlineBinary even though data is loaded
	if _, ok := tagData["InlineBinary"]; ok {
		t.Error("BulkDataURI element should use BulkDataURI, not InlineBinary")
	}
}

func TestFromJSON_BulkDataURI(t *testing.T) {
	// Create JSON with BulkDataURI
	jsonStr := `{
		"7FE00010": {
			"vr": "OW",
			"BulkDataURI": "http://example.com/dicom/pixeldata"
		}
	}`

	// Deserialize from JSON
	ds, err := FromJSON([]byte(jsonStr))
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}

	// Get PixelData element
	pixelDataTag := tag.New(0x7FE0, 0x0010)
	elem, found := ds.Get(pixelDataTag)
	if !found {
		t.Fatal("PixelData element not found in dataset")
	}

	// Verify VR
	if elem.ValueRepresentation().Code() != vr.CodeOW {
		t.Errorf("VR = %s, want OW", elem.ValueRepresentation().Code())
	}

	// Verify buffer is BulkDataUriByteBuffer
	buf := elem.Buffer()
	bulkDataBuf, ok := buf.(*buffer.BulkDataURIByteBuffer)
	if !ok {
		t.Fatalf("Buffer is not BulkDataUriByteBuffer, got %T", buf)
	}

	// Verify URI
	expectedURI := "http://example.com/dicom/pixeldata"
	if bulkDataBuf.BulkDataURI() != expectedURI {
		t.Errorf("BulkDataURI() = %s, want %s", bulkDataBuf.BulkDataURI(), expectedURI)
	}

	// Verify IsBulkDataUri
	if !bulkDataBuf.IsBulkDataURI() {
		t.Error("IsBulkDataUri() = false, want true")
	}

	// Data should not be loaded yet
	if bulkDataBuf.IsMemory() {
		t.Error("IsMemory() = true, want false (data not loaded yet)")
	}
}

func TestFromJSON_BulkDataURI_AllVRTypes(t *testing.T) {
	tests := []struct {
		name   string
		vr     string
		vrCode string
	}{
		{"OtherByte", "OB", vr.CodeOB},
		{"OtherWord", "OW", vr.CodeOW},
		{"OtherDouble", "OD", vr.CodeOD},
		{"OtherFloat", "OF", vr.CodeOF},
		{"OtherLong", "OL", vr.CodeOL},
		{"OtherVeryLong", "OV", vr.CodeOV},
		{"Unknown", "UN", vr.CodeUN},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonStr := `{
				"00091010": {
					"vr": "` + tt.vr + `",
					"BulkDataURI": "http://example.com/data/bulk` + tt.vr + `"
				}
			}`

			ds, err := FromJSON([]byte(jsonStr))
			if err != nil {
				t.Fatalf("FromJSON() error = %v", err)
			}

			testTag := tag.New(0x0009, 0x1010)
			elem, found := ds.Get(testTag)
			if !found {
				t.Fatal("Element not found in dataset")
			}

			if elem.ValueRepresentation().Code() != tt.vrCode {
				t.Errorf("VR = %s, want %s", elem.ValueRepresentation().Code(), tt.vrCode)
			}

			buf := elem.Buffer()
			bulkDataBuf, ok := buf.(*buffer.BulkDataURIByteBuffer)
			if !ok {
				t.Fatalf("Buffer is not BulkDataUriByteBuffer, got %T", buf)
			}

			expectedURI := "http://example.com/data/bulk" + tt.vr
			if bulkDataBuf.BulkDataURI() != expectedURI {
				t.Errorf("BulkDataURI() = %s, want %s", bulkDataBuf.BulkDataURI(), expectedURI)
			}
		})
	}
}

func TestJSON_BulkDataURI_Roundtrip(t *testing.T) {
	// Create original dataset
	ds1 := dataset.New()

	// Add regular element
	patientName := tag.New(0x0010, 0x0010)
	pnVR, _ := vr.Parse(vr.CodePN)
	ds1.Add(element.NewString(patientName, pnVR, []string{"Doe^John"}))

	// Add BulkDataURI element
	uri := "http://example.com/dicom/waveform"
	buf := buffer.NewBulkDataURI(uri)
	waveformTag := tag.New(0x5400, 0x1010) // WaveformData
	ds1.Add(element.NewOtherWordFromBuffer(waveformTag, buf))

	// Serialize to JSON
	jsonData, err := ToJSON(ds1)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}

	// Deserialize from JSON
	ds2, err := FromJSON(jsonData)
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}

	// Verify regular element
	elem1, found1 := ds2.Get(patientName)
	if !found1 {
		t.Fatal("PatientName element not found after roundtrip")
	}
	strElem, ok := elem1.(*element.String)
	if !ok {
		t.Fatalf("PatientName is not String element, got %T", elem1)
	}
	values := strElem.GetValues()
	if len(values) != 1 || values[0] != "Doe^John" {
		t.Errorf("PatientName = %v, want [Doe^John]", values)
	}

	// Verify BulkDataURI element
	elem2, found2 := ds2.Get(waveformTag)
	if !found2 {
		t.Fatal("WaveformData element not found after roundtrip")
	}
	bulkDataBuf, ok := elem2.Buffer().(*buffer.BulkDataURIByteBuffer)
	if !ok {
		t.Fatalf("WaveformData buffer is not BulkDataUriByteBuffer, got %T", elem2.Buffer())
	}
	if bulkDataBuf.BulkDataURI() != uri {
		t.Errorf("BulkDataURI() = %s, want %s", bulkDataBuf.BulkDataURI(), uri)
	}
}

func TestToJSON_MixedBulkDataAndInline(t *testing.T) {
	ds := dataset.New()

	// Add InlineBinary element (normal OB with data)
	tag1 := tag.New(0x0009, 0x1001)
	data1 := []byte{0xFF, 0xD8, 0xFF, 0xE0} // Small data, inline
	buf1 := buffer.NewMemory(data1)
	ds.Add(element.NewOtherByteFromBuffer(tag1, buf1))

	// Add BulkDataURI element
	tag2 := tag.New(0x0009, 0x1002)
	uri := "http://example.com/dicom/large"
	buf2 := buffer.NewBulkDataURI(uri)
	ds.Add(element.NewOtherByteFromBuffer(tag2, buf2))

	// Serialize to JSON
	jsonData, err := ToJSON(ds)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Verify first element has InlineBinary
	tag1Data := result["00091001"].(map[string]interface{})
	if _, ok := tag1Data["InlineBinary"]; !ok {
		t.Error("First element should have InlineBinary")
	}
	if _, ok := tag1Data["BulkDataURI"]; ok {
		t.Error("First element should not have BulkDataURI")
	}

	// Verify second element has BulkDataURI
	tag2Data := result["00091002"].(map[string]interface{})
	if _, ok := tag2Data["BulkDataURI"]; !ok {
		t.Error("Second element should have BulkDataURI")
	}
	if _, ok := tag2Data["InlineBinary"]; ok {
		t.Error("Second element should not have InlineBinary")
	}
	if tag2Data["BulkDataURI"] != uri {
		t.Errorf("BulkDataURI = %v, want %s", tag2Data["BulkDataURI"], uri)
	}
}

func TestFromJSON_InvalidBulkDataURI(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		errMsg  string
	}{
		{
			name: "BulkDataURI not string",
			jsonStr: `{
				"7FE00010": {
					"vr": "OW",
					"BulkDataURI": 12345
				}
			}`,
			errMsg: "invalid BulkDataURI",
		},
		{
			name: "BulkDataURI null",
			jsonStr: `{
				"7FE00010": {
					"vr": "OW",
					"BulkDataURI": null
				}
			}`,
			errMsg: "BulkDataURI cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FromJSON([]byte(tt.jsonStr))
			if err == nil {
				t.Error("FromJSON() expected error, got nil")
				return
			}
			if !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("FromJSON() error = %v, want substring %s", err, tt.errMsg)
			}
		})
	}
}

func TestBulkDataUriByteBuffer_SimulatedLazyLoad(t *testing.T) {
	// This test simulates the full lifecycle:
	// 1. Deserialize JSON with BulkDataURI
	// 2. Fetch data from URI (simulated)
	// 3. Access the data

	jsonStr := `{
		"7FE00010": {
			"vr": "OB",
			"BulkDataURI": "http://example.com/dicom/pixeldata"
		}
	}`

	// Step 1: Deserialize
	ds, err := FromJSON([]byte(jsonStr))
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}

	pixelDataTag := tag.New(0x7FE0, 0x0010)
	elem, found := ds.Get(pixelDataTag)
	if !found {
		t.Fatal("PixelData element not found")
	}

	bulkDataBuf, ok := elem.Buffer().(*buffer.BulkDataURIByteBuffer)
	if !ok {
		t.Fatalf("Buffer is not BulkDataUriByteBuffer")
	}

	// Initially no data
	if bulkDataBuf.IsMemory() {
		t.Error("Buffer should not have data initially")
	}

	// Step 2: Simulate fetching data from URI
	// In real implementation, this would be an HTTP request
	fetchedData := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	bulkDataBuf.SetData(fetchedData)

	// Step 3: Now data is available
	if !bulkDataBuf.IsMemory() {
		t.Error("Buffer should have data after SetData")
	}

	if bulkDataBuf.Size() != 8 {
		t.Errorf("Size() = %d, want 8", bulkDataBuf.Size())
	}

	// Verify data can be retrieved
	output := make([]byte, 4)
	if err := bulkDataBuf.GetByteRange(0, 4, output); err != nil {
		t.Errorf("GetByteRange() error = %v", err)
	}

	expected := []byte{0x01, 0x02, 0x03, 0x04}
	for i, b := range output {
		if b != expected[i] {
			t.Errorf("GetByteRange() output[%d] = %d, want %d", i, b, expected[i])
		}
	}
}
