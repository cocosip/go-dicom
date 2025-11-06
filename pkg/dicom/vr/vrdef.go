// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package vr

import "reflect"

// Standard VR instances
var (
	// None represents no VR.
	None = &VR{
		code:            "NONE",
		name:            "No Value Representation",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    false,
		paddingValue:    padZero,
		maximumLength:   0,
		unitSize:        0,
		byteSwap:        0,
		valueType:       reflect.TypeOf((*interface{})(nil)).Elem(),
	}

	// AE - Application Entity
	AE = &VR{
		code:            "AE",
		name:            "Application Entity",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   16,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateAE,
	}

	// AS - Age String
	AS = &VR{
		code:            "AS",
		name:            "Age String",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   4,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateAS,
	}

	// AT - Attribute Tag
	AT = &VR{
		code:            "AT",
		name:            "Attribute Tag",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   4,
		unitSize:        4,
		byteSwap:        2,
		valueType:       reflect.TypeOf(uint32(0)),
	}

	// CS - Code String
	CS = &VR{
		code:            "CS",
		name:            "Code String",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   16,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateCS,
	}

	// DA - Date
	DA = &VR{
		code:            "DA",
		name:            "Date",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   18,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateDA,
	}

	// DS - Decimal String
	DS = &VR{
		code:            "DS",
		name:            "Decimal String",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   16,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateDS,
	}

	// DT - Date Time
	DT = &VR{
		code:            "DT",
		name:            "Date Time",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   54,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateDT,
	}

	// FD - Floating Point Double
	FD = &VR{
		code:            "FD",
		name:            "Floating Point Double",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   8,
		unitSize:        8,
		byteSwap:        8,
		valueType:       reflect.TypeOf(float64(0)),
	}

	// FL - Floating Point Single
	FL = &VR{
		code:            "FL",
		name:            "Floating Point Single",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   4,
		unitSize:        4,
		byteSwap:        4,
		valueType:       reflect.TypeOf(float32(0)),
	}

	// IS - Integer String
	IS = &VR{
		code:            "IS",
		name:            "Integer String",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   12,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(int32(0)),
		stringValidator: ValidateIS,
	}

	// LO - Long String
	LO = &VR{
		code:            "LO",
		name:            "Long String",
		isString:        true,
		isStringEncoded: true,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   64,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateLO,
	}

	// LT - Long Text
	LT = &VR{
		code:            "LT",
		name:            "Long Text",
		isString:        true,
		isStringEncoded: true,
		is16bitLength:   true,
		isMultiValue:    false,
		paddingValue:    padSpace,
		maximumLength:   10240,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateLT,
	}

	// OB - Other Byte
	OB = &VR{
		code:            "OB",
		name:            "Other Byte",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   0,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf([]byte(nil)),
	}

	// OD - Other Double
	OD = &VR{
		code:            "OD",
		name:            "Other Double",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   0,
		unitSize:        8,
		byteSwap:        8,
		valueType:       reflect.TypeOf([]float64(nil)),
	}

	// OF - Other Float
	OF = &VR{
		code:            "OF",
		name:            "Other Float",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   0,
		unitSize:        4,
		byteSwap:        4,
		valueType:       reflect.TypeOf([]float32(nil)),
	}

	// OL - Other Long
	OL = &VR{
		code:            "OL",
		name:            "Other Long",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   0,
		unitSize:        4,
		byteSwap:        4,
		valueType:       reflect.TypeOf([]uint32(nil)),
	}

	// OW - Other Word
	OW = &VR{
		code:            "OW",
		name:            "Other Word",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   0,
		unitSize:        2,
		byteSwap:        2,
		valueType:       reflect.TypeOf([]uint16(nil)),
	}

	// OV - Other Very Long
	OV = &VR{
		code:            "OV",
		name:            "Other Very Long",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   0,
		unitSize:        8,
		byteSwap:        8,
		valueType:       reflect.TypeOf([]uint64(nil)),
	}

	// PN - Person Name
	PN = &VR{
		code:            "PN",
		name:            "Person Name",
		isString:        true,
		isStringEncoded: true,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   64,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidatePN,
	}

	// SH - Short String
	SH = &VR{
		code:            "SH",
		name:            "Short String",
		isString:        true,
		isStringEncoded: true,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   16,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateSH,
	}

	// SL - Signed Long
	SL = &VR{
		code:            "SL",
		name:            "Signed Long",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   4,
		unitSize:        4,
		byteSwap:        4,
		valueType:       reflect.TypeOf(int32(0)),
	}

	// SQ - Sequence of Items
	SQ = &VR{
		code:            "SQ",
		name:            "Sequence of Items",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    false,
		paddingValue:    padSpace,
		maximumLength:   0,
		unitSize:        0,
		byteSwap:        0,
		valueType:       reflect.TypeOf((*interface{})(nil)).Elem(),
	}

	// SS - Signed Short
	SS = &VR{
		code:            "SS",
		name:            "Signed Short",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   2,
		unitSize:        2,
		byteSwap:        2,
		valueType:       reflect.TypeOf(int16(0)),
	}

	// ST - Short Text
	ST = &VR{
		code:            "ST",
		name:            "Short Text",
		isString:        true,
		isStringEncoded: true,
		is16bitLength:   true,
		isMultiValue:    false,
		paddingValue:    padSpace,
		maximumLength:   1024,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateST,
	}

	// SV - Signed Very Long
	SV = &VR{
		code:            "SV",
		name:            "Signed Very Long",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   8,
		unitSize:        8,
		byteSwap:        8,
		valueType:       reflect.TypeOf(int64(0)),
	}

	// TM - Time
	TM = &VR{
		code:            "TM",
		name:            "Time",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   16,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateTM,
	}

	// UC - Unlimited Characters
	UC = &VR{
		code:            "UC",
		name:            "Unlimited Characters",
		isString:        true,
		isStringEncoded: true,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padSpace,
		maximumLength:   0,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateUC,
	}

	// UI - Unique Identifier
	UI = &VR{
		code:            "UI",
		name:            "Unique Identifier",
		isString:        true,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   64,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateUI,
	}

	// UL - Unsigned Long
	UL = &VR{
		code:            "UL",
		name:            "Unsigned Long",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   4,
		unitSize:        4,
		byteSwap:        4,
		valueType:       reflect.TypeOf(uint32(0)),
	}

	// UN - Unknown
	UN = &VR{
		code:            "UN",
		name:            "Unknown",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   0,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf([]byte(nil)),
	}

	// UR - Universal Resource Identifier or Locator
	UR = &VR{
		code:            "UR",
		name:            "Universal Resource Identifier or Locator",
		isString:        true,
		isStringEncoded: true,
		is16bitLength:   false,
		isMultiValue:    false,
		paddingValue:    padSpace,
		maximumLength:   0,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateUR,
	}

	// US - Unsigned Short
	US = &VR{
		code:            "US",
		name:            "Unsigned Short",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   true,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   2,
		unitSize:        2,
		byteSwap:        2,
		valueType:       reflect.TypeOf(uint16(0)),
	}

	// UT - Unlimited Text
	UT = &VR{
		code:            "UT",
		name:            "Unlimited Text",
		isString:        true,
		isStringEncoded: true,
		is16bitLength:   false,
		isMultiValue:    false,
		paddingValue:    padSpace,
		maximumLength:   0,
		unitSize:        1,
		byteSwap:        1,
		valueType:       reflect.TypeOf(""),
		stringValidator: ValidateUT,
	}

	// UV - Unsigned Very Long
	UV = &VR{
		code:            "UV",
		name:            "Unsigned Very Long",
		isString:        false,
		isStringEncoded: false,
		is16bitLength:   false,
		isMultiValue:    true,
		paddingValue:    padZero,
		maximumLength:   8,
		unitSize:        8,
		byteSwap:        8,
		valueType:       reflect.TypeOf(uint64(0)),
	}
)

// registry maps VR code strings to VR instances
var registry = map[string]*VR{
	"NONE": None,
	"AE":   AE,
	"AS":   AS,
	"AT":   AT,
	"CS":   CS,
	"DA":   DA,
	"DS":   DS,
	"DT":   DT,
	"FD":   FD,
	"FL":   FL,
	"IS":   IS,
	"LO":   LO,
	"LT":   LT,
	"OB":   OB,
	"OD":   OD,
	"OF":   OF,
	"OL":   OL,
	"OV":   OV,
	"OW":   OW,
	"PN":   PN,
	"SH":   SH,
	"SL":   SL,
	"SQ":   SQ,
	"SS":   SS,
	"ST":   ST,
	"SV":   SV,
	"TM":   TM,
	"UC":   UC,
	"UI":   UI,
	"UL":   UL,
	"UN":   UN,
	"UR":   UR,
	"US":   US,
	"UT":   UT,
	"UV":   UV,
}
