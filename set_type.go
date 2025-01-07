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
var _ basetypes.SetTypable = SetType{}
var _ basetypes.SetTypable = SetTypeOf[struct{}]{}

// * SetType

type SetType struct {
	basetypes.SetType
}

func (t SetType) Equal(o attr.Type) bool {
	switch o.(type) {
	case SetType:
		other, ok := o.(SetType)
		if !ok {
			return false
		}

		return t.SetType.Equal(other.SetType)
	case basetypes.ObjectType:
		other, ok := o.(basetypes.ObjectType)
		if !ok {
			return false
		}

		return t.SetType.Equal(other)
	default:
		return false
	}
}

func (t SetType) String() string {
	return "supertypes.SetType[" + t.ElementType().String() + "]"
}

func (t SetType) ValueFromSet(_ context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	value := SetValue{
		SetValue: in,
	}

	return value, nil
}

func (t SetType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.SetType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	SetValue, ok := attrValue.(basetypes.SetValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	SetValuable, diags := t.ValueFromSet(ctx, SetValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return SetValuable, nil
}

func (t SetType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Set{
		ElementType: t.ElementType().TerraformType(ctx),
	}
}

func (t SetType) ElementType() attr.Type {
	if t.SetType.ElemType == nil {
		return missingType{}
	}

	return t.SetType.ElemType
}

func (t SetType) ValueType(ctx context.Context) attr.Value {
	return SetValue{
		SetValue: t.SetType.ValueType(ctx).(basetypes.SetValue),
	}
}

// * -----------------
// * SetTypeOf is a type that implements SetTypable for a specific type.

type SetTypeOf[T any] struct {
	basetypes.SetType
}

func NewSetTypeOf[T any](ctx context.Context) SetTypeOf[T] {
	return SetTypeOf[T]{SetType: basetypes.SetType{ElemType: ElementTypeMust[T](ctx)}}
}

// Equal returns true if the given type is equal to this type.
func (t SetTypeOf[T]) Equal(o attr.Type) bool {
	switch o.(type) {
	case SetTypeOf[T]:
		other, ok := o.(SetTypeOf[T])
		if !ok {
			return false
		}

		return t.SetType.Equal(other.SetType)
	case basetypes.SetType:
		other, ok := o.(basetypes.SetType)
		if !ok {
			return false
		}

		return t.SetType.Equal(other)
	default:
		return false
	}
}

// ValueFromSet converts a basetypes.SetValue to a SetValueOf.
func (t SetTypeOf[T]) ValueFromSet(ctx context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewSetValueOfNull[T](ctx), diags
	}
	if in.IsUnknown() {
		return NewSetValueOfUnknown[T](ctx), diags
	}

	v, d := basetypes.NewSetValue(ElementTypeMust[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewSetValueOfUnknown[T](ctx), diags
	}

	value := SetValueOf[T]{
		SetValue: v,
	}

	return value, diags
}

// String returns a string representation of the type.
func (t SetTypeOf[T]) String() string {
	var zero T
	return fmt.Sprintf("supertypes.SetTypeOf[%T]", zero)
}

// ValueFromTerraform converts a tftypes.Value to a SetValueOf.
func (t SetTypeOf[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.SetType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	SetValue, ok := attrValue.(basetypes.SetValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	SetValuable, diags := t.ValueFromSet(ctx, SetValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return SetValuable, nil
}

// ValueType returns the value type of this Set.
func (t SetTypeOf[T]) ValueType(ctx context.Context) attr.Value {
	return SetValueOf[T]{}
}

// ElementType returns the element type of this Set.
func (t SetTypeOf[T]) ElementType() attr.Type {
	return ElementTypeMust[T](context.Background())
}
