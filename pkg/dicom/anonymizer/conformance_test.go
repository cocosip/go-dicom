// Copyright (c) 2025 go-dicom contributors.
// Licensed under the MIT License.

package anonymizer

import (
	"strings"
	"testing"
	"time"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
	"github.com/cocosip/go-dicom/pkg/dicom/element"
	"github.com/cocosip/go-dicom/pkg/dicom/tag"
	"github.com/cocosip/go-dicom/pkg/dicom/vr"
)

// nonKBasicTags lists every (GGGG,EEEE) row of PS3.15 Table E.1-1 whose Basic
// Profile action is not "K". The embedded default profile must match all of
// them (exact row or a group pattern).
var nonKBasicTags = []string{
	"0000,1000",
	"0000,1001",
	"0002,0003",
	"0004,1511",
	"0008,0012",
	"0008,0013",
	"0008,0014",
	"0008,0015",
	"0008,0017",
	"0008,0018",
	"0008,0019",
	"0008,0020",
	"0008,0021",
	"0008,0022",
	"0008,0023",
	"0008,0024",
	"0008,0025",
	"0008,002A",
	"0008,0030",
	"0008,0031",
	"0008,0032",
	"0008,0033",
	"0008,0034",
	"0008,0035",
	"0008,0050",
	"0008,0054",
	"0008,0055",
	"0008,0058",
	"0008,0080",
	"0008,0081",
	"0008,0082",
	"0008,0090",
	"0008,0092",
	"0008,0094",
	"0008,0096",
	"0008,009C",
	"0008,009D",
	"0008,0106",
	"0008,0107",
	"0008,0201",
	"0008,1000",
	"0008,1010",
	"0008,1030",
	"0008,103E",
	"0008,1040",
	"0008,1041",
	"0008,1048",
	"0008,1049",
	"0008,1050",
	"0008,1052",
	"0008,1060",
	"0008,1062",
	"0008,1070",
	"0008,1072",
	"0008,1080",
	"0008,1084",
	"0008,1088",
	"0008,1110",
	"0008,1111",
	"0008,1120",
	"0008,1140",
	"0008,1155",
	"0008,1195",
	"0008,1301",
	"0008,1302",
	"0008,1303",
	"0008,1304",
	"0008,2111",
	"0008,2112",
	"0008,3010",
	"0008,4000",
	"0010,0010",
	"0010,0011",
	"0010,0012",
	"0010,0013",
	"0010,0014",
	"0010,0015",
	"0010,0016",
	"0010,0020",
	"0010,0021",
	"0010,0030",
	"0010,0032",
	"0010,0040",
	"0010,0041",
	"0010,0042",
	"0010,0043",
	"0010,0044",
	"0010,0045",
	"0010,0046",
	"0010,0047",
	"0010,0050",
	"0010,0101",
	"0010,0102",
	"0010,1000",
	"0010,1001",
	"0010,1002",
	"0010,1005",
	"0010,1010",
	"0010,1020",
	"0010,1030",
	"0010,1040",
	"0010,1050",
	"0010,1060",
	"0010,1080",
	"0010,1081",
	"0010,1090",
	"0010,1100",
	"0010,2000",
	"0010,2110",
	"0010,2150",
	"0010,2152",
	"0010,2154",
	"0010,2155",
	"0010,2160",
	"0010,2161",
	"0010,2162",
	"0010,2180",
	"0010,21A0",
	"0010,21B0",
	"0010,21C0",
	"0010,21D0",
	"0010,21F0",
	"0010,2203",
	"0010,2297",
	"0010,2299",
	"0010,4000",
	"0012,0010",
	"0012,0020",
	"0012,0021",
	"0012,0022",
	"0012,0023",
	"0012,0030",
	"0012,0031",
	"0012,0032",
	"0012,0040",
	"0012,0041",
	"0012,0042",
	"0012,0043",
	"0012,0050",
	"0012,0051",
	"0012,0055",
	"0012,0060",
	"0012,0071",
	"0012,0072",
	"0012,0073",
	"0012,0081",
	"0012,0082",
	"0012,0086",
	"0012,0087",
	"0014,407C",
	"0014,407E",
	"0016,002B",
	"0016,004B",
	"0016,004D",
	"0016,004E",
	"0016,004F",
	"0016,0050",
	"0016,0051",
	"0016,0070",
	"0016,0071",
	"0016,0072",
	"0016,0073",
	"0016,0074",
	"0016,0075",
	"0016,0076",
	"0016,0077",
	"0016,0078",
	"0016,0079",
	"0016,007A",
	"0016,007B",
	"0016,007C",
	"0016,007D",
	"0016,007E",
	"0016,007F",
	"0016,0080",
	"0016,0081",
	"0016,0082",
	"0016,0083",
	"0016,0084",
	"0016,0085",
	"0016,0086",
	"0016,0087",
	"0016,0088",
	"0016,0089",
	"0016,008A",
	"0016,008B",
	"0016,008C",
	"0016,008D",
	"0016,008E",
	"0018,0010",
	"0018,0027",
	"0018,0035",
	"0018,1000",
	"0018,1002",
	"0018,1004",
	"0018,1005",
	"0018,1007",
	"0018,1008",
	"0018,1009",
	"0018,100A",
	"0018,100B",
	"0018,1010",
	"0018,1011",
	"0018,1012",
	"0018,1014",
	"0018,1030",
	"0018,1042",
	"0018,1043",
	"0018,1072",
	"0018,1073",
	"0018,1078",
	"0018,1079",
	"0018,11BB",
	"0018,1200",
	"0018,1201",
	"0018,1202",
	"0018,1203",
	"0018,1204",
	"0018,1205",
	"0018,1400",
	"0018,2042",
	"0018,4000",
	"0018,5011",
	"0018,700A",
	"0018,700C",
	"0018,700E",
	"0018,9074",
	"0018,9151",
	"0018,9185",
	"0018,9367",
	"0018,9369",
	"0018,936A",
	"0018,9371",
	"0018,9373",
	"0018,937B",
	"0018,937F",
	"0018,9424",
	"0018,9516",
	"0018,9517",
	"0018,9623",
	"0018,9701",
	"0018,9804",
	"0018,9919",
	"0018,9937",
	"0018,A002",
	"0018,A003",
	"0020,000D",
	"0020,000E",
	"0020,0010",
	"0020,0027",
	"0020,0052",
	"0020,0200",
	"0020,3401",
	"0020,3403",
	"0020,3405",
	"0020,3406",
	"0020,4000",
	"0020,9158",
	"0020,9161",
	"0020,9164",
	"0028,1199",
	"0028,1214",
	"0028,4000",
	"0032,0012",
	"0032,0032",
	"0032,0033",
	"0032,0034",
	"0032,0035",
	"0032,1000",
	"0032,1001",
	"0032,1010",
	"0032,1011",
	"0032,1020",
	"0032,1021",
	"0032,1030",
	"0032,1032",
	"0032,1033",
	"0032,1040",
	"0032,1041",
	"0032,1050",
	"0032,1051",
	"0032,1060",
	"0032,1066",
	"0032,1067",
	"0032,1070",
	"0032,4000",
	"0034,0001",
	"0034,0002",
	"0034,0005",
	"0034,0007",
	"0038,0004",
	"0038,0010",
	"0038,0011",
	"0038,0014",
	"0038,001A",
	"0038,001B",
	"0038,001C",
	"0038,001D",
	"0038,001E",
	"0038,0020",
	"0038,0021",
	"0038,0030",
	"0038,0032",
	"0038,0040",
	"0038,0050",
	"0038,0060",
	"0038,0061",
	"0038,0062",
	"0038,0064",
	"0038,0300",
	"0038,0400",
	"0038,0500",
	"0038,4000",
	"003A,0020",
	"003A,0203",
	"003A,020C",
	"003A,0310",
	"003A,0314",
	"003A,0329",
	"003A,032B",
	"0040,0001",
	"0040,0002",
	"0040,0003",
	"0040,0004",
	"0040,0005",
	"0040,0006",
	"0040,0007",
	"0040,0009",
	"0040,000B",
	"0040,0010",
	"0040,0011",
	"0040,0012",
	"0040,0241",
	"0040,0242",
	"0040,0243",
	"0040,0244",
	"0040,0245",
	"0040,0250",
	"0040,0251",
	"0040,0253",
	"0040,0254",
	"0040,0275",
	"0040,0280",
	"0040,0310",
	"0040,050A",
	"0040,0512",
	"0040,0513",
	"0040,051A",
	"0040,0551",
	"0040,0554",
	"0040,0555",
	"0040,0556",
	"0040,0562",
	"0040,0600",
	"0040,0602",
	"0040,0610",
	"0040,06FA",
	"0040,1001",
	"0040,1002",
	"0040,1004",
	"0040,1005",
	"0040,100A",
	"0040,1010",
	"0040,1011",
	"0040,1101",
	"0040,1102",
	"0040,1103",
	"0040,1104",
	"0040,1400",
	"0040,2001",
	"0040,2004",
	"0040,2005",
	"0040,2008",
	"0040,2009",
	"0040,2010",
	"0040,2011",
	"0040,2016",
	"0040,2017",
	"0040,2400",
	"0040,3001",
	"0040,4005",
	"0040,4008",
	"0040,4010",
	"0040,4011",
	"0040,4023",
	"0040,4025",
	"0040,4027",
	"0040,4028",
	"0040,4030",
	"0040,4034",
	"0040,4035",
	"0040,4036",
	"0040,4037",
	"0040,4050",
	"0040,4051",
	"0040,4052",
	"0040,A023",
	"0040,A024",
	"0040,A027",
	"0040,A030",
	"0040,A032",
	"0040,A033",
	"0040,A034",
	"0040,A035",
	"0040,A073",
	"0040,A075",
	"0040,A078",
	"0040,A07A",
	"0040,A07C",
	"0040,A082",
	"0040,A088",
	"0040,A110",
	"0040,A112",
	"0040,A120",
	"0040,A121",
	"0040,A122",
	"0040,A123",
	"0040,A124",
	"0040,A13A",
	"0040,A171",
	"0040,A172",
	"0040,A192",
	"0040,A193",
	"0040,A307",
	"0040,A352",
	"0040,A353",
	"0040,A354",
	"0040,A358",
	"0040,A402",
	"0040,A730",
	"0040,B020",
	"0040,B034",
	"0040,B036",
	"0040,B03B",
	"0040,B03F",
	"0040,DB06",
	"0040,DB07",
	"0040,DB0C",
	"0040,DB0D",
	"0040,E004",
	"0040,E012",
	"0042,0011",
	"0044,0004",
	"0044,000B",
	"0044,0010",
	"0044,0104",
	"0044,0105",
	"0050,001B",
	"0050,0020",
	"0050,0021",
	"0062,0021",
	"0064,0003",
	"0068,6226",
	"0068,6270",
	"006A,0003",
	"006A,0005",
	"006A,0006",
	"0070,0001",
	"0070,0006",
	"0070,0082",
	"0070,0083",
	"0070,0084",
	"0070,0086",
	"0070,031A",
	"0070,1101",
	"0070,1102",
	"0072,000A",
	"0072,005E",
	"0072,005F",
	"0072,0061",
	"0072,0063",
	"0072,0065",
	"0072,0066",
	"0072,0068",
	"0072,006A",
	"0072,006B",
	"0072,006C",
	"0072,006D",
	"0072,006E",
	"0072,0070",
	"0072,0071",
	"0074,1234",
	"0074,1236",
	"0088,0140",
	"0088,0200",
	"0088,0904",
	"0088,0906",
	"0088,0910",
	"0088,0912",
	"0100,0420",
	"0400,0100",
	"0400,0105",
	"0400,0115",
	"0400,0310",
	"0400,0402",
	"0400,0403",
	"0400,0404",
	"0400,0550",
	"0400,0551",
	"0400,0552",
	"0400,0561",
	"0400,0562",
	"0400,0563",
	"0400,0564",
	"0400,0565",
	"0400,0600",
	"2030,0020",
	"2100,0040",
	"2100,0050",
	"2100,0070",
	"2100,0140",
	"2200,0002",
	"2200,0005",
	"3002,0121",
	"3002,0123",
	"3004,007F",
	"3006,0002",
	"3006,0004",
	"3006,0006",
	"3006,0008",
	"3006,0009",
	"3006,0024",
	"3006,0026",
	"3006,0028",
	"3006,002D",
	"3006,002E",
	"3006,0038",
	"3006,004D",
	"3006,004E",
	"3006,0085",
	"3006,0088",
	"3006,00A6",
	"3006,00C2",
	"3008,0024",
	"3008,0025",
	"3008,0054",
	"3008,0056",
	"3008,0105",
	"3008,0162",
	"3008,0164",
	"3008,0166",
	"3008,0168",
	"3008,0250",
	"3008,0251",
	"300A,0002",
	"300A,0003",
	"300A,0004",
	"300A,0006",
	"300A,0007",
	"300A,000B",
	"300A,000E",
	"300A,0013",
	"300A,0016",
	"300A,0054",
	"300A,0072",
	"300A,0083",
	"300A,00B2",
	"300A,00C3",
	"300A,00DD",
	"300A,0196",
	"300A,01A6",
	"300A,01B2",
	"300A,0216",
	"300A,022C",
	"300A,022E",
	"300A,02EB",
	"300A,0608",
	"300A,0609",
	"300A,0611",
	"300A,0615",
	"300A,0619",
	"300A,0623",
	"300A,062A",
	"300A,0650",
	"300A,0676",
	"300A,067C",
	"300A,067D",
	"300A,0700",
	"300A,0734",
	"300A,0736",
	"300A,073A",
	"300A,0741",
	"300A,0742",
	"300A,0760",
	"300A,0783",
	"300A,0785",
	"300A,078E",
	"300A,0792",
	"300A,0794",
	"300A,079A",
	"300C,0113",
	"300C,0127",
	"300E,0004",
	"300E,0005",
	"300E,0008",
	"3010,0006",
	"3010,000B",
	"3010,000F",
	"3010,0013",
	"3010,0015",
	"3010,0017",
	"3010,001B",
	"3010,002D",
	"3010,0031",
	"3010,0033",
	"3010,0034",
	"3010,0035",
	"3010,0036",
	"3010,0037",
	"3010,0038",
	"3010,003B",
	"3010,0043",
	"3010,004C",
	"3010,004D",
	"3010,0054",
	"3010,0056",
	"3010,005A",
	"3010,005C",
	"3010,0061",
	"3010,006E",
	"3010,006F",
	"3010,0077",
	"3010,007A",
	"3010,007B",
	"3010,007F",
	"3010,0081",
	"3010,0085",
	"4000,0010",
	"4000,4000",
	"4008,0040",
	"4008,0042",
	"4008,0100",
	"4008,0101",
	"4008,0102",
	"4008,0108",
	"4008,0109",
	"4008,010A",
	"4008,010B",
	"4008,010C",
	"4008,0111",
	"4008,0112",
	"4008,0113",
	"4008,0114",
	"4008,0115",
	"4008,0118",
	"4008,0119",
	"4008,011A",
	"4008,0200",
	"4008,0202",
	"4008,0300",
	"4008,4000",
	"FFFA,FFFA",
	"FFFC,FFFC",
}

