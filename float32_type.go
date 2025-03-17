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
var _ basetypes.Float32Typable = Float32Type{}

type Float32Type struct {
	basetypes.Float32Type
}

func (t Float32Type) Equal(o attr.Type) bool {
	other, ok := o.(Float32Type)
	if !ok {
		return false
	}

	return t.Float32Type.Equal(other.Float32Type)
}

func (t Float32Type) String() string {
	return "supertypes.Float32Type"
}

func (t Float32Type) ValueFromFloat32(_ context.Context, in basetypes.Float32Value) (basetypes.Float32Valuable, diag.Diagnostics) {
	value := Float32Value{
		Float32Value: in,
	}

	return value, nil
}

func (t Float32Type) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.Float32Type.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	Float32Value, ok := attrValue.(basetypes.Float32Value)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	Float32Valuable, diags := t.ValueFromFloat32(ctx, Float32Value)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting Float32Value to Float32Valuable: %v", diags)
	}

	return Float32Valuable, nil
}

func (t Float32Type) ValueType(ctx context.Context) attr.Value {
	return Float32Value{
		Float32Value: t.Float32Type.ValueType(ctx).(basetypes.Float32Value),
	}
}
