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
var (
	_ basetypes.MapTypable = MapNestedType{}

	_ basetypes.MapTypable = MapNestedObjectTypeOf[struct{}]{}
	_ NestedObjectType     = MapNestedObjectTypeOf[struct{}]{}
)

// MapNestedType is the attribute type of a MapNestedValue.
type MapNestedType struct {
	basetypes.MapType
}

// MapNestedObjectTypeOf is the attribute type of a MapNestedObjectValueOf.
type MapNestedObjectTypeOf[T any] struct {
	basetypes.MapType
}

func (t MapNestedObjectTypeOf[T]) Equal(o attr.Type) bool {
	switch o.(type) {
	case MapNestedType:
		other, ok := o.(MapNestedType)
		if !ok {
			return false
		}

		return t.MapType.Equal(other.MapType)
	case basetypes.ObjectType:
		other, ok := o.(basetypes.ObjectType)
		if !ok {
			return false
		}

		return t.MapType.Equal(other)
	case MapNestedObjectTypeOf[T]:
		other, ok := o.(MapNestedObjectTypeOf[T])
		if !ok {
			return false
		}

		return t.MapType.Equal(other.MapType)
	default:
		return false
	}
}

// String returns a string representation of the type.
func (t MapNestedType) String() string {
	return "supertypes.MapType[" + t.ElementType().String() + "]"
}

func (t MapNestedType) Equal(o attr.Type) bool {
	switch o.(type) {
	case MapNestedType:
		other, ok := o.(MapNestedType)
		if !ok {
			return false
		}

		return t.MapType.Equal(other.MapType)
	case basetypes.MapType:
		other, ok := o.(basetypes.MapType)
		if !ok {
			return false
		}

		return t.MapType.Equal(other)
	default:
		return false
	}
}

func (t MapNestedObjectTypeOf[T]) String() string {
	var zero T
	return fmt.Sprintf("MapNestedObjectTypeOf[%T]", zero)
}

func (t MapNestedType) ValueFromMap(_ context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	// TODO improve this
	value := MapNestedValue{
		MapValue: in,
	}

	return value, nil
}

func (t MapNestedObjectTypeOf[T]) ValueFromMap(ctx context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewMapNestedObjectValueOfNull[T](ctx), diags
	}
	if in.IsUnknown() {
		return NewMapNestedObjectValueOfUnknown[T](ctx), diags
	}

	listValue, d := basetypes.NewMapValue(NewObjectTypeOf[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewMapNestedObjectValueOfUnknown[T](ctx), diags
	}

	value := MapNestedObjectValueOf[T]{
		MapValue: listValue,
	}

	return value, diags
}

func (t MapNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.MapType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	MapValue, ok := attrValue.(basetypes.MapValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	MapValuable, diags := t.ValueFromMap(ctx, MapValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting MapValue to MapValuable: %v", diags)
	}

	return MapValuable, nil
}

func (t MapNestedObjectTypeOf[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.MapType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	listValue, ok := attrValue.(basetypes.MapValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromMap(ctx, listValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting MapValue to MapValuable: %v", diags)
	}

	return listValuable, nil
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the AttributeType to something Terraform
// can understand.
func (t MapNestedType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Map{
		ElementType: t.MapType.ElementType().TerraformType(ctx),
	}
}

func (t MapNestedType) ElementType() attr.Type {
	if t.MapType.ElemType == nil {
		return missingType{}
	}

	return t.MapType.ElemType
}

func (t MapNestedObjectTypeOf[T]) ValueType(ctx context.Context) attr.Value {
	return MapNestedObjectValueOf[T]{}
}

func (t MapNestedObjectTypeOf[T]) NewObjectPtr(ctx context.Context) (any, diag.Diagnostics) {
	return nestedObjectTypeNewObjectPtr[T](ctx)
}

func (t MapNestedObjectTypeOf[T]) NewObjectSlice(ctx context.Context, _, _ int) (any, diag.Diagnostics) {
	return nestedMapTypeNewMap[T](ctx)
}
func (t MapNestedObjectTypeOf[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	return NewMapNestedObjectValueOfNull[T](ctx), diags
}

func (t MapNestedObjectTypeOf[T]) ValueFromObjectPtr(ctx context.Context, ptr any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v, ok := ptr.(map[string]*T); ok {
		return NewMapNestedObjectValueOfPtr(ctx, v), diags
	}

	diags.Append(diag.NewErrorDiagnostic("Invalid map value", fmt.Sprintf("incorrect type: want %T, got %T", (map[string]*T)(nil), ptr)))
	return nil, diags
}
func (t MapNestedObjectTypeOf[T]) ValueFromObjectSlice(ctx context.Context, slice any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v, ok := slice.(map[string]*T); ok {
		return NewMapNestedObjectValueOfMap(ctx, v), diags
	}
	diags.Append(diag.NewErrorDiagnostic("Invalid slice value", fmt.Sprintf("incorrect type: want %T, got %T", (map[string]*T)(nil), slice)))
	return nil, diags
}