func TestDefaultProfileMatchesEveryNonKBasicRow(t *testing.T) {
	profile := NewSecurityProfile(BasicProfile)
	for _, value := range nonKBasicTags {
		parsed, err := tag.Parse("(" + value + ")")
		if err != nil {
			t.Fatalf("parse tag %q: %v", value, err)
		}
		if _, found := profile.FindAction(parsed); !found {
			t.Errorf("Table E.1-1 row (%s) is not covered by the default profile", value)
		}
	}
	// Rows represented by group patterns rather than exact tags.
	for _, mask := range []string{"(0019,0010)", "(5000,0005)", "(6000,4000)", "(6000,3000)"} {
		parsed, err := tag.Parse(mask)
		if err != nil {
			t.Fatalf("parse tag %q: %v", mask, err)
		}
		if _, found := profile.FindAction(parsed); !found {
			t.Errorf("pattern row %s is not covered by the default profile", mask)
		}
	}
}

func TestDefaultProfileRemovesTableE1_1Identifiers(t *testing.T) {
	gpsElements := []uint16{
		0x002B, 0x004B, 0x004E, 0x004F, 0x0050, 0x0051,
		0x0070, 0x0071, 0x0072, 0x0073, 0x0074, 0x0075, 0x0076, 0x0077, 0x0078,
		0x0079, 0x007A, 0x007B, 0x007C, 0x007D, 0x007E, 0x007F, 0x0080, 0x0081,
		0x0082, 0x0083, 0x0084, 0x0085, 0x0086, 0x0087, 0x0088, 0x0089, 0x008A,
		0x008B, 0x008C, 0x008D, 0x008E,
	}

	ds := dataset.New()
	ds.SetAutoValidate(false)
	addString := func(elemTag *tag.Tag, v *vr.VR, value string) {
		t.Helper()
		if err := ds.Add(element.NewString(elemTag, v, []string{value})); err != nil {
			t.Fatalf("add %s: %v", elemTag, err)
		}
	}
	addString(tag.NameToUse, vr.PN, "Alias^Name")
	addString(tag.PatientTelecomInformation, vr.LO, "+1-555-0100")
	addString(tag.RetrieveAETitle, vr.AE, "RETRIEVE")
	addString(tag.StationAETitle, vr.AE, "STATION")
	addString(tag.InstanceCreationDate, vr.DA, "20260102")
	addString(tag.InstanceCreationTime, vr.TM, "120000")

	udiItem := dataset.New()
	udiItem.SetAutoValidate(false)
	_ = udiItem.Add(element.NewString(tag.DeviceSerialNumber, vr.LO, []string{"SN-12345"}))
	if err := ds.Add(dataset.NewSequenceWithItems(tag.UDISequence, []*dataset.Dataset{udiItem})); err != nil {
		t.Fatalf("add UDI sequence: %v", err)
	}
	for _, groupElement := range gpsElements {
		addString(tag.New(0x0016, groupElement), vr.UT, "GPS data")
	}
	addString(tag.New(0x0400, 0x0100), vr.UI, "1.2.826.0.1.3680043.10.999")
	addString(tag.FailedSOPInstanceUIDList, vr.UI, "1.2.3.1\\1.2.3.2")
	if err := ds.Add(dataset.NewSequenceWithItems(tag.GenderIdentityCodeSequence, []*dataset.Dataset{dataset.New()})); err != nil {
		t.Fatalf("add gender identity sequence: %v", err)
	}
	addString(tag.GenderIdentityComment, vr.UT, "comment")
	addString(tag.EthnicGroups, vr.UC, "Ethnicity")
	addString(tag.PatientName, vr.PN, "Doe^Jane")

	if err := NewAnonymizer(NewSecurityProfile(BasicProfile)).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}

	for _, removed := range []*tag.Tag{
		tag.NameToUse,
		tag.PatientTelecomInformation,
		tag.RetrieveAETitle,
		tag.StationAETitle,
		tag.UDISequence,
		tag.GenderIdentityCodeSequence,
		tag.GenderIdentityComment,
		tag.EthnicGroups,
	} {
		if ds.Contains(removed) {
			t.Errorf("%s was not removed by the Basic profile", removed)
		}
	}
	// X/D and X/Z/D rows are dummied rather than removed.
	for _, dummied := range []struct {
		tag  *tag.Tag
		want string
	}{
		{tag: tag.InstanceCreationDate, want: "19000101"},
		{tag: tag.InstanceCreationTime, want: "000000"},
	} {
		got, ok := ds.GetString(dummied.tag)
		if !ok {
			t.Errorf("%s is missing, want a dummy value", dummied.tag)
			continue
		}
		if got != dummied.want {
			t.Errorf("%s = %q, want %q", dummied.tag, got, dummied.want)
		}
	}
	for _, groupElement := range gpsElements {
		if ds.Contains(tag.New(0x0016, groupElement)) {
			t.Errorf("(0016,%04X) GPS attribute was not removed", groupElement)
		}
	}

	signature := tag.New(0x0400, 0x0100)
	value, ok := ds.GetString(signature)
	if !ok || value == "1.2.826.0.1.3680043.10.999" {
		t.Errorf("(0400,0100) = %q, want a replaced UID", value)
	}
	failed, ok := ds.GetStrings(tag.FailedSOPInstanceUIDList)
	if !ok || len(failed) != 2 {
		t.Fatalf("Failed SOP Instance UID List = %v, want 2 remapped values", failed)
	}
	if failed[0] == "1.2.3.1" || failed[1] == "1.2.3.2" {
		t.Errorf("Failed SOP Instance UID List was not remapped: %v", failed)
	}
	if failed[0] == failed[1] {
		t.Errorf("Failed SOP Instance UID List values collapsed: %v", failed)
	}
	if got, _ := ds.GetString(tag.PatientIdentityRemoved); got != testPatientRemovedYes {
		t.Errorf("PatientIdentityRemoved = %q, want YES", got)
	}
}

