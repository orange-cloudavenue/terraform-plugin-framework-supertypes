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
var _ basetypes.Int32Typable = Int32Type{}

type Int32Type struct {
	basetypes.Int32Type
}

func (t Int32Type) Equal(o attr.Type) bool {
	other, ok := o.(Int32Type)
	if !ok {
		return false
	}

	return t.Int32Type.Equal(other.Int32Type)
}

func (t Int32Type) String() string {
	return "supertypes.Int32Type"
}

func (t Int32Type) ValueFromInt32(_ context.Context, in basetypes.Int32Value) (basetypes.Int32Valuable, diag.Diagnostics) {
	value := Int32Value{
		Int32Value: in,
	}

	return value, nil
}

func (t Int32Type) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.Int32Type.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	Int32Value, ok := attrValue.(basetypes.Int32Value)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	Int32Valuable, diags := t.ValueFromInt32(ctx, Int32Value)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting Int32Value to Int32Valuable: %v", diags)
	}

	return Int32Valuable, nil
}

func (t Int32Type) ValueType(ctx context.Context) attr.Value {
	return Int32Value{
		Int32Value: t.Int32Type.ValueType(ctx).(basetypes.Int32Value),
	}
}
