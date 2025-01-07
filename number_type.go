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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.NumberTypable = NumberType{}

type NumberType struct {
	basetypes.NumberType
}

func (t NumberType) Equal(o attr.Type) bool {
	other, ok := o.(NumberType)
	if !ok {
		return false
	}

	return t.NumberType.Equal(other.NumberType)
}

func (t NumberType) String() string {
	return "supertypes.NumberType"
}

func (t NumberType) ValueFromNumber(_ context.Context, in basetypes.NumberValue) (basetypes.NumberValuable, diag.Diagnostics) {
	value := NumberValue{
		NumberValue: in,
	}

	return value, nil
}

func (t NumberType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.NumberType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	NumberValue, ok := attrValue.(basetypes.NumberValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	NumberValuable, diags := t.ValueFromNumber(ctx, NumberValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting NumberValue to NumberValuable: %v", diags)
	}

	return NumberValuable, nil
}

func (t NumberType) ValueType(ctx context.Context) attr.Value {
	return NumberValue{
		NumberValue: t.NumberType.ValueType(ctx).(basetypes.NumberValue),
	}
}
