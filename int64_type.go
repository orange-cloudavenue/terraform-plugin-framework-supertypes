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
var _ basetypes.Int64Typable = Int64Type{}

type Int64Type struct {
	basetypes.Int64Type
}

func (t Int64Type) Equal(o attr.Type) bool {
	other, ok := o.(Int64Type)
	if !ok {
		return false
	}

	return t.Int64Type.Equal(other.Int64Type)
}

func (t Int64Type) String() string {
	return "supertypes.Int64Type"
}

func (t Int64Type) ValueFromInt64(_ context.Context, in basetypes.Int64Value) (basetypes.Int64Valuable, diag.Diagnostics) {
	value := Int64Value{
		Int64Value: in,
	}

	return value, nil
}

func (t Int64Type) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.Int64Type.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	Int64Value, ok := attrValue.(basetypes.Int64Value)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	Int64Valuable, diags := t.ValueFromInt64(ctx, Int64Value)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting Int64Value to Int64Valuable: %v", diags)
	}

	return Int64Valuable, nil
}

func (t Int64Type) ValueType(ctx context.Context) attr.Value {
	return Int64Value{
		Int64Value: t.Int64Type.ValueType(ctx).(basetypes.Int64Value),
	}
}
