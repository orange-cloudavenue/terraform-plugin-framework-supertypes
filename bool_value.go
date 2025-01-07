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
var _ basetypes.BoolValuable = BoolValue{}

type BoolValue struct {
	basetypes.BoolValue
}

func (v BoolValue) Equal(o attr.Value) bool {
	other, ok := o.(BoolValue)

	if !ok {
		return false
	}

	return v.BoolValue.Equal(other.BoolValue)
}

func (v BoolValue) Type(ctx context.Context) attr.Type {
	return BoolType{
		BoolType: v.BoolValue.Type(ctx).(basetypes.BoolType),
	}
}

func NewBoolNull() BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolNull(),
	}
}

func NewBoolUnknown() BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolUnknown(),
	}
}

func NewBoolValue(s bool) BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolValue(s),
	}
}

func NewBoolPointerValue(s *bool) BoolValue {
	return BoolValue{
		BoolValue: basetypes.NewBoolPointerValue(s),
	}
}

// * CustomFunc

// Get returns the known Bool value. If Bool is null or unknown, returns false.
func (v *BoolValue) Get() bool {
	return v.BoolValue.ValueBool()
}

// GetPtr returns a pointer to the known int64 value, nil for a
// null value, or a pointer to 0 for an unknown value.
func (v *BoolValue) GetPtr() *bool {
	return v.BoolValue.ValueBoolPointer()
}

// Set sets the Bool value.
func (v *BoolValue) Set(s bool) {

	v.BoolValue = basetypes.NewBoolValue(s)
}

// SetPtr sets a pointer to the Bool value.
func (v *BoolValue) SetPtr(s *bool) {
	if s == nil {
		v.BoolValue = basetypes.NewBoolNull()
		return
	}

	v.BoolValue = basetypes.NewBoolPointerValue(s)
}

// SetNull sets the Bool value to null.
func (v *BoolValue) SetNull() {
	v.BoolValue = basetypes.NewBoolNull()
}

// SetUnknown sets the Bool value to unknown.
func (v *BoolValue) SetUnknown() {
	v.BoolValue = basetypes.NewBoolUnknown()
}

// IsKnown returns true if the value is not null and not unknown.
func (v BoolValue) IsKnown() bool {
	return !v.BoolValue.IsNull() && !v.BoolValue.IsUnknown()
}
