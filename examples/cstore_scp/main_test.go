// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package main

import "testing"

func TestSanitizeAEPathComponent(t *testing.T) {
	for _, test := range []struct {
		name string
		ae   string
		want string
	}{
		{name: "plain title", ae: "STORESCU", want: "STORESCU"},
		{name: "trimmed title", ae: "  STORESCP  ", want: "STORESCP"},
		{name: "path separators", ae: "../evil", want: "___evil"},
		{name: "windows drive", ae: `C:\temp`, want: "C__temp"},
		{name: "symbols and spaces", ae: "A B/C", want: "A_B_C"},
		{name: "empty title", ae: "", want: "unknown_ae"},
		{name: "overlong title", ae: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", want: "ABCDEFGHIJKLMNOP"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := sanitizeAEPathComponent(test.ae); got != test.want {
				t.Fatalf("sanitizeAEPathComponent(%q) = %q, want %q", test.ae, got, test.want)
			}
			if got := sanitizeAEPathComponent(test.ae); got == "." || got == ".." {
				t.Fatalf("sanitizeAEPathComponent(%q) = %q, which is a traversal component", test.ae, got)
			}
		})
	}
}