func TestAnonymizerKeepsVerificationFlag(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewString(tag.VerificationFlag, vr.CS, []string{"UNVERIFIED"})); err != nil {
		t.Fatal(err)
	}
	if err := NewAnonymizer(NewSecurityProfile(BasicProfile)).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}
	if got, ok := ds.GetString(tag.VerificationFlag); !ok || got != "UNVERIFIED" {
		t.Fatalf("Verification Flag = %q, %v, want UNVERIFIED and present", got, ok)
	}
}

func TestPatientIdentityRemovedReflectsRetainedAttributes(t *testing.T) {
	tests := []struct {
		name         string
		profile      *SecurityProfile
		wantRemoved  string
		wantRetained []*tag.Tag
	}{
		{
			name:         "empty profile retains patient name",
			profile:      NewSecurityProfile(0),
			wantRemoved:  "NO",
			wantRetained: []*tag.Tag{tag.PatientName},
		},
		{
			name:         "minimal profile retains institution",
			profile:      NewMinimalProfile(),
			wantRemoved:  "NO",
			wantRetained: []*tag.Tag{tag.InstitutionName},
		},
		{
			name:         "retain patient characteristics reports no",
			profile:      NewSecurityProfile(BasicProfile | RetainPatientChars),
			wantRemoved:  "NO",
			wantRetained: []*tag.Tag{tag.PatientSex},
		},
		{
			name:        "basic profile removes identifiers",
			profile:     NewSecurityProfile(BasicProfile),
			wantRemoved: testPatientRemovedYes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dataset.New()
			for _, item := range []element.Element{
				element.NewPersonName(tag.PatientName, []string{"Doe^Jane"}),
				element.NewString(tag.PatientSex, vr.CS, []string{"F"}),
				element.NewString(tag.InstitutionName, vr.LO, []string{"General Hospital"}),
			} {
				if err := ds.Add(item); err != nil {
					t.Fatalf("add %s: %v", item.Tag(), err)
				}
			}
			if err := NewAnonymizer(tt.profile).AnonymizeInPlace(ds); err != nil {
				t.Fatalf("AnonymizeInPlace() error = %v", err)
			}
			if got, _ := ds.GetString(tag.PatientIdentityRemoved); got != tt.wantRemoved {
				t.Errorf("PatientIdentityRemoved = %q, want %q", got, tt.wantRemoved)
			}
			method, _ := ds.GetString(tag.DeidentificationMethod)
			for _, retained := range tt.wantRetained {
				if !strings.Contains(method, retained.String()) {
					t.Errorf("DeidentificationMethod = %q, want retained %s", method, retained)
				}
			}
		})
	}
}

