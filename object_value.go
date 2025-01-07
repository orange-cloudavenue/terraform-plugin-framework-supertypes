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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ObjectValuable = ObjectValue{}

type ObjectValue struct {
	basetypes.ObjectValue
}

func (v ObjectValue) Equal(o attr.Value) bool {
	other, ok := o.(ObjectValue)

	if !ok {
		return false
	}

	return v.ObjectValue.Equal(other.ObjectValue)
}

func (v ObjectValue) Type(_ context.Context) attr.Type {
	return ObjectType{}
}

func NewObjectNull(attributeTypes map[string]attr.Type) ObjectValue {
	return ObjectValue{
		ObjectValue: basetypes.NewObjectNull(attributeTypes),
	}
}

func NewObjectUnknown(attributeTypes map[string]attr.Type) ObjectValue {
	return ObjectValue{
		ObjectValue: basetypes.NewObjectUnknown(attributeTypes),
	}
}

// * CustomFunc

func (v *ObjectValue) Get(ctx context.Context, target interface{}, opts basetypes.ObjectAsOptions) diag.Diagnostics {
	return v.ObjectValue.As(ctx, target, opts)
}

func (v *ObjectValue) Set(ctx context.Context, structure any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ObjectValue, d = basetypes.NewObjectValueFrom(ctx, v.AttributeTypes(ctx), structure)
	return d
}

func (v *ObjectValue) SetNull(ctx context.Context) {
	v.ObjectValue = basetypes.NewObjectNull(v.AttributeTypes(ctx))
}

func (v *ObjectValue) SetUnknown(ctx context.Context) {
	v.ObjectValue = basetypes.NewObjectUnknown(v.AttributeTypes(ctx))
}

func (v ObjectValue) IsKnown() bool {
	return !v.ObjectValue.IsNull() && !v.ObjectValue.IsUnknown()
}

// ObjectValueOf represents a Terraform Plugin Framework Object value whose corresponding Go type is the structure T.
type ObjectValueOf[T any] struct {
	basetypes.ObjectValue
}

var _ basetypes.ObjectValuable = ObjectValueOf[struct{}]{}

func (v ObjectValueOf[T]) Equal(o attr.Value) bool {
	other, ok := o.(ObjectValueOf[T])

	if !ok {
		return false
	}

	return v.ObjectValue.Equal(other.ObjectValue)
}

func (v ObjectValueOf[T]) Type(ctx context.Context) attr.Type {
	return NewObjectTypeOf[T](ctx)
}

func NewObjectValueOfNull[T any](ctx context.Context) ObjectValueOf[T] {
	return ObjectValueOf[T]{ObjectValue: basetypes.NewObjectNull(AttributeTypesMust[T](ctx))}
}

func NewObjectValueOfUnknown[T any](ctx context.Context) ObjectValueOf[T] {
	return ObjectValueOf[T]{ObjectValue: basetypes.NewObjectUnknown(AttributeTypesMust[T](ctx))}
}

func NewObjectValueOf[T any](ctx context.Context, t *T) ObjectValueOf[T] {
	return ObjectValueOf[T]{ObjectValue: MustDiag(basetypes.NewObjectValueFrom(ctx, AttributeTypesMust[T](ctx), t))}
}

func (v ObjectValueOf[T]) Get(ctx context.Context) (*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	ptr := new(T)

	diags.Append(v.ObjectValue.As(ctx, ptr, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil, diags
	}

	return ptr, diags
}

func (v *ObjectValueOf[T]) Set(ctx context.Context, t *T) (diags diag.Diagnostics) {
	v.ObjectValue, diags = basetypes.NewObjectValueFrom(ctx, AttributeTypesMust[T](ctx), t)
	return diags
}

// IsKnown returns whether the value is known.
func (v ObjectValueOf[T]) IsKnown() bool {
	if !v.IsNull() && !v.IsUnknown() {
		return true
	}

	return false
}

func (v *ObjectValueOf[T]) SetNull(ctx context.Context) {
	*v = NewObjectValueOfNull[T](ctx)
}

func (v *ObjectValueOf[T]) SetUnknown(ctx context.Context) {
	*v = NewObjectValueOfUnknown[T](ctx)
}
