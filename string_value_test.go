/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

package supertypes

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestStringValueToTerraformValue(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       StringValue
		expectation interface{}
	}
	tests := map[string]testCase{
		"known": {
			input:       NewStringValue("test"),
			expectation: tftypes.NewValue(tftypes.String, "test"),
		},
		"unknown": {
			input:       NewStringUnknown(),
			expectation: tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		},
		"null": {
			input:       NewStringNull(),
			expectation: tftypes.NewValue(tftypes.String, nil),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			got, err := test.input.ToTerraformValue(ctx)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
				return
			}
			if !cmp.Equal(got, test.expectation) {
				t.Errorf("Expected %+v, got %+v", test.expectation, got)
			}
		})
	}
}

func TestStringValueEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       StringValue
		candidate   attr.Value
		expectation bool
	}
	tests := map[string]testCase{
		"known-known-same": {
			input:       NewStringValue("test"),
			candidate:   NewStringValue("test"),
			expectation: true,
		},
		"known-known-diff": {
			input:       NewStringValue("test"),
			candidate:   NewStringValue("not-test"),
			expectation: false,
		},
		"known-unknown": {
			input:       NewStringValue("test"),
			candidate:   NewStringUnknown(),
			expectation: false,
		},
		"known-null": {
			input:       NewStringValue("test"),
			candidate:   NewStringNull(),
			expectation: false,
		},
		"unknown-value": {
			input:       NewStringUnknown(),
			candidate:   NewStringValue("test"),
			expectation: false,
		},
		"unknown-unknown": {
			input:       NewStringUnknown(),
			candidate:   NewStringUnknown(),
			expectation: true,
		},
		"unknown-null": {
			input:       NewStringUnknown(),
			candidate:   NewStringNull(),
			expectation: false,
		},
		"null-known": {
			input:       NewStringNull(),
			candidate:   NewStringValue("test"),
			expectation: false,
		},
		"null-unknown": {
			input:       NewStringNull(),
			candidate:   NewStringUnknown(),
			expectation: false,
		},
		"null-null": {
			input:       NewStringNull(),
			candidate:   NewStringNull(),
			expectation: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := test.input.Equal(test.candidate)
			if !cmp.Equal(got, test.expectation) {
				t.Errorf("Expected %v, got %v", test.expectation, got)
			}
		})
	}
}

func TestStringValueIsNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    StringValue
		expected bool
	}{
		"known": {
			input:    NewStringValue("test"),
			expected: false,
		},
		"null": {
			input:    NewStringNull(),
			expected: true,
		},
		"unknown": {
			input:    NewStringUnknown(),
			expected: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.IsNull()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestStringValueIsUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    StringValue
		expected bool
	}{
		"known": {
			input:    NewStringValue("test"),
			expected: false,
		},
		"null": {
			input:    NewStringNull(),
			expected: false,
		},
		"unknown": {
			input:    NewStringUnknown(),
			expected: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.IsUnknown()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestStringValueString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       StringValue
		expectation string
	}
	tests := map[string]testCase{
		"known-non-empty": {
			input:       NewStringValue("test"),
			expectation: `"test"`,
		},
		"known-empty": {
			input:       NewStringValue(""),
			expectation: `""`,
		},
		"known-quotes": {
			input:       NewStringValue(`testing is "fun"`),
			expectation: `"testing is \"fun\""`,
		},
		"unknown": {
			input:       NewStringUnknown(),
			expectation: "<unknown>",
		},
		"null": {
			input:       NewStringNull(),
			expectation: "<null>",
		},
		"zero-value": {
			input:       StringValue{},
			expectation: `<null>`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := test.input.String()
			if !cmp.Equal(got, test.expectation) {
				t.Errorf("Expected %q, got %q", test.expectation, got)
			}
		})
	}
}

