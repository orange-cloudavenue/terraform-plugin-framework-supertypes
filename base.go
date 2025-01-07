package supertypes

import (
	// These are all the packages that are used in the templates.
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	// Used for the templates.
	_ "github.com/fatih/color"
	_ "github.com/iancoleman/strcase"
)

const (
	errorTestValidate = "An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. Please report the following to the provider developer:\n\n"
	errorTestConvert  = "An unexpected error was encountered trying to convert the Terraform value. This is always an error in the provider. Please report the following to the provider developer:\n\n"
)

func nestedObjectTypeNewObjectPtr[T any](_ context.Context) (*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	return new(T), diags
}

func nestedObjectTypeNewObjectSlice[T any](_ context.Context, xlen, xcap int) ([]*T, diag.Diagnostics) { //nolint:unparam
	var diags diag.Diagnostics

	return make([]*T, xlen, xcap), diags
}

func nestedMapTypeNewMap[T any](_ context.Context) (map[string]*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	return make(map[string]*T), diags
}

func nestedObjectValueSlice[T any](ctx context.Context, val valueWithElements) ([]*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	elements := val.Elements()
	n := len(elements)
	slice := make([]*T, n)
	for i := 0; i < n; i++ {
		ptr, d := nestedObjectValueObjectPtrFromElement[T](ctx, elements[i])
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}

		slice[i] = ptr
	}

	return slice, diags
}

func nestedObjectValueMap[T any](ctx context.Context, val valueWithElementsMap) (map[string]*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	elements := val.Elements()
	m := make(map[string]*T)
	for e := range elements {
		ptr, d := nestedObjectValueObjectPtrFromElement[T](ctx, elements[e])
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}

		m[e] = ptr
	}

	return m, diags
}

func nestedObjectValueObjectPtrFromElement[T any](ctx context.Context, val attr.Value) (*T, diag.Diagnostics) {
	var diags diag.Diagnostics

	ptr := new(T)
	diags.Append(val.(ObjectValueOf[T]).ObjectValue.As(ctx, ptr, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil, diags
	}

	return ptr, diags
}

//go:generate go run template.go
