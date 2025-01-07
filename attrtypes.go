package supertypes

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	for i := 0; i < typ.NumField(); i++ {
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
func ElementType[T any](_ context.Context) (attr.Type, error) {
	var t T
	val := reflect.ValueOf(t)
	typ := val.Type()

	switch typ.Kind() {
	case reflect.String:
		return StringType{}, nil
	case reflect.Bool:
		return BoolType{}, nil
	case reflect.Int64:
		return Int64Type{}, nil
	case reflect.Float64:
		return Float64Type{}, nil
	default:
		return nil, fmt.Errorf("%T has unsupported type: %s", t, typ)
	}
}

func ElementTypeMust[T any](ctx context.Context) attr.Type {
	return Must(ElementType[T](ctx))
}
