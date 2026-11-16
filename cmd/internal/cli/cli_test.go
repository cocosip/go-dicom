// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package cli

import "testing"

func TestCommandCheckArgs(t *testing.T) {
	command := Command{MinArgs: 1, MaxArgs: 1}

	if err := command.checkArgs(nil); err == nil {
		t.Error("checkArgs(nil) = nil, want missing-argument error")
	}
	if err := command.checkArgs([]string{"file.dcm"}); err != nil {
		t.Errorf("checkArgs(one argument) = %v, want nil", err)
	}
	if err := command.checkArgs([]string{"file.dcm", "extra"}); err == nil {
		t.Error("checkArgs(two arguments) = nil, want unexpected-argument error")
	}
}
