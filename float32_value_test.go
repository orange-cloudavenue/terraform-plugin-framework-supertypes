/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package supertypes

import (
	"context"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-framework/attr"
)

func TestNewFloat32ValueEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Float32Value
		candidate   attr.Value
		expectation bool
	}
	tests := map[string]testCase{
		"known-known-same": {
			input:       NewFloat32Value(123),
			candidate:   NewFloat32Value(123),
			expectation: true,
		},
		"known-known-diff": {
			input:       NewFloat32Value(123),
			candidate:   NewFloat32Value(456),
			expectation: false,
		},
		"known-unknown": {
			input:       NewFloat32Value(123),
			candidate:   NewFloat64Unknown(),
			expectation: false,
		},
		"known-null": {
			input:       NewFloat32Value(123),
			candidate:   NewFloat32Null(),
			expectation: false,
		},
		"unknown-value": {
			input:       NewFloat32Unknown(),
			candidate:   NewFloat32Value(123),
			expectation: false,
		},
		"unknown-unknown": {
			input:       NewFloat32Unknown(),
			candidate:   NewFloat32Unknown(),
			expectation: true,
		},
		"unknown-null": {
			input:       NewFloat32Unknown(),
			candidate:   NewFloat32Null(),
			expectation: false,
		},
		"null-known": {
			input:       NewFloat32Null(),
			candidate:   NewFloat32Value(123),
			expectation: false,
		},
		"null-unknown": {
			input:       NewFloat32Null(),
			candidate:   NewFloat32Unknown(),
			expectation: false,
		},
		"null-null": {
			input:       NewFloat32Null(),
			candidate:   NewFloat32Null(),
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

func TestFloat32ValueIsNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Float32Value
		expected bool
	}{
		"known": {
			input:    NewFloat32Value(2.4),
			expected: false,
		},
		"null": {
			input:    NewFloat32Null(),
			expected: true,
		},
		"unknown": {
			input:    NewFloat32Unknown(),
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

func TestFloat32ValueIsUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Float32Value
		expected bool
	}{
		"known": {
			input:    NewFloat32Value(2.4),
			expected: false,
		},
		"null": {
			input:    NewFloat32Null(),
			expected: false,
		},
		"unknown": {
			input:    NewFloat32Unknown(),
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

func TestFloat32ValueString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Float32Value
		expectation string
	}
	tests := map[string]testCase{
		"less-than-one": {
			input:       NewFloat32Value(0.12340984302980000),
			expectation: "0.123410",
		},
		"more-than-one": {
			input:       NewFloat32Value(173219.328125),
			expectation: "173219.328125",
		},
		"negative-more-than-one": {
			input:       NewFloat32Value(-0.12340984302980000),
			expectation: "-0.123410",
		},
		"negative-less-than-one": {
			input:       NewFloat32Value(-173219.328125),
			expectation: "-173219.328125",
		},
		"min-float32": {
			input:       NewFloat32Value(math.SmallestNonzeroFloat64),
			expectation: "0.000000",
		},
		"max-float32": {
			input:       NewFloat32Value(math.MaxFloat32),
			expectation: "340282346638528859811704183484516925440.000000",
		},
		"unknown": {
			input:       NewFloat32Unknown(),
			expectation: "<unknown>",
		},
		"null": {
			input:       NewFloat32Null(),
			expectation: "<null>",
		},
		"zero-value": {
			input:       Float32Value{},
			expectation: "<null>",
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

func TestFloat32ValueValueFloat64(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Float32Value
		expected float32
	}{
		"known": {
			input:    NewFloat32Value(2.4),
			expected: 2.4,
		},
		"null": {
			input:    NewFloat32Null(),
			expected: 0.0,
		},
		"unknown": {
			input:    NewFloat32Unknown(),
			expected: 0.0,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.ValueFloat32()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestFloat32ValueValueFloat32Pointer(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Float32Value
		expected *float32
	}{
		"known": {
			input:    NewFloat32Value(2.4),
			expected: pointer(float32(2.4)),
		},
		"null": {
			input:    NewFloat32Null(),
			expected: nil,
		},
		"unknown": {
			input:    NewFloat32Unknown(),
			expected: pointer(float32(0.0)),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.ValueFloat32Pointer()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestNewFloat32PointerValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		value    *float32
		expected Float32Value
	}{
		"nil": {
			value:    nil,
			expected: NewFloat32Null(),
		},
		"value": {
			value:    pointer(float32(1.2)),
			expected: NewFloat32Value(1.2),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewFloat32PointerValue(testCase.value)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestNewFloat32PointerValueOrNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		value    *float32
		expected Float32Value
	}{
		"nil": {
			value:    nil,
			expected: NewFloat32Null(),
		},
		"value": {
			value:    pointer(float32(1.2)),
			expected: NewFloat32Value(1.2),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewFloat32PointerValueOrNull(testCase.value)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestFloat32Value(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		assert.Equal(t, float32(1.23), v.Get())
	})

	t.Run("GetPtr", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		p := v.GetPtr()
		assert.NotNil(t, p)
		assert.Equal(t, float32(1.23), *p)
	})

	t.Run("Set", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		v.Set(4.56)
		assert.Equal(t, float32(4.56), v.Get())
	})

	t.Run("SetPtr", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		p := float32(4.56)
		v.SetPtr(&p)
		assert.Equal(t, float32(4.56), v.Get())

		v.SetPtr(nil)
		assert.True(t, v.IsNull())
	})

	t.Run("SetNull", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		v.SetNull()
		assert.True(t, v.Float32Value.IsNull())
	})

	t.Run("SetUnknown", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		v.SetUnknown()
		assert.True(t, v.Float32Value.IsUnknown())
	})

	t.Run("IsKnown", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		assert.True(t, v.IsKnown())
		v.SetNull()
		assert.False(t, v.IsKnown())
		v.SetUnknown()
		assert.False(t, v.IsKnown())
	})

	t.Run("NotEqual", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		v2 := NewInt32Value(1)
		assert.False(t, v.Equal(v2))
	})

	t.Run("Type", func(t *testing.T) {
		v := NewFloat32Value(1.23)
		assert.IsType(t, Float32Type{}, v.Type(context.Background()))
	})
}
