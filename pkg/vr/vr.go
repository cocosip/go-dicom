package vr

import (
	"errors"
	"fmt"
	"go-dicom/pkg/tag"
	"reflect"
	"time"
)

//DICOM Value Representation
type DicomVR struct {
	//Code used to represent VR.
	Code string
	//Descriptive name of VR.
	Name string
	//Value is a string.
	IsString bool
	//String is encoded using the specified character set.
	IsStringEncoded bool
	//Length field of value is a 16-bit short integer.
	Is16bitLength bool
	//Value can contain multiple items.
	IsMultiValue bool
	//Byte value used to pad value to even length.
	PaddingValue byte
	//Maximum length of a single value.
	MaximumLength uint
	//Size of each individual value unit for fixed length value types.
	UnitSize int
	//Number of bytes to swap when changing endian representation of value. Usually equal to the "UnitSize"
	ByteSwap int
	//Type used to represent VR value.
	ValueType reflect.Type

	stringValidator func(content string) error
}

//Code String for DICOM Value Representation
const (
	CodeAE string = "AE"
    //Age String
	CodeAS string = "AS"
    //Attribute Tag
	CodeAT string = "AT"
	//Code String
	CodeCS string = "CS"
	//Date
	CodeDA string = "DA"
	//Decimal String
	CodeDS string = "DS"
	//Date Time
	CodeDT string = "DT"
	//Floating Point Double
	CodeFD string = "FD"
	//Floating Point Single
	CodeFL string = "FL"
	//Integer String
	CodeIS string = "IS"
	//Long String
	CodeLO string = "LO"
	//Long Text
	CodeLT string = "LT"
	//Other Byte
	CodeOB string = "OB"
	//Other Double
	CodeOD string = "OD"
	//Other Float
	CodeOF string = "OF"
	//Other Long
	CodeOL string = "OL"
	//Other Very Long
	CodeOV string = "OV"
	//Other Word
	CodeOW string = "OW"
	//Person Name
	CodePN string = "PN"
	//Short String
	CodeSH string = "SH"
	//Signed Long
	CodeSL string = "SL"
	//Sequence of Items
	CodeSQ string = "SQ"
	//Signed Short
	CodeSS string = "SS"
	//Short Text
	CodeST string = "ST"
	//Signed Very Long
	CodeSV string = "SV"
	//Time
	CodeTM string = "TM"
	//Unlimited Characters
	CodeUC string = "UC"
	//Unique Identifier
	CodeUI string = "UI"
	//Unsigned Long
	CodeUL string = "UL"
	//Unknown
	CodeUN string = "UN"
	//Universal Resource Identifier or Universal Resource Locator (URI/URL)
	CodeUR string = "UR"
	//Unsigned Short
	CodeUS string = "US"
	//Unlimited Text
	CodeUT string = "UT"
	//Unsigned Very Long
	CodeUV string = "UV"
)

const (
	PadZero byte = 0x00
	PadSpace byte =0x20
)


//Implicit VR in Explicit VR context
var Implicit DicomVR = DicomVR{}

//No VR
var NONE DicomVR = DicomVR{
	Code:            "NONE",
	Name:            "No Value Representation",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   false,
	IsMultiValue:    false,
	PaddingValue:    PadZero,
	MaximumLength:   0,
	UnitSize:        0,
	ByteSwap:        0,
	ValueType:       reflect.TypeOf(reflect.Interface),
}

//Application Entity
var AE DicomVR = DicomVR{
	Code:            CodeAE,
	Name:            "Application Entity",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   16,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidateAE,
}

//Age String
var AS DicomVR = DicomVR{
	Code:            CodeAS,
	Name:            "Age String",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   4,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidateAS,
}

//Attribute Tag
var AT DicomVR = DicomVR{
	Code:            CodeAT,
	Name:            "Attribute Tag",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   4,
	UnitSize:        4,
	ByteSwap:        2,
	ValueType:       reflect.TypeOf(tag.DicomTag{}),
}

//Code String
var CS DicomVR = DicomVR{
	Code:            CodeCS,
	Name:            "Code String",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   16,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidateCS,
}

//Date
var DA DicomVR = DicomVR{
	Code:            CodeDA,
	Name:            "Date",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   18,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(time.Time{}),
	stringValidator: ValidateDA,
}

//Decimal String
var DS DicomVR = DicomVR{
	Code:            CodeDS,
	Name:            "Decimal String",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   16,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.Float32),
	stringValidator: ValidateDS,
}

//Date Time
var DT DicomVR = DicomVR{
	Code:            CodeDT,
	Name:            "Date Time",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   54,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(time.Time{}),
	stringValidator: ValidateDT,
}

//Floating Point Double
var FD DicomVR = DicomVR{
	Code:            CodeFD,
	Name:            "Floating Point Double",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   8,
	UnitSize:        8,
	ByteSwap:        8,
	ValueType:       reflect.TypeOf(reflect.Float32),
}

