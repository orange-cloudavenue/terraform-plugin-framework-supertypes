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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ObjectValuable = SingleNestedValue{}

type SingleNestedValue struct {
	basetypes.ObjectValue
}

func (v SingleNestedValue) Type(ctx context.Context) attr.Type {
	return v.ObjectValue.Type(ctx)
}

func (v SingleNestedValue) ToObjectValue(_ context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return v.ObjectValue, nil
}

func NewSingleNestedNull(attributeTypes map[string]attr.Type) SingleNestedValue {
	return SingleNestedValue{
		ObjectValue: basetypes.NewObjectNull(attributeTypes),
	}
}

func NewSingleNestedUnknown(attributeTypes map[string]attr.Type) SingleNestedValue {
	return SingleNestedValue{
		ObjectValue: basetypes.NewObjectUnknown(attributeTypes),
	}
}

// * CustomFunc

func (v *SingleNestedValue) Get(ctx context.Context, target interface{}, opts basetypes.ObjectAsOptions) (diag diag.Diagnostics) {
	return v.ObjectValue.As(ctx, target, opts)
}

func (v *SingleNestedValue) Set(ctx context.Context, structure any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ObjectValue, d = types.ObjectValueFrom(ctx, v.AttributeTypes(ctx), structure)
	return d
}

func (v *SingleNestedValue) SetNull(ctx context.Context) {
	v.ObjectValue = basetypes.NewObjectNull(v.AttributeTypes(ctx))
}

func (v *SingleNestedValue) SetUnknown(ctx context.Context) {
	v.ObjectValue = basetypes.NewObjectUnknown(v.AttributeTypes(ctx))
}

func (v SingleNestedValue) IsKnown() bool {
	return !v.ObjectValue.IsNull() && !v.ObjectValue.IsUnknown()
}

// SingleNestedObjectValueOf represents a Terraform Plugin Framework Single value whose corresponding Go type is the structure T.
type SingleNestedObjectValueOf[T any] struct {
	basetypes.ObjectValue
}

var _ basetypes.ObjectValuable = SingleNestedObjectValueOf[struct{}]{}

func (v SingleNestedObjectValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(SingleNestedObjectValueOf[T])

	if !ok {
		return false
	}

	return v.ObjectValue.Equal(other.ObjectValue)
}

func (v SingleNestedObjectValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewSingleNestedObjectTypeOf[T](ctx)
}

func NewSingleNestedObjectValueOfNull[T any](ctx context.Context) SingleNestedObjectValueOf[T] {
	return SingleNestedObjectValueOf[T]{ObjectValue: basetypes.NewObjectNull(AttributeTypesMust[T](ctx))}
}

func NewSingleNestedObjectValueOfUnknown[T any](ctx context.Context) SingleNestedObjectValueOf[T] {
	return SingleNestedObjectValueOf[T]{ObjectValue: basetypes.NewObjectUnknown(AttributeTypesMust[T](ctx))}
}

func NewSingleNestedObjectValueOf[T any](ctx context.Context, t *T) SingleNestedObjectValueOf[T] {
	return SingleNestedObjectValueOf[T]{ObjectValue: MustDiag(basetypes.NewObjectValueFrom(ctx, AttributeTypesMust[T](ctx), t))}
}

func (v SingleNestedObjectValueOf[T]) Get(ctx context.Context) (*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	ptr := new(T)

	diags.Append(v.ObjectValue.As(ctx, ptr, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil, diags
	}

	return ptr, diags
}

func (v *SingleNestedObjectValueOf[T]) Set(ctx context.Context, t *T) (diags diag.Diagnostics) {
	v.ObjectValue, diags = basetypes.NewObjectValueFrom(ctx, AttributeTypesMust[T](ctx), t)
	return diags
}

// IsKnown returns whether the value is known.
func (v SingleNestedObjectValueOf[T]) IsKnown() bool {
	if !v.IsNull() && !v.IsUnknown() {
		return true
	}

	return false
}

func newSingleNestedObjectValueOf[T any](ctx context.Context, t *T) SingleNestedObjectValueOf[T] {
	return SingleNestedObjectValueOf[T]{ObjectValue: MustDiag(basetypes.NewObjectValueFrom(ctx, AttributeTypesMust[T](ctx), t))}
}

func (v *SingleNestedObjectValueOf[T]) SetNull(ctx context.Context) {
	*v = NewSingleNestedObjectValueOfNull[T](ctx)
}

func (v *SingleNestedObjectValueOf[T]) SetUnknown(ctx context.Context) {
	*v = NewSingleNestedObjectValueOfUnknown[T](ctx)
}