func TestProfileAcceptsTableE1_1UIDAsteriskNotation(t *testing.T) {
	profile, err := NewProfileFromReader(strings.NewReader("0008,2112;X/Z/U*;;K;;;;;;;;\n"), BasicProfile)
	if err != nil {
		t.Fatalf("NewProfileFromReader() error = %v", err)
	}
	parsed, err := tag.Parse("(0008,2112)")
	if err != nil {
		t.Fatal(err)
	}
	action, found := profile.FindAction(parsed)
	if !found || action != ActionZ {
		t.Fatalf("action = %v, found = %v, want Z, true", action, found)
	}
}

func TestTableE1_1ActionsMatchSRContentAttributes(t *testing.T) {
	tests := []struct {
		tag    string
		action SecurityProfileAction
	}{
		{tag: "0040,A027", action: ActionD}, // Verifying Organization
		{tag: "0040,A073", action: ActionD}, // Verifying Observer Sequence
		{tag: "0040,A075", action: ActionD}, // Verifying Observer Name
		{tag: "0040,A088", action: ActionZ}, // Verifying Observer Identification Code Sequence
		{tag: "0040,A120", action: ActionD}, // DateTime
		{tag: "0040,A121", action: ActionD}, // Date
		{tag: "0040,A122", action: ActionD}, // Time
		{tag: "0040,A123", action: ActionD}, // Person Name
		{tag: "0040,A124", action: ActionU}, // UID
		{tag: "0040,A730", action: ActionD}, // Content Sequence
		{tag: "0040,A493", action: 0},       // Verification Flag: no rule (Type 1)
	}
	profile := NewSecurityProfile(BasicProfile)
	for _, tt := range tests {
		parsed, err := tag.Parse("(" + tt.tag + ")")
		if err != nil {
			t.Fatalf("parse tag %s: %v", tt.tag, err)
		}
		action, found := profile.FindAction(parsed)
		if tt.action == 0 {
			if found {
				t.Errorf("(%s) action = %v, want no rule", tt.tag, action)
			}
			continue
		}
		if !found || action != tt.action {
			t.Errorf("(%s) action = %v, found = %v, want %v", tt.tag, action, found, tt.action)
		}
	}
}