//Floating Point Single
var FL DicomVR = DicomVR{
	Code:            CodeFL,
	Name:            "Floating Point Single",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   4,
	UnitSize:        4,
	ByteSwap:        4,
	ValueType:       reflect.TypeOf(reflect.Float32),
}

//Integer String
var IS DicomVR = DicomVR{
	Code:            CodeIS,
	Name:            "Integer String",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   12,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.Int),
	stringValidator: ValidateIS,
}

//Long String
var LO DicomVR = DicomVR{
	Code:            CodeLO,
	Name:            "Long String",
	IsString:        true,
	IsStringEncoded: true,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   64,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidateLO,
}

//Long Text
var LT DicomVR = DicomVR{
	Code:            CodeLT,
	Name:            "Long Text",
	IsString:        true,
	IsStringEncoded: true,
	Is16bitLength:   true,
	IsMultiValue:    false,
	PaddingValue:    PadSpace,
	MaximumLength:   10240,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidateLT,
}

//Other Byte
var OB DicomVR = DicomVR{
	Code:            CodeOB,
	Name:            "Other Byte",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   false,
	IsMultiValue:    true,
	MaximumLength:   0,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf([]byte{}),
}

//Other Double
var OD DicomVR = DicomVR{
	Code:            CodeOD,
	Name:            "Other Double",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   false,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   0,
	UnitSize:        8,
	ByteSwap:        8,
	ValueType:       reflect.TypeOf([]float32{}),
}

//Other Float
var OF DicomVR = DicomVR{
	Code:            CodeOF,
	Name:            "Other Float",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   false,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   0,
	UnitSize:        4,
	ByteSwap:        4,
	ValueType:       reflect.TypeOf([]float32{}),
}

//Other Long
var OL DicomVR = DicomVR{
	Code:            CodeOL,
	Name:            "Other Long",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   false,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   0,
	UnitSize:        4,
	ByteSwap:        4,
	ValueType:       reflect.TypeOf([]uint{}),
}

//Other Word
var OW DicomVR = DicomVR{
	Code:            CodeOW,
	Name:            "Other Word",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   false,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   0,
	UnitSize:        2,
	ByteSwap:        2,
	ValueType:       reflect.TypeOf([]uint16{}),
}

//Other Very Long
var OV DicomVR = DicomVR{
	Code:            CodeOV,
	Name:            "Other Very Long",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   false,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   0,
	UnitSize:        8,
	ByteSwap:        8,
	ValueType:       reflect.TypeOf([]uint64{}),
}

//Person Name
var PN DicomVR = DicomVR{
	Code:            CodePN,
	Name:            "Person Name",
	IsString:        true,
	IsStringEncoded: true,
	Is16bitLength:   true,
	PaddingValue:    PadSpace,
	MaximumLength:   64,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidatePN,
}

//Short String
var SH DicomVR = DicomVR{
	Code:            CodeSH,
	Name:            "Short String",
	IsString:        true,
	IsStringEncoded: true,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   16,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidateSH,
}

//Signed Long
var SL DicomVR = DicomVR{
	Code:            CodeSL,
	Name:            "Signed Long",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   4,
	UnitSize:        4,
	ByteSwap:        4,
	ValueType:       reflect.TypeOf(reflect.Int),
}

//Sequence of Items
var SQ DicomVR = DicomVR{
	Code:            CodeSQ,
	Name:            "Sequence of Items",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   false,
	IsMultiValue:    false,
	PaddingValue:    PadSpace,
	MaximumLength:   0,
	UnitSize:        0,
	ByteSwap:        0,
}

//Signed Short
var SS DicomVR = DicomVR{
	Code:            CodeSS,
	Name:            "Signed Short",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   2,
	UnitSize:        2,
	ByteSwap:        2,
	ValueType:       reflect.TypeOf(reflect.Int16),
}

//Short Text
var ST DicomVR = DicomVR{
	Code:            CodeST,
	Name:            "Short Text",
	IsString:        true,
	IsStringEncoded: true,
	Is16bitLength:   true,
	IsMultiValue:    false,
	PaddingValue:    PadSpace,
	MaximumLength:   1024,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidateST,
}

//Signed Very Long
var SV DicomVR = DicomVR{
	Code:            CodeSV,
	Name:            "Signed Very Long",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   8,
	UnitSize:        8,
	ByteSwap:        8,
	ValueType:       reflect.TypeOf(reflect.Int64),
}

//Time
var TM DicomVR = DicomVR{
	Code:            CodeTM,
	Name:            "Time",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   16,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(time.Time{}),
	stringValidator: ValidateTM,
}

