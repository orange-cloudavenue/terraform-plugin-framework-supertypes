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
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestInt32ValueToTerraformValue(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Int32Value
		expectation interface{}
	}
	tests := map[string]testCase{
		"known": {
			input:       NewInt32Value(123),
			expectation: tftypes.NewValue(tftypes.Number, big.NewFloat(123)),
		},
		"unknown": {
			input:       NewInt32Unknown(),
			expectation: tftypes.NewValue(tftypes.Number, tftypes.UnknownValue),
		},
		"null": {
			input:       NewInt32Null(),
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

func TestInt32ValueEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Int32Value
		candidate   attr.Value
		expectation bool
	}
	tests := map[string]testCase{
		"known-known-same": {
			input:       NewInt32Value(123),
			candidate:   NewInt32Value(123),
			expectation: true,
		},
		"known-known-diff": {
			input:       NewInt32Value(123),
			candidate:   NewInt32Value(456),
			expectation: false,
		},
		"known-unknown": {
			input:       NewInt32Value(123),
			candidate:   NewInt32Unknown(),
			expectation: false,
		},
		"known-null": {
			input:       NewInt32Value(123),
			candidate:   NewInt32Null(),
			expectation: false,
		},
		"unknown-value": {
			input:       NewInt32Unknown(),
			candidate:   NewInt32Value(123),
			expectation: false,
		},
		"unknown-unknown": {
			input:       NewInt32Unknown(),
			candidate:   NewInt32Unknown(),
			expectation: true,
		},
		"unknown-null": {
			input:       NewInt32Unknown(),
			candidate:   NewInt32Null(),
			expectation: false,
		},
		"null-known": {
			input:       NewInt32Null(),
			candidate:   NewInt32Value(123),
			expectation: false,
		},
		"null-unknown": {
			input:       NewInt32Null(),
			candidate:   NewInt32Unknown(),
			expectation: false,
		},
		"null-null": {
			input:       NewInt32Null(),
			candidate:   NewInt32Null(),
			expectation: true,
		},
		"not-int32-type": {
			input:       NewInt32Value(123),
			candidate:   NewStringValue("123"),
			expectation: false,
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

func TestInt32ValueIsNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Int32Value
		expected bool
	}{
		"known": {
			input:    NewInt32Value(24),
			expected: false,
		},
		"null": {
			input:    NewInt32Null(),
			expected: true,
		},
		"unknown": {
			input:    NewInt32Unknown(),
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

func TestInt32ValueIsUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Int32Value
		expected bool
	}{
		"known": {
			input:    NewInt32Value(24),
			expected: false,
		},
		"null": {
			input:    NewInt32Null(),
			expected: false,
		},
		"unknown": {
			input:    NewInt32Unknown(),
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

func TestInt32ValueString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input       Int32Value
		expectation string
	}
	tests := map[string]testCase{
		"known-less-than-one": {
			input:       NewInt32Value(-1234098430),
			expectation: "-1234098430",
		},
		"known-more-than-one": {
			input:       NewInt32Value(923879381),
			expectation: "923879381",
		},
		"known-min-int32": {
			input:       NewInt32Value(math.MinInt32),
			expectation: "-2147483648",
		},
		"known-max-int32": {
			input:       NewInt32Value(math.MaxInt32),
			expectation: "2147483647",
		},
		"unknown": {
			input:       NewInt32Unknown(),
			expectation: "<unknown>",
		},
		"null": {
			input:       NewInt32Null(),
			expectation: "<null>",
		},
		"zero-value": {
			input:       Int32Value{},
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

func TestInt32ValueValueInt32(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Int32Value
		expected int32
	}{
		"known": {
			input:    NewInt32Value(24),
			expected: 24,
		},
		"null": {
			input:    NewInt32Null(),
			expected: 0,
		},
		"unknown": {
			input:    NewInt32Unknown(),
			expected: 0,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.ValueInt32()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInt32ValueValueInt32Pointer(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    Int32Value
		expected *int32
	}{
		"known": {
			input:    NewInt32Value(24),
			expected: pointer(int32(24)),
		},
		"null": {
			input:    NewInt32Null(),
			expected: nil,
		},
		"unknown": {
			input:    NewInt32Unknown(),
			expected: pointer(int32(0)),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.ValueInt32Pointer()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestNewInt32PointerValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		value    *int32
		expected Int32Value
	}{
		"nil": {
			value:    nil,
			expected: NewInt32Null(),
		},
		"value": {
			value:    pointer(int32(123)),
			expected: NewInt32Value(123),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewInt32PointerValue(testCase.value)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInt32Value(t *testing.T) {
	// Test NewInt32Null
	v := NewInt32Null()
	assert.True(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.False(t, v.IsKnown())
	assert.Equal(t, int32(0), v.Get())

	// Test NewInt32Unknown
	v = NewInt32Unknown()
	assert.False(t, v.IsNull())
	assert.True(t, v.IsUnknown())
	assert.False(t, v.IsKnown())
	assert.Equal(t, int32(0), v.Get())

	// Test NewInt32Value
	v = NewInt32Value(42)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(42), v.Get())

	// Test NewInt32PointerValue
	i := int32(42)
	v = NewInt32PointerValue(&i)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(42), v.Get())

	// Test Set
	v.Set(84)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(84), v.Get())

	// Test SetPtr
	i = 168
	v.SetPtr(&i)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(168), v.Get())

	// Test SetInt
	v.SetInt(42)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(42), v.Get())

	// Test SetInt8
	v.SetInt8(8)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(8), v.Get())

	// Test SetInt16
	v.SetInt16(16)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(16), v.Get())

	// Test SetInt32
	v.SetInt32(32)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(32), v.Get())

	// Test SetInt64
	v.SetInt64(64)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(64), v.Get())

	// Test SetIntPtr
	iInt := int(128)
	v.SetIntPtr(&iInt)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(128), v.Get())
	assert.NotNil(t, v.GetIntPtr())

	// Test SetInt8Ptr
	i8 := int8(8)
	v.SetInt8Ptr(&i8)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(8), v.Get())
	assert.NotNil(t, v.GetInt8Ptr())

	// Test SetInt16Ptr
	i16 := int16(16)
	v.SetInt16Ptr(&i16)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(16), v.Get())
	assert.NotNil(t, v.GetInt16Ptr())

	// Test SetInt32Ptr
	i32 := int32(32)
	v.SetInt32Ptr(&i32)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(32), v.Get())
	assert.NotNil(t, v.GetInt32Ptr())

	// Test SetInt64Ptr
	i64 := int64(64)
	v.SetInt64Ptr(&i64)
	assert.False(t, v.IsNull())
	assert.False(t, v.IsUnknown())
	assert.True(t, v.IsKnown())
	assert.Equal(t, int32(64), v.Get())
	assert.NotNil(t, v.GetInt64Ptr())

	// Test GetPtr
	iPtr := v.GetPtr()
	assert.NotNil(t, iPtr)
	assert.Equal(t, int32(64), *iPtr)

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
	assert.Equal(t, int32(0), v.Get())

	// Test SetUnknown
	v.SetUnknown()
	assert.False(t, v.IsNull())
	assert.True(t, v.IsUnknown())
	assert.False(t, v.IsKnown())
	assert.Equal(t, int32(0), v.Get())

	// Test Equal
	v1 := NewInt32Value(42)
	v2 := NewInt32Value(42)
	v3 := NewInt32Value(84)
	assert.True(t, v1.Equal(v2))
	assert.False(t, v1.Equal(v3))
}