func TestStringValueValueString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    StringValue
		expected string
	}{
		"known": {
			input:    NewStringValue("test"),
			expected: "test",
		},
		"null": {
			input:    NewStringNull(),
			expected: "",
		},
		"unknown": {
			input:    NewStringUnknown(),
			expected: "",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.ValueString()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestStringValueValueStringPointer(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    StringValue
		expected *string
	}{
		"known": {
			input:    NewStringValue("test"),
			expected: pointer("test"),
		},
		"null": {
			input:    NewStringNull(),
			expected: nil,
		},
		"unknown": {
			input:    NewStringUnknown(),
			expected: pointer(""),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.ValueStringPointer()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestNewStringPointerValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		value    *string
		expected StringValue
	}{
		"nil": {
			value:    nil,
			expected: NewStringNull(),
		},
		"value": {
			value:    pointer("test"),
			expected: NewStringValue("test"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewStringPointerValue(testCase.value)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestStringValue(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		v := NewStringValue("hello")
		assert.Equal(t, "hello", v.Get())
	})

	t.Run("GetPtr", func(t *testing.T) {
		v := NewStringValue("hello")
		ptr := v.GetPtr()
		assert.NotNil(t, ptr)
		assert.Equal(t, "hello", *ptr)

		v.SetNull()
		ptr = v.GetPtr()
		assert.Nil(t, ptr)

		v.SetUnknown()
		ptr = v.GetPtr()
		assert.NotNil(t, ptr)
		assert.Equal(t, "", *ptr)
	})

	t.Run("Set", func(t *testing.T) {
		v := NewStringValue("hello")
		assert.Equal(t, "hello", v.Get())

		v.Set("world")
		assert.Equal(t, "world", v.Get())

		v.Set("")
		assert.True(t, v.IsNull())
	})

	t.Run("SetPtr", func(t *testing.T) {
		v := NewStringValue("hello")
		assert.Equal(t, "hello", v.Get())

		ptr := "world"
		v.SetPtr(&ptr)
		assert.Equal(t, "world", v.Get())

		v.SetPtr(nil)
		assert.True(t, v.IsNull())
	})

	t.Run("SetNull", func(t *testing.T) {
		v := NewStringValue("hello")
		assert.False(t, v.IsNull())

		v.SetNull()
		assert.True(t, v.IsNull())
	})

	t.Run("SetUnknown", func(t *testing.T) {
		v := NewStringValue("hello")
		assert.False(t, v.IsUnknown())

		v.SetUnknown()
		assert.True(t, v.IsUnknown())
	})
}

func TestStringValueIsKnown(t *testing.T) {
	// Test NewStringValue
	s := "hello"
	v := NewStringValue(s)
	if v.Get() != s {
		t.Errorf("Expected %q, got %q", s, v.Get())
	}

	// Test NewStringPointerValue
	sp := &s
	v = NewStringPointerValue(sp)
	if v.Get() != s {
		t.Errorf("Expected %q, got %q", s, v.Get())
	}

	// Test Set
	s2 := "world"
	v.Set(s2)
	if v.Get() != s2 {
		t.Errorf("Expected %q, got %q", s2, v.Get())
	}

	// Test SetPtr
	sp2 := &s2
	v.SetPtr(sp2)
	if v.Get() != s2 {
		t.Errorf("Expected %q, got %q", s2, v.Get())
	}

	// Test SetNull
	v.SetNull()
	if v.Get() != "" {
		t.Errorf("Expected empty string, got %q", v.Get())
	}

	// Test SetUnknown
	v.SetUnknown()
	if v.Get() != "" {
		t.Errorf("Expected empty string, got %q", v.Get())
	}

	// Test IsKnown
	v.Set(s)
	if !v.IsKnown() {
		t.Errorf("Expected IsKnown to be true, got false")
	}
	v.SetNull()
	if v.IsKnown() {
		t.Errorf("Expected IsKnown to be false, got true")
	}
	v.SetUnknown()
	if v.IsKnown() {
		t.Errorf("Expected IsKnown to be false, got true")
	}
}