func TestCleanStructuredContentOptionMapsToClean(t *testing.T) {
	parsed, err := tag.Parse("(0040,A730)")
	if err != nil {
		t.Fatal(err)
	}
	basic := NewSecurityProfile(BasicProfile)
	if action, found := basic.FindAction(parsed); !found || action != ActionD {
		t.Fatalf("Basic Content Sequence action = %v, found = %v, want D", action, found)
	}
	clean := NewSecurityProfile(BasicProfile | CleanStructdCont)
	if action, found := clean.FindAction(parsed); !found || action != ActionC {
		t.Fatalf("Clean Structd Cont Content Sequence action = %v, found = %v, want C", action, found)
	}
}

func TestPatientCharacteristicsOptionSemantics(t *testing.T) {
	sex, err := tag.Parse("(0010,0040)")
	if err != nil {
		t.Fatal(err)
	}
	age, err := tag.Parse("(0010,1010)")
	if err != nil {
		t.Fatal(err)
	}
	basic := NewSecurityProfile(BasicProfile)
	if action, _ := basic.FindAction(sex); action != ActionZ {
		t.Errorf("Basic Patient's Sex action = %v, want Z", action)
	}
	if action, _ := basic.FindAction(age); action != ActionX {
		t.Errorf("Basic Patient's Age action = %v, want X", action)
	}

	retain := NewSecurityProfile(BasicProfile | RetainPatientChars)
	if action, _ := retain.FindAction(sex); action != ActionK {
		t.Errorf("Retain Patient Characteristics Sex action = %v, want K", action)
	}
	if action, _ := retain.FindAction(age); action != ActionK {
		t.Errorf("Retain Patient Characteristics Age action = %v, want K", action)
	}

	// Clean Descriptions must not turn the enumerated Patient's Sex into a dummy.
	clean := NewSecurityProfile(BasicProfile | CleanDesc)
	if action, _ := clean.FindAction(sex); action != ActionZ {
		t.Errorf("Clean Descriptions Patient's Sex action = %v, want Z", action)
	}
}

