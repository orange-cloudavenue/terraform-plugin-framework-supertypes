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
	"math"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestInt64ValueToTerraformValue(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Int64Value
		expectation interface{}
	}
	tests := map[string]testCase{
		"known": {
			input:       NewInt64Value(123),
			expectation: tftypes.NewValue(tftypes.Number, big.NewFloat(123)),
		},
		"unknown": {
			input:       NewInt64Unknown(),
			expectation: tftypes.NewValue(tftypes.Number, tftypes.UnknownValue),
		},
		"null": {
			input:       NewInt64Null(),
			expectation: tftypes.NewValue(tftypes.Number, nil),
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
			if !cmp.Equal(got, test.expectation, cmp.Comparer(numberComparer)) {
				t.Errorf("Expected %+v, got %+v", test.expectation, got)
			}
		})
	}
}

func TestInt64ValueEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Int64Value
		candidate   attr.Value
		expectation bool
	}
	tests := map[string]testCase{
		"known-known-same": {
			input:       NewInt64Value(123),
			candidate:   NewInt64Value(123),
			expectation: true,
		},
		"known-known-diff": {
			input:       NewInt64Value(123),
			candidate:   NewInt64Value(456),
			expectation: false,
		},
		"known-unknown": {
			input:       NewInt64Value(123),
			candidate:   NewInt64Unknown(),
			expectation: false,
		},
		"known-null": {
			input:       NewInt64Value(123),
			candidate:   NewInt64Null(),
			expectation: false,
		},
		"unknown-value": {
			input:       NewInt64Unknown(),
			candidate:   NewInt64Value(123),
			expectation: false,
		},
		"unknown-unknown": {
			input:       NewInt64Unknown(),
			candidate:   NewInt64Unknown(),
			expectation: true,
		},
		"unknown-null": {
			input:       NewInt64Unknown(),
			candidate:   NewInt64Null(),
			expectation: false,
		},
		"null-known": {
			input:       NewInt64Null(),
			candidate:   NewInt64Value(123),
			expectation: false,
		},
		"null-unknown": {
			input:       NewInt64Null(),
			candidate:   NewInt64Unknown(),
			expectation: false,
		},
		"null-null": {
			input:       NewInt64Null(),
			candidate:   NewInt64Null(),
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

func TestInt64ValueIsNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Int64Value
		expected bool
	}{
		"known": {
			input:    NewInt64Value(24),
			expected: false,
		},
		"null": {
			input:    NewInt64Null(),
			expected: true,
		},
		"unknown": {
			input:    NewInt64Unknown(),
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

func TestInt64ValueIsUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Int64Value
		expected bool
	}{
		"known": {
			input:    NewInt64Value(24),
			expected: false,
		},
		"null": {
			input:    NewInt64Null(),
			expected: false,
		},
		"unknown": {
			input:    NewInt64Unknown(),
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

func TestInt64ValueString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Int64Value
		expectation string
	}
	tests := map[string]testCase{
		"known-less-than-one": {
			input:       NewInt64Value(-12340984302980000),
			expectation: "-12340984302980000",
		},
		"known-more-than-one": {
			input:       NewInt64Value(92387938173219327),
			expectation: "92387938173219327",
		},
		"known-min-int64": {
			input:       NewInt64Value(math.MinInt64),
			expectation: "-9223372036854775808",
		},
		"known-max-int64": {
			input:       NewInt64Value(math.MaxInt64),
			expectation: "9223372036854775807",
		},
		"unknown": {
			input:       NewInt64Unknown(),
			expectation: "<unknown>",
		},
		"null": {
			input:       NewInt64Null(),
			expectation: "<null>",
		},
		"zero-value": {
			input:       Int64Value{},
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

func TestInt64ValueValueInt64(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Int64Value
		expected int64
	}{
		"known": {
			input:    NewInt64Value(24),
			expected: 24,
		},
		"null": {
			input:    NewInt64Null(),
			expected: 0,
		},
		"unknown": {
			input:    NewInt64Unknown(),
			expected: 0,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.ValueInt64()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInt64ValueValueInt64Pointer(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Int64Value
		expected *int64
	}{
		"known": {
			input:    NewInt64Value(24),
			expected: pointer(int64(24)),
		},
		"null": {
			input:    NewInt64Null(),
			expected: nil,
		},
		"unknown": {
			input:    NewInt64Unknown(),
			expected: pointer(int64(0)),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.ValueInt64Pointer()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestNewInt64PointerValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		value    *int64
		expected Int64Value
	}{
		"nil": {
			value:    nil,
			expected: NewInt64Null(),
		},
		"value": {
			value:    pointer(int64(123)),
			expected: NewInt64Value(123),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewInt64PointerValue(testCase.value)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInt64Value(t *testing.T) {
	// Test NewInt64Null
	v := NewInt64Null()
	assert.True(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.False(t, v.IsKnown())
	assert.Equal(t, int64(0), v.Get())

	// Test NewInt64Unknown
	v = NewInt64Unknown()
	assert.False(t, v.IsNull())
	assert.True(t, v.IsUnknown())
	assert.False(t, v.IsKnown())
	assert.Equal(t, int64(0), v.Get())

	// Test NewInt64Value
	v = NewInt64Value(42)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(42), v.Get())

	// Test NewInt64PointerValue
	i := int64(42)
	v = NewInt64PointerValue(&i)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(42), v.Get())

	// Test Set
	v.Set(84)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(84), v.Get())

	// Test SetPtr
	i = 168
	v.SetPtr(&i)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(168), v.Get())

	// Test SetInt
	v.SetInt(42)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(42), v.Get())

	// Test SetInt8
	v.SetInt8(8)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(8), v.Get())

	// Test SetInt16
	v.SetInt16(16)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(16), v.Get())

	// Test SetInt32
	v.SetInt32(32)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(32), v.Get())

	// Test SetInt64
	v.SetInt64(64)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(64), v.Get())

	// Test SetIntPtr
	iInt := int(128)
	v.SetIntPtr(&iInt)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(128), v.Get())
	assert.NotNil(t, v.GetIntPtr())

	// Test SetInt8Ptr
	i8 := int8(8)
	v.SetInt8Ptr(&i8)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(8), v.Get())
	assert.NotNil(t, v.GetInt8Ptr())

	// Test SetInt16Ptr
	i16 := int16(16)
	v.SetInt16Ptr(&i16)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(16), v.Get())
	assert.NotNil(t, v.GetInt16Ptr())

	// Test SetInt32Ptr
	i32 := int32(32)
	v.SetInt32Ptr(&i32)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(32), v.Get())
	assert.NotNil(t, v.GetInt32Ptr())

	// Test SetInt64Ptr
	i64 := int64(64)
	v.SetInt64Ptr(&i64)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int64(64), v.Get())
	assert.NotNil(t, v.GetInt64Ptr())

	// Test GetPtr
	iPtr := v.GetPtr()
	assert.NotNil(t, iPtr)
	assert.Equal(t, int64(64), *iPtr)

	// Test GetInt
	assert.Equal(t, 64, v.GetInt())

	// Test GetInt8
	assert.Equal(t, int8(64), v.GetInt8())

	// Test GetInt16
	assert.Equal(t, int16(64), v.GetInt16())

	// Test GetInt32
	assert.Equal(t, int32(64), v.GetInt32())

	// Test GetInt64
	assert.Equal(t, int64(64), v.GetInt64())

	// Test GetIntPtr
	iIntPtr := v.GetIntPtr()
	assert.NotNil(t, iIntPtr)
	assert.Equal(t, 64, *iIntPtr)

	// Test GetInt8Ptr
	i8Ptr := v.GetInt8Ptr()
	assert.NotNil(t, i8Ptr)
	assert.Equal(t, int8(64), *i8Ptr)

	// Test GetInt16Ptr
	i16Ptr := v.GetInt16Ptr()
	assert.NotNil(t, i16Ptr)
	assert.Equal(t, int16(64), *i16Ptr)

	// Test GetInt32Ptr
	i32Ptr := v.GetInt32Ptr()
	assert.NotNil(t, i32Ptr)
	assert.Equal(t, int32(64), *i32Ptr)

	// Test GetInt64Ptr
	i64Ptr := v.GetInt64Ptr()
	assert.NotNil(t, i64Ptr)
	assert.Equal(t, int64(64), *i64Ptr)

	v.SetNull()

	// Test GetPtr is nil
	iPtr = v.GetPtr()
	assert.Nil(t, iPtr)

	// Test GetIntPtr is nil
	iIntPtr = v.GetIntPtr()
	assert.Nil(t, iIntPtr)

	// Test GetInt8Ptr is nil
	i8Ptr = v.GetInt8Ptr()
	assert.Nil(t, i8Ptr)

	// Test GetInt16Ptr is nil
	i16Ptr = v.GetInt16Ptr()
	assert.Nil(t, i16Ptr)

	// Test GetInt32Ptr is nil
	i32Ptr = v.GetInt32Ptr()
	assert.Nil(t, i32Ptr)

	// Test GetInt64Ptr is nil
	i64Ptr = v.GetInt64Ptr()
	assert.Nil(t, i64Ptr)

	// Set Ptr to nil
	v.SetPtr(nil)
	assert.True(t, v.IsNull())

	// Set IntPtr to nil
	v.SetIntPtr(nil)
	assert.True(t, v.IsNull())

	// Set Int8Ptr to nil
	v.SetInt8Ptr(nil)
	assert.True(t, v.IsNull())

	// Set Int16Ptr to nil
	v.SetInt16Ptr(nil)
	assert.True(t, v.IsNull())

	// Set Int32Ptr to nil
	v.SetInt32Ptr(nil)
	assert.True(t, v.IsNull())

	// Set Int64Ptr to nil
	v.SetInt64Ptr(nil)
	assert.True(t, v.IsNull())

	// Test SetNull
	v.SetNull()
	assert.True(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.False(t, v.IsKnown())
	assert.Equal(t, int64(0), v.Get())

	// Test SetUnknown
	v.SetUnknown()
	assert.False(t, v.IsNull())
	assert.True(t, v.IsUnknown())
	assert.False(t, v.IsKnown())
	assert.Equal(t, int64(0), v.Get())

	// Test Equal
	v1 := NewInt64Value(42)
	v2 := NewInt64Value(42)
	v3 := NewInt64Value(84)
	assert.True(t, v1.Equal(v2))
	assert.False(t, v1.Equal(v3))
}