//Unlimited Characters
var UC DicomVR = DicomVR{
	Code:            CodeUC,
	Name:            "Unlimited Characters",
	IsString:        true,
	IsStringEncoded: true,
	Is16bitLength:   false,
	IsMultiValue:    true,
	PaddingValue:    PadSpace,
	MaximumLength:   0,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
}

//Unique Identifier
var UI DicomVR = DicomVR{
	Code:            CodeUI,
	Name:            "Unique Identifier",
	IsString:        true,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   64,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
	stringValidator: ValidateUI,
}

//Unsigned Long
var UL DicomVR = DicomVR{
	Code:            CodeUL,
	Name:            "Unsigned Long",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   4,
	UnitSize:        4,
	ByteSwap:        4,
	ValueType:       reflect.TypeOf(reflect.Uint),
}

//Unknown
var UN DicomVR = DicomVR{
	Code:            CodeUN,
	Name:            "Unknown",
	IsString:        false,
	IsStringEncoded: false,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   0,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf([]byte{}),
}

//Universal Resource Identifier or Universal Resource Locator (URI/URL)
var UR DicomVR = DicomVR{
	Code:            CodeUR,
	Name:            "Universal Resource Identifier or Locator",
	IsString:        true,
	IsStringEncoded: true,
	Is16bitLength:   false,
	IsMultiValue:    false,
	PaddingValue:    PadSpace,
	MaximumLength:   0,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
}

//Unsigned Short
var US DicomVR = DicomVR{
	Code:            CodeUS,
	Name:            "Unsigned Short",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   2,
	UnitSize:        2,
	ByteSwap:        2,
	ValueType:       reflect.TypeOf(reflect.Uint16),
}

//Unlimited Text
var UT DicomVR = DicomVR{
	Code:            CodeUT,
	Name:            "Unlimited Text",
	IsString:        true,
	IsStringEncoded: true,
	Is16bitLength:   false,
	IsMultiValue:    false,
	PaddingValue:    PadSpace,
	MaximumLength:   0,
	UnitSize:        1,
	ByteSwap:        1,
	ValueType:       reflect.TypeOf(reflect.String),
}

//Unsigned Very Long
var UV DicomVR = DicomVR{
	Code:            CodeUV,
	Name:            "Unsigned Very Long",
	IsString:        false,
	IsStringEncoded: false,
	Is16bitLength:   true,
	IsMultiValue:    true,
	PaddingValue:    PadZero,
	MaximumLength:   8,
	UnitSize:        8,
	ByteSwap:        8,
	ValueType:       reflect.TypeOf(reflect.Uint64),
}

//validates a string content. Throws Dicom
func (vr DicomVR) ValidateString(content string) {
	if vr.stringValidator != nil {
		vr.stringValidator(content)
	}
}

//Gets a string representation of this VR.
func (vr DicomVR) ToString() string {
	return vr.Code
}

//Get VR for given string value.
func Parse(vr string) (DicomVR,error) {
	dicomVR, valid := TryParse(vr)
	if !valid {
		return DicomVR{}, errors.New(fmt.Sprintf("Unknown VR: '%s'", vr))
	}
	return dicomVR, nil
}

// Try to get VR for given string value.
func TryParse(vr string)  (DicomVR,bool) {
	result, valid := tryParse(vr, false)
	return result, valid
}

func tryParse(vr string,valid bool) (DicomVR,bool) {
	valid = true
	switch vr {
	case "NONE":
		return NONE, valid
	case CodeAE:
		return AE, valid
	case CodeAS:
		return AS, valid
	case CodeAT:
		return AT, valid
	case CodeCS:
		return CS, valid
	case CodeDA:
		return DA, valid
	case CodeDS:
		return DS, valid
	case CodeDT:
		return DT, valid
	case CodeFD:
		return FD, valid
	case CodeFL:
		return FL, valid
	case CodeIS:
		return IS, valid
	case CodeLO:
		return LO, valid
	case CodeLT:
		return LT, valid
	case CodeOB:
		return OB, valid
	case CodeOD:
		return OD, valid
	case CodeOF:
		return OF, valid
	case CodeOL:
		return OL, valid
	case CodeOW:
		return OW, valid
	case CodeOV:
		return OV, valid
	case CodePN:
		return PN, valid
	case CodeSH:
		return SH, valid
	case CodeSL:
		return SL, valid
	case CodeSQ:
		return SQ, valid
	case CodeSS:
		return SS, valid
	case CodeST:
		return ST, valid
	case CodeSV:
		return SV, valid
	case CodeTM:
		return TM, valid
	case CodeUC:
		return UC, valid
	case CodeUI:
		return UI, valid
	case CodeUL:
		return UL, valid
	case CodeUN:
		return UN, valid
	case CodeUR:
		return UR, valid
	case CodeUS:
		return US, valid
	case CodeUT:
		return UT, valid
	case CodeUV:
		return UV, valid
	default:
		valid = false
		return NONE, valid
	}
}


