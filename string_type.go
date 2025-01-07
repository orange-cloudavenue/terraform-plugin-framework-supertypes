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
var _ basetypes.StringTypable = StringType{}

type StringType struct {
	basetypes.StringType
}

func (t StringType) Equal(o attr.Type) bool {
	other, ok := o.(StringType)
	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t StringType) String() string {
	return "supertypes.StringType"
}

func (t StringType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	value := StringValue{
		StringValue: in,
	}

	return value, nil
}

func (t StringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	StringValue, ok := attrValue.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	StringValuable, diags := t.ValueFromString(ctx, StringValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return StringValuable, nil
}

func (t StringType) ValueType(ctx context.Context) attr.Value {
	return StringValue{
		StringValue: t.StringType.ValueType(ctx).(basetypes.StringValue),
	}
}
