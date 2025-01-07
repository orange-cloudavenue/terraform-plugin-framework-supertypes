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
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.NumberValuable = NumberValue{}

type NumberValue struct {
	basetypes.NumberValue
}

func (v NumberValue) Equal(o attr.Value) bool {
	other, ok := o.(NumberValue)

	if !ok {
		return false
	}

	return v.NumberValue.Equal(other.NumberValue)
}

func (v NumberValue) Type(ctx context.Context) attr.Type {
	return NumberType{
		NumberType: v.NumberValue.Type(ctx).(basetypes.NumberType),
	}
}

func NewNumberNull() NumberValue {
	return NumberValue{
		NumberValue: basetypes.NewNumberNull(),
	}
}

func NewNumberUnknown() NumberValue {
	return NumberValue{
		NumberValue: basetypes.NewNumberUnknown(),
	}
}

func NewNumberValue(s *big.Float) NumberValue {
	return NumberValue{
		NumberValue: basetypes.NewNumberValue(s),
	}
}

// * CustomFunc

// Get returns the known Number value. If Number is null or unknown, returns 0.0.
func (v *NumberValue) Get() *big.Float {
	return v.NumberValue.ValueBigFloat()
}

// Set sets the Number value.
func (v *NumberValue) Set(s *big.Float) {

	v.NumberValue = basetypes.NewNumberValue(s)
}

// SetNull sets the Number value to null.
func (v *NumberValue) SetNull() {
	v.NumberValue = basetypes.NewNumberNull()
}

// SetUnknown sets the Number value to unknown.
func (v *NumberValue) SetUnknown() {
	v.NumberValue = basetypes.NewNumberUnknown()
}

// IsKnown returns true if the value is not null and not unknown.
func (v NumberValue) IsKnown() bool {
	return !v.NumberValue.IsNull() && !v.NumberValue.IsUnknown()
}
