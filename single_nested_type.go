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
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ basetypes.ObjectTypable = SingleNestedType{}
	_ basetypes.ObjectTypable = SingleNestedObjectTypeOf[struct{}]{}
)

type SingleNestedType struct {
	basetypes.ObjectType
}

func (t SingleNestedType) Equal(o attr.Type) bool {
	switch o.(type) {
	case SingleNestedType:
		other, ok := o.(SingleNestedType)
		if !ok {
			return false
		}

		return t.ObjectType.Equal(other.ObjectType)
	case basetypes.ObjectType:
		other, ok := o.(basetypes.ObjectType)
		if !ok {
			return false
		}

		return t.ObjectType.Equal(other)
	default:
		return false
	}
}

func (t SingleNestedType) String() string {
	var res strings.Builder
	res.WriteString("supertypes.SingleNestedType[")
	keys := make([]string, 0, len(t.AttrTypes))
	for k := range t.AttrTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for pos, key := range keys {
		if pos != 0 {
			res.WriteString(", ")
		}
		res.WriteString(`"` + key + `":`)
		res.WriteString(t.AttrTypes[key].String())
	}
	res.WriteString("]")
	return res.String()
}

func (t SingleNestedType) ValueFromObject(_ context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	// CustomSingleValue defined in the value type section
	value := SingleNestedValue{
		ObjectValue: in,
	}

	return value, nil
}

func (t SingleNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ObjectType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	ObjectValue, ok := attrValue.(basetypes.ObjectValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	value, diags := t.ValueFromObject(ctx, ObjectValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SingleValue to SingleValuable: %v", diags)
	}

	return value, nil
}

// * TypeOf

type SingleNestedObjectTypeOf[T any] struct {
	basetypes.ObjectType
}

func NewSingleNestedObjectTypeOf[T any](ctx context.Context) SingleNestedObjectTypeOf[T] {
	return SingleNestedObjectTypeOf[T]{basetypes.ObjectType{AttrTypes: AttributeTypesMust[T](ctx)}}
}

func (t SingleNestedObjectTypeOf[T]) Equal(o attr.Type) bool {
	switch o.(type) {
	case SingleNestedType:
		other, ok := o.(SingleNestedType)
		if !ok {
			return false
		}

		return t.ObjectType.Equal(other.ObjectType)
	case basetypes.ObjectType:
		other, ok := o.(basetypes.ObjectType)
		if !ok {
			return false
		}

		return t.ObjectType.Equal(other)
	case SingleNestedObjectTypeOf[T]:

		other, ok := o.(SingleNestedObjectTypeOf[T])
		if !ok {
			return false
		}

		return t.ObjectType.Equal(other.ObjectType)
	default:
		return false
	}
}

func (t SingleNestedObjectTypeOf[T]) String() string {
	var zero T
	return fmt.Sprintf("SingleNestedObjectTypeOf[%T]", zero)
}

func (t SingleNestedObjectTypeOf[T]) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewSingleNestedObjectValueOfNull[T](ctx), diags
	}
	if in.IsUnknown() {
		return NewSingleNestedObjectValueOfUnknown[T](ctx), diags
	}

	objectValue, d := basetypes.NewObjectValue(AttributeTypesMust[T](ctx), in.Attributes())
	diags.Append(d...)
	if diags.HasError() {
		return NewSingleNestedObjectValueOfUnknown[T](ctx), diags
	}

	value := SingleNestedObjectValueOf[T]{
		ObjectValue: objectValue,
	}

	return value, diags
}

func (t SingleNestedObjectTypeOf[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ObjectType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	objectValue, ok := attrValue.(basetypes.ObjectValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	objectValuable, diags := t.ValueFromObject(ctx, objectValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ObjectValue to SingleNestedValuable: %v", diags)
	}

	return objectValuable, nil
}

func (t SingleNestedObjectTypeOf[T]) ValueType(ctx context.Context) attr.Value {
	return SingleNestedObjectValueOf[T]{}
}
