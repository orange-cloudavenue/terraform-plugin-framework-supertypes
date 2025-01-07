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
	_ basetypes.ListTypable = ListNestedType{}

	_ basetypes.ListTypable = ListNestedObjectTypeOf[struct{}]{}
	_ NestedObjectType      = ListNestedObjectTypeOf[struct{}]{}
)

// ListNestedType is the attribute type of a ListNestedValue.
type ListNestedType struct {
	basetypes.ListType
}

// ListNestedObjectTypeOf is the attribute type of a ListNestedObjectValueOf.
type ListNestedObjectTypeOf[T any] struct {
	basetypes.ListType
}

func (t ListNestedObjectTypeOf[T]) Equal(o attr.Type) bool {
	switch o.(type) {
	case ListNestedType:
		other, ok := o.(ListNestedType)
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
	case ListNestedObjectTypeOf[T]:
		other, ok := o.(ListNestedObjectTypeOf[T])
		if !ok {
			return false
		}

		return t.ListType.Equal(other.ListType)
	default:
		return false
	}
}

// String returns a string representation of the type.
func (t ListNestedType) String() string {
	return "supertypes.ListType[" + t.ElementType().String() + "]"
}

func (t ListNestedType) Equal(o attr.Type) bool {
	switch o.(type) {
	case ListNestedType:
		other, ok := o.(ListNestedType)
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

func (t ListNestedObjectTypeOf[T]) String() string {
	var zero T
	return fmt.Sprintf("ListNestedObjectTypeOf[%T]", zero)
}

func (t ListNestedType) ValueFromList(_ context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	// TODO improve this
	value := ListNestedValue{
		ListValue: in,
	}

	return value, nil
}

func (t ListNestedObjectTypeOf[T]) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewListNestedObjectValueOfNull[T](ctx), diags
	}
	if in.IsUnknown() {
		return NewListNestedObjectValueOfUnknown[T](ctx), diags
	}

	listValue, d := basetypes.NewListValue(NewObjectTypeOf[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewListNestedObjectValueOfUnknown[T](ctx), diags
	}

	value := ListNestedObjectValueOf[T]{
		ListValue: listValue,
	}

	return value, diags
}

func (t ListNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t ListNestedObjectTypeOf[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	listValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromList(ctx, listValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the AttributeType to something Terraform
// can understand.
func (t ListNestedType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.List{
		ElementType: t.ListType.ElementType().TerraformType(ctx),
	}
}

func (t ListNestedType) ElementType() attr.Type {
	if t.ListType.ElemType == nil {
		return missingType{}
	}

	return t.ListType.ElemType
}

func (t ListNestedObjectTypeOf[T]) ValueType(ctx context.Context) attr.Value {
	return ListNestedObjectValueOf[T]{}
}

func (t ListNestedObjectTypeOf[T]) NewObjectPtr(ctx context.Context) (any, diag.Diagnostics) {
	return nestedObjectTypeNewObjectPtr[T](ctx)
}

func (t ListNestedObjectTypeOf[T]) NewObjectSlice(ctx context.Context, len, cap int) (any, diag.Diagnostics) {
	return nestedObjectTypeNewObjectSlice[T](ctx, len, cap)
}
func (t ListNestedObjectTypeOf[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	return NewListNestedObjectValueOfNull[T](ctx), diags
}

func (t ListNestedObjectTypeOf[T]) ValueFromObjectPtr(ctx context.Context, ptr any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v, ok := ptr.(*T); ok {
		return NewListNestedObjectValueOfPtr(ctx, v), diags
	}

	diags.Append(diag.NewErrorDiagnostic("Invalid pointer value", fmt.Sprintf("incorrect type: want %T, got %T", (*T)(nil), ptr)))
	return nil, diags
}
func (t ListNestedObjectTypeOf[T]) ValueFromObjectSlice(ctx context.Context, slice any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v, ok := slice.([]*T); ok {
		return NewListNestedObjectValueOfSlice(ctx, v), diags
	}
	diags.Append(diag.NewErrorDiagnostic("Invalid slice value", fmt.Sprintf("incorrect type: want %T, got %T", (*[]T)(nil), slice)))
	return nil, diags
}
