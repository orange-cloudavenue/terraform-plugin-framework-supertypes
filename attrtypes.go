/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package supertypes

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// AttributeTypes returns a map of attribute types for the specified type T.
// T must be a struct and reflection is used to find exported fields of T with the `tfsdk` tag.
func AttributeTypes[T any](ctx context.Context) (map[string]attr.Type, error) {
	var t T
	val := reflect.ValueOf(t)
	typ := val.Type()

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%T has unsupported type: %s", t, typ)
	}

	attributeTypes := make(map[string]attr.Type)
	for i := range typ.NumField() {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue // Skip unexported fields.
		}
		tag := field.Tag.Get(`tfsdk`)
		if tag == "-" {
			continue // Skip explicitly excluded fields.
		}
		if tag == "" {
			return nil, fmt.Errorf(`%T needs a struct tag for "tfsdk" on %s`, t, field.Name)
		}

		if v, ok := val.Field(i).Interface().(attr.Value); ok {
			attributeTypes[tag] = v.Type(ctx)
		}
	}

	return attributeTypes, nil
}

func AttributeTypesMust[T any](ctx context.Context) map[string]attr.Type {
	return Must(AttributeTypes[T](ctx))
}

// ElementType returns the element type of the specified type T.
// T must be a slice or map and reflection is used to find the element type.
func ElementType[T any](ctx context.Context) (attr.Type, error) {
	var t T

	supportedTypes := []attr.Type{
		basetypes.StringType{},
		basetypes.BoolType{},
		basetypes.Int32Type{},
		basetypes.Int64Type{},
		basetypes.Float32Type{},
		basetypes.Float64Type{},
	}

	if v, ok := any(t).(attr.Value); ok {
		vType := v.Type(ctx)

		vTypeReflect := reflect.TypeOf(vType)
		for _, supportedType := range supportedTypes {
			supportedTypeReflect := reflect.TypeOf(supportedType)

			// Check if types match directly or if embedded type matches
			if vTypeReflect == supportedTypeReflect ||
				(vTypeReflect.Kind() == reflect.Struct && vTypeReflect.NumField() > 0 &&
					vTypeReflect.Field(0).Type == supportedTypeReflect) {
				return vType, nil
			}
		}

		return nil, fmt.Errorf("%T has unsupported type: %s", t, vType.String())
	}

	val := reflect.ValueOf(t)
	typ := val.Type()

	switch typ.Kind() {
	case reflect.String:
		return StringType{}, nil
	case reflect.Bool:
		return BoolType{}, nil
	case reflect.Int32:
		return Int32Type{}, nil
	case reflect.Int64:
		return Int64Type{}, nil
	case reflect.Float32:
		return Float32Type{}, nil
	case reflect.Float64:
		return Float64Type{}, nil
	default:
		return nil, fmt.Errorf("%T has unsupported type: %s", t, typ)
	}
}

func ElementTypeMust[T any](ctx context.Context) attr.Type {
	return Must(ElementType[T](ctx))
}
