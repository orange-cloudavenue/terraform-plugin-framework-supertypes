/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// -------------------------------------------------------------------- //
// !      DO NOT EDIT. This file is auto-generated from template      ! //
// -------------------------------------------------------------------- //
package supertypes

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.Float64Valuable = Float64Value{}

type Float64Value struct {
	basetypes.Float64Value
}

func (v Float64Value) Equal(o attr.Value) bool {
	other, ok := o.(Float64Value)

	if !ok {
		return false
	}

	return v.Float64Value.Equal(other.Float64Value)
}

func (v Float64Value) Type(ctx context.Context) attr.Type {
	return Float64Type{
		Float64Type: v.Float64Value.Type(ctx).(basetypes.Float64Type),
	}
}

func NewFloat64Null() Float64Value {
	return Float64Value{
		Float64Value: basetypes.NewFloat64Null(),
	}
}

func NewFloat64Unknown() Float64Value {
	return Float64Value{
		Float64Value: basetypes.NewFloat64Unknown(),
	}
}

func NewFloat64Value(s float64) Float64Value {
	return Float64Value{
		Float64Value: basetypes.NewFloat64Value(s),
	}
}

func NewFloat64PointerValue(s *float64) Float64Value {
	return Float64Value{
		Float64Value: basetypes.NewFloat64PointerValue(s),
	}
}

func NewFloat64PointerValueOrNull(s *float64) Float64Value {
	if s == nil {
		return NewFloat64Null()
	}

	return Float64Value{
		Float64Value: basetypes.NewFloat64PointerValue(s),
	}
}

// * CustomFunc

// Get returns the known Float64 value. If Float64 is null or unknown, returns 0.0.
func (v *Float64Value) Get() float64 {
	return v.Float64Value.ValueFloat64()
}

// GetPtr returns a pointer to the known int64 value, nil for a
// null value, or a pointer to 0 for an unknown value.
func (v *Float64Value) GetPtr() *float64 {
	return v.Float64Value.ValueFloat64Pointer()
}

// Set sets the Float64 value.
func (v *Float64Value) Set(s float64) {

	v.Float64Value = basetypes.NewFloat64Value(s)
}

// SetPtr sets a pointer to the Float64 value.
func (v *Float64Value) SetPtr(s *float64) {
	if s == nil {
		v.Float64Value = basetypes.NewFloat64Null()
		return
	}

	v.Float64Value = basetypes.NewFloat64PointerValue(s)
}

// SetNull sets the Float64 value to null.
func (v *Float64Value) SetNull() {
	v.Float64Value = basetypes.NewFloat64Null()
}

// SetUnknown sets the Float64 value to unknown.
func (v *Float64Value) SetUnknown() {
	v.Float64Value = basetypes.NewFloat64Unknown()
}

// IsKnown returns true if the value is not null and not unknown.
func (v Float64Value) IsKnown() bool {
	return !v.Float64Value.IsNull() && !v.Float64Value.IsUnknown()
}
