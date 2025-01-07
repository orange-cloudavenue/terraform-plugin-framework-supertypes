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
var _ basetypes.Float64Typable = Float64Type{}

type Float64Type struct {
	basetypes.Float64Type
}

func (t Float64Type) Equal(o attr.Type) bool {
	other, ok := o.(Float64Type)
	if !ok {
		return false
	}

	return t.Float64Type.Equal(other.Float64Type)
}

func (t Float64Type) String() string {
	return "supertypes.Float64Type"
}

func (t Float64Type) ValueFromFloat64(_ context.Context, in basetypes.Float64Value) (basetypes.Float64Valuable, diag.Diagnostics) {
	value := Float64Value{
		Float64Value: in,
	}

	return value, nil
}

func (t Float64Type) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.Float64Type.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	Float64Value, ok := attrValue.(basetypes.Float64Value)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	Float64Valuable, diags := t.ValueFromFloat64(ctx, Float64Value)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting Float64Value to Float64Valuable: %v", diags)
	}

	return Float64Valuable, nil
}

func (t Float64Type) ValueType(ctx context.Context) attr.Value {
	return Float64Value{
		Float64Value: t.Float64Type.ValueType(ctx).(basetypes.Float64Value),
	}
}
