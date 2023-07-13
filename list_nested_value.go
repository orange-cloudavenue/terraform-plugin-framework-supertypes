package supertypes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure the implementation satisfies the expected interfaces.
var _ basetypes.ListValuable = ListValue{}

type ListNestedValue struct {
	basetypes.ListValue
}

// func (v ListNestedValue) Equal(o attr.Value) bool {
// 	other, ok := o.(ListValue)

// 	if !ok {
// 		return false
// 	}

// 	return v.ListValue.Equal(other.ListValue)
// }

func (v ListNestedValue) Type(ctx context.Context) attr.Type {
	// CustomListNestedType defined in the schema type section
	return v.ListValue.Type(ctx)
}

func (v ListNestedValue) ToListValue(_ context.Context) (basetypes.ListValue, diag.Diagnostics) {
	return v.ListValue, nil
}

func ListNestedNull(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListNull(elementType),
	}
}

func ListNestedUnknown(elementType attr.Type) ListValue {
	return ListValue{
		ListValue: basetypes.NewListUnknown(elementType),
	}
}

// * CustomFunc

func (v *ListNestedValue) Get(ctx context.Context, target interface{}, allowUnhandled bool) (diag diag.Diagnostics) {
	return v.ListValue.ElementsAs(ctx, target, allowUnhandled)
}

func (v *ListNestedValue) Set(ctx context.Context, elements any) diag.Diagnostics {
	var d diag.Diagnostics
	v.ListValue, d = types.ListValueFrom(ctx, v.ElementType(ctx), elements)
	return d
}

func (v *ListNestedValue) SetNull(ctx context.Context) {
	v.ListValue = basetypes.NewListNull(v.ElementType(ctx))
}

func (v *ListNestedValue) SetUnknown(ctx context.Context) {
	v.ListValue = basetypes.NewListUnknown(v.ElementType(ctx))
}

func (v ListNestedValue) IsKnown() bool {
	return !v.ListValue.IsNull() && !v.ListValue.IsUnknown()
}
