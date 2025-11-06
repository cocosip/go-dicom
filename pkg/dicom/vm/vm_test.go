// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package vm_test

import (
	"testing"

	"github.com/cocosip/go-dicom/pkg/dicom/vm"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		wantMin         int
		wantMax         int
		wantMult        int
		wantUnlimited   bool
		wantStr         string
		wantErr         bool
	}{
		{"single value", "1", 1, 1, 1, false, "1", false},
		{"fixed range", "1-2", 1, 2, 1, false, "1-2", false},
		{"unlimited", "1-n", 1, int(^uint(0) >> 1), 1, true, "1-n", false},
		{"multiplicity 2", "2-2n", 2, int(^uint(0) >> 1), 2, true, "2-2n", false},
		{"multiplicity 3", "3-3n", 3, int(^uint(0) >> 1), 3, true, "3-3n", false},
		{"fixed large", "1-99", 1, 99, 1, false, "1-99", false},
		{"single digit", "4", 4, 4, 4, false, "4", false},
		{"invalid empty", "", 0, 0, 0, false, "", true},
		{"invalid format", "abc", 0, 0, 0, false, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vm.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if got.Minimum() != tt.wantMin {
				t.Errorf("Parse(%q).Minimum() = %v, want %v", tt.input, got.Minimum(), tt.wantMin)
			}
			if got.Maximum() != tt.wantMax {
				t.Errorf("Parse(%q).Maximum() = %v, want %v", tt.input, got.Maximum(), tt.wantMax)
			}
			if got.Multiplicity() != tt.wantMult {
				t.Errorf("Parse(%q).Multiplicity() = %v, want %v", tt.input, got.Multiplicity(), tt.wantMult)
			}
			if got.IsUnlimited() != tt.wantUnlimited {
				t.Errorf("Parse(%q).IsUnlimited() = %v, want %v", tt.input, got.IsUnlimited(), tt.wantUnlimited)
			}
			if got.String() != tt.wantStr {
				t.Errorf("Parse(%q).String() = %v, want %v", tt.input, got.String(), tt.wantStr)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		vm    *vm.VM
		count int
		want  bool
	}{
		{"VM1 with 1", vm.VM1, 1, true},
		{"VM1 with 2", vm.VM1, 2, false},
		{"VM1 with 0", vm.VM1, 0, false},
		{"VM1_2 with 1", vm.VM1_2, 1, true},
		{"VM1_2 with 2", vm.VM1_2, 2, true},
		{"VM1_2 with 3", vm.VM1_2, 3, false},
		{"VM1_n with 1", vm.VM1_n, 1, true},
		{"VM1_n with 100", vm.VM1_n, 100, true},
		{"VM1_n with 0", vm.VM1_n, 0, false},
		{"VM2_2n with 2", vm.VM2_2n, 2, true},
		{"VM2_2n with 4", vm.VM2_2n, 4, true},
		{"VM2_2n with 6", vm.VM2_2n, 6, true},
		{"VM2_2n with 3", vm.VM2_2n, 3, false},
		{"VM2_2n with 1", vm.VM2_2n, 1, false},
		{"VM3_3n with 3", vm.VM3_3n, 3, true},
		{"VM3_3n with 6", vm.VM3_3n, 6, true},
		{"VM3_3n with 9", vm.VM3_3n, 9, true},
		{"VM3_3n with 4", vm.VM3_3n, 4, false},
		{"VM3_3n with 2", vm.VM3_3n, 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vm.IsValid(tt.count); got != tt.want {
				t.Errorf("%s.IsValid(%d) = %v, want %v", tt.vm.String(), tt.count, got, tt.want)
			}
		})
	}
}

func TestStandardVMConstants(t *testing.T) {
	tests := []struct {
		name string
		vm   *vm.VM
		want string
	}{
		{"VM1", vm.VM1, "1"},
		{"VM1_2", vm.VM1_2, "1-2"},
		{"VM1_3", vm.VM1_3, "1-3"},
		{"VM1_8", vm.VM1_8, "1-8"},
		{"VM1_32", vm.VM1_32, "1-32"},
		{"VM1_99", vm.VM1_99, "1-99"},
		{"VM1_n", vm.VM1_n, "1-n"},
		{"VM2", vm.VM2, "2"},
		{"VM2_n", vm.VM2_n, "2-n"},
		{"VM2_2n", vm.VM2_2n, "2-2n"},
		{"VM3", vm.VM3, "3"},
		{"VM3_n", vm.VM3_n, "3-n"},
		{"VM3_3n", vm.VM3_3n, "3-3n"},
		{"VM4", vm.VM4, "4"},
		{"VM6", vm.VM6, "6"},
		{"VM16", vm.VM16, "16"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vm.String(); got != tt.want {
				t.Errorf("%s.String() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestMustParse(t *testing.T) {
	// Should not panic for valid input
	result := vm.MustParse("1-2")
	if result.String() != "1-2" {
		t.Errorf("MustParse(\"1-2\") = %v, want 1-2", result.String())
	}

	// Should panic for invalid input
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustParse with invalid input should panic")
		}
	}()
	vm.MustParse("invalid")
}

func TestParseCache(t *testing.T) {
	// Parse the same VM twice
	vm1, err := vm.Parse("1-2")
	if err != nil {
		t.Fatalf("Parse(\"1-2\") error = %v", err)
	}

	vm2, err := vm.Parse("1-2")
	if err != nil {
		t.Fatalf("Parse(\"1-2\") error = %v", err)
	}

	// Should return the same cached instance
	if vm1 != vm2 {
		t.Error("Parse should return cached instance for same input")
	}
}
