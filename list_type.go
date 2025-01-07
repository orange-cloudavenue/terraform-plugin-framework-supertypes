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
var _ basetypes.ListTypable = ListType{}
var _ basetypes.ListTypable = ListTypeOf[struct{}]{}

// * ListType

type ListType struct {
	basetypes.ListType
}

func (t ListType) Equal(o attr.Type) bool {
	switch o.(type) {
	case ListType:
		other, ok := o.(ListType)
		if !ok {
			return false
		}

		return t.ListType.Equal(other.ListType)
	case basetypes.ObjectType:
		other, ok := o.(basetypes.ObjectType)
		if !ok {
			return false
		}

		return t.ListType.Equal(other)
	default:
		return false
	}
}

func (t ListType) String() string {
	return "supertypes.ListType[" + t.ElementType().String() + "]"
}

func (t ListType) ValueFromList(_ context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	value := ListValue{
		ListValue: in,
	}

	return value, nil
}

func (t ListType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	ListValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	ListValuable, diags := t.ValueFromList(ctx, ListValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return ListValuable, nil
}

func (t ListType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.List{
		ElementType: t.ElementType().TerraformType(ctx),
	}
}

func (t ListType) ElementType() attr.Type {
	if t.ListType.ElemType == nil {
		return missingType{}
	}

	return t.ListType.ElemType
}

func (t ListType) ValueType(ctx context.Context) attr.Value {
	return ListValue{
		ListValue: t.ListType.ValueType(ctx).(basetypes.ListValue),
	}
}

// * -----------------
// * ListTypeOf is a type that implements ListTypable for a specific type.

type ListTypeOf[T any] struct {
	basetypes.ListType
}

func NewListTypeOf[T any](ctx context.Context) ListTypeOf[T] {
	return ListTypeOf[T]{ListType: basetypes.ListType{ElemType: ElementTypeMust[T](ctx)}}
}

// Equal returns true if the given type is equal to this type.
func (t ListTypeOf[T]) Equal(o attr.Type) bool {
	switch o.(type) {
	case ListTypeOf[T]:
		other, ok := o.(ListTypeOf[T])
		if !ok {
			return false
		}

		return t.ListType.Equal(other.ListType)
	case basetypes.ListType:
		other, ok := o.(basetypes.ListType)
		if !ok {
			return false
		}

		return t.ListType.Equal(other)
	default:
		return false
	}
}

// ValueFromList converts a basetypes.ListValue to a ListValueOf.
func (t ListTypeOf[T]) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewListValueOfNull[T](ctx), diags
	}
	if in.IsUnknown() {
		return NewListValueOfUnknown[T](ctx), diags
	}

	v, d := basetypes.NewListValue(ElementTypeMust[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewListValueOfUnknown[T](ctx), diags
	}

	value := ListValueOf[T]{
		ListValue: v,
	}

	return value, diags
}

// String returns a string representation of the type.
func (t ListTypeOf[T]) String() string {
	var zero T
	return fmt.Sprintf("supertypes.ListTypeOf[%T]", zero)
}

// ValueFromTerraform converts a tftypes.Value to a ListValueOf.
func (t ListTypeOf[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	ListValue, ok := attrValue.(basetypes.ListValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	ListValuable, diags := t.ValueFromList(ctx, ListValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return ListValuable, nil
}

// ValueType returns the value type of this List.
func (t ListTypeOf[T]) ValueType(ctx context.Context) attr.Value {
	return ListValueOf[T]{}
}

// ElementType returns the element type of this List.
func (t ListTypeOf[T]) ElementType() attr.Type {
	return ElementTypeMust[T](context.Background())
}