func TestAnonymizerShiftsDateTimeWithTimezone(t *testing.T) {
	ds := dataset.New()
	if err := ds.Add(element.NewString(tag.AcquisitionDateTime, vr.DT, []string{"20260102120000.123+0800"})); err != nil {
		t.Fatal(err)
	}
	profile := NewSecurityProfile(BasicProfile | RetainLongModifDates)
	if err := NewAnonymizer(profile).AnonymizeInPlace(ds); err != nil {
		t.Fatalf("AnonymizeInPlace() error = %v", err)
	}
	want := time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -10000).Format("20060102") + "120000.123+0800"
	if got, _ := ds.GetString(tag.AcquisitionDateTime); got != want {
		t.Fatalf("Acquisition DateTime = %q, want %q", got, want)
	}
}

func TestShiftDateTimePreservesPrecisionAndTimezone(t *testing.T) {
	shiftedDate := time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -10000).Format("20060102")
	tests := []struct {
		input string
		want  string
	}{
		{input: "20260102120000+0800", want: shiftedDate + "120000+0800"},
		{input: "20260102120000.123", want: shiftedDate + "120000.123"},
		{input: "20260102120000.123456-0500", want: shiftedDate + "120000.123456-0500"},
		{input: "20260102", want: shiftedDate},
		{input: "2026010212", want: shiftedDate + "12"},
	}
	for _, tt := range tests {
		got, ok := shiftDTValue(tt.input, -10000)
		if !ok {
			t.Fatalf("shiftDTValue(%q) failed", tt.input)
		}
		if got != tt.want {
			t.Errorf("shiftDTValue(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

