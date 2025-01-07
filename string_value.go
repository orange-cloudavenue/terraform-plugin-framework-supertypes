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
var _ basetypes.StringValuable = StringValue{}

type StringValue struct {
	basetypes.StringValue
}

func (v StringValue) Equal(o attr.Value) bool {
	other, ok := o.(StringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v StringValue) Type(ctx context.Context) attr.Type {
	return StringType{
		StringType: v.StringValue.Type(ctx).(basetypes.StringType),
	}
}

func NewStringNull() StringValue {
	return StringValue{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewStringUnknown() StringValue {
	return StringValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewStringValue(s string) StringValue {
	return StringValue{
		StringValue: basetypes.NewStringValue(s),
	}
}

func NewStringPointerValue(s *string) StringValue {
	return StringValue{
		StringValue: basetypes.NewStringPointerValue(s),
	}
}

// * CustomFunc

// Get returns the known String value. If String is null or unknown, returns "".
func (v *StringValue) Get() string {
	return v.StringValue.ValueString()
}

// GetPtr returns a pointer to the known int64 value, nil for a
// null value, or a pointer to 0 for an unknown value.
func (v *StringValue) GetPtr() *string {
	return v.StringValue.ValueStringPointer()
}

// Set sets the String value.
func (v *StringValue) Set(s string) {

	if s == "" {
		v.StringValue = basetypes.NewStringNull()
		return
	}

	v.StringValue = basetypes.NewStringValue(s)
}

// SetPtr sets a pointer to the String value.
func (v *StringValue) SetPtr(s *string) {
	if s == nil {
		v.StringValue = basetypes.NewStringNull()
		return
	}

	v.StringValue = basetypes.NewStringPointerValue(s)
}

// SetNull sets the String value to null.
func (v *StringValue) SetNull() {
	v.StringValue = basetypes.NewStringNull()
}

// SetUnknown sets the String value to unknown.
func (v *StringValue) SetUnknown() {
	v.StringValue = basetypes.NewStringUnknown()
}

// IsKnown returns true if the value is not null and not unknown.
func (v StringValue) IsKnown() bool {
	return !v.StringValue.IsNull() && !v.StringValue.IsUnknown()
}
