// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package association

import (
	"testing"
)

func TestServiceApplicationInfo(t *testing.T) {
	t.Run("New from bytes", func(t *testing.T) {
		data := []byte{1, 0, 1}
		info := NewServiceApplicationInfo(data)

		if info.Count() != 3 {
			t.Errorf("Count() = %d, want 3", info.Count())
		}

		if v, ok := info.Get(1); !ok || v != 1 {
			t.Errorf("Get(1) = %d, %v, want 1, true", v, ok)
		}
		if v, ok := info.Get(2); !ok || v != 0 {
			t.Errorf("Get(2) = %d, %v, want 0, true", v, ok)
		}
	})

	t.Run("Set and Get", func(t *testing.T) {
		info := NewServiceApplicationInfoEmpty()

		if err := info.Set(1, 1); err != nil {
			t.Errorf("Set() error = %v", err)
		}
		if err := info.Set(2, 0); err != nil {
			t.Errorf("Set() error = %v", err)
		}

		if !info.Contains(1) {
			t.Error("Contains(1) = false, want true")
		}
		if info.Contains(3) {
			t.Error("Contains(3) = true, want false")
		}

		if v := info.GetBool(1, false); !v {
			t.Error("GetBool(1) = false, want true")
		}
	})

	t.Run("SetBool", func(t *testing.T) {
		info := NewServiceApplicationInfoEmpty()

		if err := info.SetBool(1, true); err != nil {
			t.Errorf("SetBool() error = %v", err)
		}
		if v, _ := info.Get(1); v != 1 {
			t.Errorf("Get(1) = %d, want 1", v)
		}

		if err := info.SetBool(2, false); err != nil {
			t.Errorf("SetBool() error = %v", err)
		}
		if v, _ := info.Get(2); v != 0 {
			t.Errorf("Get(2) = %d, want 0", v)
		}
	})

	t.Run("Values", func(t *testing.T) {
		info := NewServiceApplicationInfo([]byte{1, 0, 1, 1})
		values := info.Values()

		if len(values) != 4 {
			t.Errorf("Values() len = %d, want 4", len(values))
		}
	})

	t.Run("Remove", func(t *testing.T) {
		info := NewServiceApplicationInfo([]byte{1, 0, 1})
		info.Remove(2)

		if info.Contains(2) {
			t.Error("Contains(2) = true after Remove, want false")
		}
	})

	t.Run("Invalid index", func(t *testing.T) {
		info := NewServiceApplicationInfoEmpty()
		if err := info.Set(0, 1); err == nil {
			t.Error("Set(0) should return error")
		}
	})

	t.Run("String representation", func(t *testing.T) {
		info := NewServiceApplicationInfo([]byte{1, 0, 1})
		s := info.String()
		if s == "" {
			t.Error("String() returned empty string")
		}
	})

	t.Run("Empty values", func(t *testing.T) {
		info := NewServiceApplicationInfoEmpty()
		if info.Values() != nil {
			t.Error("Values() should return nil for empty info")
		}
	})
}
