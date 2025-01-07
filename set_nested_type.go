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
	_ basetypes.SetTypable = SetNestedType{}

	_ basetypes.SetTypable = SetNestedObjectTypeOf[struct{}]{}
	_ NestedObjectType     = SetNestedObjectTypeOf[struct{}]{}
)

// SetNestedType is the attribute type of a SetNestedValue.
type SetNestedType struct {
	basetypes.SetType
}

// SetNestedObjectTypeOf is the attribute type of a SetNestedObjectValueOf.
type SetNestedObjectTypeOf[T any] struct {
	basetypes.SetType
}

func (t SetNestedObjectTypeOf[T]) Equal(o attr.Type) bool {
	switch o.(type) {
	case SetNestedType:
		other, ok := o.(SetNestedType)
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
	case SetNestedObjectTypeOf[T]:
		other, ok := o.(SetNestedObjectTypeOf[T])
		if !ok {
			return false
		}

		return t.SetType.Equal(other.SetType)
	default:
		return false
	}
}

// String returns a string representation of the type.
func (t SetNestedType) String() string {
	return "supertypes.SetType[" + t.ElementType().String() + "]"
}

func (t SetNestedType) Equal(o attr.Type) bool {
	switch o.(type) {
	case SetNestedType:
		other, ok := o.(SetNestedType)
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

func (t SetNestedObjectTypeOf[T]) String() string {
	var zero T
	return fmt.Sprintf("SetNestedObjectTypeOf[%T]", zero)
}

func (t SetNestedType) ValueFromSet(_ context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	// TODO improve this
	value := SetNestedValue{
		SetValue: in,
	}

	return value, nil
}

func (t SetNestedObjectTypeOf[T]) ValueFromSet(ctx context.Context, in basetypes.SetValue) (basetypes.SetValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewSetNestedObjectValueOfNull[T](ctx), diags
	}
	if in.IsUnknown() {
		return NewSetNestedObjectValueOfUnknown[T](ctx), diags
	}

	listValue, d := basetypes.NewSetValue(NewObjectTypeOf[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewSetNestedObjectValueOfUnknown[T](ctx), diags
	}

	value := SetNestedObjectValueOf[T]{
		SetValue: listValue,
	}

	return value, diags
}

func (t SetNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t SetNestedObjectTypeOf[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.SetType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	listValue, ok := attrValue.(basetypes.SetValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromSet(ctx, listValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting SetValue to SetValuable: %v", diags)
	}

	return listValuable, nil
}

// TerraformType returns the tftypes.Type that should be used to
// represent this type. This constrains what user input will be
// accepted and what kind of data can be set in state. The framework
// will use this to translate the AttributeType to something Terraform
// can understand.
func (t SetNestedType) TerraformType(ctx context.Context) tftypes.Type {
	return tftypes.Set{
		ElementType: t.SetType.ElementType().TerraformType(ctx),
	}
}

func (t SetNestedType) ElementType() attr.Type {
	if t.SetType.ElemType == nil {
		return missingType{}
	}

	return t.SetType.ElemType
}

func (t SetNestedObjectTypeOf[T]) ValueType(ctx context.Context) attr.Value {
	return SetNestedObjectValueOf[T]{}
}

func (t SetNestedObjectTypeOf[T]) NewObjectPtr(ctx context.Context) (any, diag.Diagnostics) {
	return nestedObjectTypeNewObjectPtr[T](ctx)
}

func (t SetNestedObjectTypeOf[T]) NewObjectSlice(ctx context.Context, len, cap int) (any, diag.Diagnostics) {
	return nestedObjectTypeNewObjectSlice[T](ctx, len, cap)
}
func (t SetNestedObjectTypeOf[T]) NullValue(ctx context.Context) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	return NewSetNestedObjectValueOfNull[T](ctx), diags
}

func (t SetNestedObjectTypeOf[T]) ValueFromObjectPtr(ctx context.Context, ptr any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v, ok := ptr.(*T); ok {
		return NewSetNestedObjectValueOfPtr(ctx, v), diags
	}

	diags.Append(diag.NewErrorDiagnostic("Invalid pointer value", fmt.Sprintf("incorrect type: want %T, got %T", (*T)(nil), ptr)))
	return nil, diags
}
func (t SetNestedObjectTypeOf[T]) ValueFromObjectSlice(ctx context.Context, slice any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v, ok := slice.([]*T); ok {
		return NewSetNestedObjectValueOfSlice(ctx, v), diags
	}
	diags.Append(diag.NewErrorDiagnostic("Invalid slice value", fmt.Sprintf("incorrect type: want %T, got %T", (*[]T)(nil), slice)))
	return nil, diags
}
