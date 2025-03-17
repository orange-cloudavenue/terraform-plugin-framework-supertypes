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
var _ basetypes.Float32Valuable = Float32Value{}

type Float32Value struct {
	basetypes.Float32Value
}

func (v Float32Value) Equal(o attr.Value) bool {
	other, ok := o.(Float32Value)

	if !ok {
		return false
	}

	return v.Float32Value.Equal(other.Float32Value)
}

func (v Float32Value) Type(ctx context.Context) attr.Type {
	return Float32Type{
		Float32Type: v.Float32Value.Type(ctx).(basetypes.Float32Type),
	}
}

func NewFloat32Null() Float32Value {
	return Float32Value{
		Float32Value: basetypes.NewFloat32Null(),
	}
}

func NewFloat32Unknown() Float32Value {
	return Float32Value{
		Float32Value: basetypes.NewFloat32Unknown(),
	}
}

func NewFloat32Value(s float32) Float32Value {
	return Float32Value{
		Float32Value: basetypes.NewFloat32Value(s),
	}
}

func NewFloat32PointerValue(s *float32) Float32Value {
	return Float32Value{
		Float32Value: basetypes.NewFloat32PointerValue(s),
	}
}

func NewFloat32PointerValueOrNull(s *float32) Float32Value {
	if s == nil {
		return NewFloat32Null()
	}

	return Float32Value{
		Float32Value: basetypes.NewFloat32PointerValue(s),
	}
}

// * CustomFunc

// Get returns the known Float32 value.
func (v *Float32Value) Get() float32 {
	return v.Float32Value.ValueFloat32()
}

// GetPtr returns a pointer to the known int64 value, nil for a
// null value, or a pointer to 0 for an unknown value.
func (v *Float32Value) GetPtr() *float32 {
	return v.Float32Value.ValueFloat32Pointer()
}

// Set sets the Float32 value.
func (v *Float32Value) Set(s float32) {

	v.Float32Value = basetypes.NewFloat32Value(s)
}

// SetPtr sets a pointer to the Float32 value.
func (v *Float32Value) SetPtr(s *float32) {
	if s == nil {
		v.Float32Value = basetypes.NewFloat32Null()
		return
	}

	v.Float32Value = basetypes.NewFloat32PointerValue(s)
}

// SetNull sets the Float32 value to null.
func (v *Float32Value) SetNull() {
	v.Float32Value = basetypes.NewFloat32Null()
}

// SetUnknown sets the Float32 value to unknown.
func (v *Float32Value) SetUnknown() {
	v.Float32Value = basetypes.NewFloat32Unknown()
}

// IsKnown returns true if the value is not null and not unknown.
func (v Float32Value) IsKnown() bool {
	return !v.Float32Value.IsNull() && !v.Float32Value.IsUnknown()
}
