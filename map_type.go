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
var _ basetypes.MapTypable = MapType{}
var _ basetypes.MapTypable = MapTypeOf[struct{}]{}

// * MapType

type MapType struct {
	basetypes.MapType
}

func (t MapType) Equal(o attr.Type) bool {
	switch o.(type) {
	case MapType:
		other, ok := o.(MapType)
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
	default:
		return false
	}
}

func (t MapType) String() string {
	return "supertypes.MapType[" + t.ElementType().String() + "]"
}

func (t MapType) ValueFromMap(_ context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	value := MapValue{
		MapValue: in,
	}

	return value, nil
}

func (t MapType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t MapType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Map{
		ElementType: t.ElementType().TerraformType(ctx),
	}
}

func (t MapType) ElementType() attr.Type {
	if t.MapType.ElemType == nil {
		return missingType{}
	}

	return t.MapType.ElemType
}

func (t MapType) ValueType(ctx context.Context) attr.Value {
	return MapValue{
		MapValue: t.MapType.ValueType(ctx).(basetypes.MapValue),
	}
}

// * -----------------
// * MapTypeOf is a type that implements MapTypable for a specific type.

type MapTypeOf[T any] struct {
	basetypes.MapType
}

func NewMapTypeOf[T any](ctx context.Context) MapTypeOf[T] {
	return MapTypeOf[T]{MapType: basetypes.MapType{ElemType: ElementTypeMust[T](ctx)}}
}

// Equal returns true if the given type is equal to this type.
func (t MapTypeOf[T]) Equal(o attr.Type) bool {
	switch o.(type) {
	case MapTypeOf[T]:
		other, ok := o.(MapTypeOf[T])
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

// ValueFromMap converts a basetypes.MapValue to a MapValueOf.
func (t MapTypeOf[T]) ValueFromMap(ctx context.Context, in basetypes.MapValue) (basetypes.MapValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewMapValueOfNull[T](ctx), diags
	}
	if in.IsUnknown() {
		return NewMapValueOfUnknown[T](ctx), diags
	}

	v, d := basetypes.NewMapValue(ElementTypeMust[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewMapValueOfUnknown[T](ctx), diags
	}

	value := MapValueOf[T]{
		MapValue: v,
	}

	return value, diags
}

// String returns a string representation of the type.
func (t MapTypeOf[T]) String() string {
	var zero T
	return fmt.Sprintf("supertypes.MapTypeOf[%T]", zero)
}

// ValueFromTerraform converts a tftypes.Value to a MapValueOf.
func (t MapTypeOf[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

// ValueType returns the value type of this Map.
func (t MapTypeOf[T]) ValueType(ctx context.Context) attr.Value {
	return MapValueOf[T]{}
}

// ElementType returns the element type of this Map.
func (t MapTypeOf[T]) ElementType() attr.Type {
	return ElementTypeMust[T](context.Background())
}
