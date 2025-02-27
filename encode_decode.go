package supertypes

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func DecodeToStruct[T any](ctx context.Context, tfStruct any, goStruct *T) (*T, error) {
	valTf, typeTf := reflectStruct(tfStruct)
	valGo, typeGo := reflectStruct(goStruct)

	if typeTf.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%T is not a struct. Actual type is unsupported: %s", tfStruct, typeTf)
	}

	if typeGo.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%T is not a struct. Actual type is unsupported: %s", goStruct, typeGo)
	}

	if err := decodeToStruct(ctx, typeTf, valTf, typeGo, valGo); err != nil {
		return nil, err
	}

	return goStruct, nil
}

func EncodeToTerraformStruct[T any](ctx context.Context, tfStruct *T, goStruct any) (*T, error) {
	valTf, typeTf := reflectStruct(tfStruct)
	valGo, typeGo := reflectStruct(goStruct)

	if typeTf.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%T is not a struct. Actual type is unsupported: %s", tfStruct, typeTf)
	}

	if typeGo.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%T is not a struct. Actual type is unsupported: %s", goStruct, typeGo)
	}

	if err := encodeToTerraformStruct(ctx, typeTf, valTf, typeGo, valGo); err != nil {
		return nil, err
	}

	return tfStruct, nil
}

func encodeToTerraformStruct(ctx context.Context, rTfType reflect.Type, rTf reflect.Value, rGoType reflect.Type, rGo reflect.Value) error { //nolint:gocyclo
	if rTf.Kind() == reflect.Ptr && rTf.IsNil() {
		return nil
	}

	if rGo.Kind() == reflect.Ptr && rGo.IsNil() {
		return nil
	}

	if rTfType.Kind() == reflect.Ptr {
		rTf = rTf.Elem()
		rTfType = rTfType.Elem()
	}

	if rGoType.Kind() == reflect.Ptr {
		rGo = rGo.Elem()
		rGoType = rGoType.Elem()
	}

	if rTfType.Kind() != reflect.Struct {
		return fmt.Errorf("%T is not a struct. Actual type is unsupported: %s", rTf.Interface(), rTfType)
	}

	if rGoType.Kind() != reflect.Struct {
		return fmt.Errorf("%T is not a struct. Actual type is unsupported: %s", rGo.Interface(), rGoType)
	}

	for i := 0; i < rTfType.NumField(); i++ {
		fieldTf := rTfType.Field(i)
		valTfField := rTf.Field(i)
		valGoField := rGo.FieldByName(fieldTf.Name)

		if valGoField.IsValid() {
			switch v := valTfField.Interface().(type) {
			case StringValue, BoolValue:
				set(valTfField, valGoField)
			case Int64Value:
				switch valGoField.Kind() {
				case reflect.Int:
					v.SetInt(valGoField.Interface().(int))
				case reflect.Int8:
					v.SetInt8(valGoField.Interface().(int8))
				case reflect.Int16:
					v.SetInt16(valGoField.Interface().(int16))
				case reflect.Int32:
					v.SetInt32(valGoField.Interface().(int32))
				case reflect.Int64:
					v.SetInt64(valGoField.Interface().(int64))
				case reflect.Ptr:
					switch valGoField.Type().Elem().Kind() {
					case reflect.Int:
						v.SetIntPtr(valGoField.Interface().(*int))
					case reflect.Int8:
						v.SetInt8Ptr(valGoField.Interface().(*int8))
					case reflect.Int16:
						v.SetInt16Ptr(valGoField.Interface().(*int16))
					case reflect.Int32:
						v.SetInt32Ptr(valGoField.Interface().(*int32))
					case reflect.Int64:
						v.SetInt64Ptr(valGoField.Interface().(*int64))
					}
				}
				valTfField.Set(reflect.ValueOf(v))
			case Float64Value:
				switch valGoField.Kind() {
				case reflect.Float32:
					v.SetFloat32(valGoField.Interface().(float32))
				case reflect.Float64:
					v.SetFloat64(valGoField.Interface().(float64))
				case reflect.Ptr:
					switch valGoField.Type().Elem().Kind() {
					case reflect.Float32:
						v.SetFloat32Ptr(valGoField.Interface().(*float32))
					case reflect.Float64:
						v.SetFloat64Ptr(valGoField.Interface().(*float64))
					}
				}
				valTfField.Set(reflect.ValueOf(v))
			default:
				// Default is not simple type. Require to get type in the goStruct to initialize the generic type in the tfStruct
				// Define if is a slice, map or a struct in the goStruct
				switch valGoField.Kind() {
				case reflect.Map:

					switch valGoField.Type().Elem().Kind() {
					case reflect.String, reflect.Int64, reflect.Float64:
						setOf(ctx, valTfField, valGoField)
					case reflect.Ptr:

						// Get the Tf value object by getOf method
						TfV := getOf(ctx, valTfField)
						for _, k := range valGoField.MapKeys() {
							if !TfV.MapIndex(k).IsValid() {
								TfV.SetMapIndex(k, reflect.New(TfV.Type().Elem().Elem()))
							}

							if err := encodeToTerraformStruct(ctx, TfV.MapIndex(k).Type(), TfV.MapIndex(k), valGoField.MapIndex(k).Type(), valGoField.MapIndex(k)); err != nil {
								return err
							}
						}

						setOf(ctx, valTfField, TfV)
					}
					// End of Map
				case reflect.Slice:

					switch valGoField.Type().Elem().Kind() {
					case reflect.String, reflect.Int64, reflect.Float64:
						setOf(ctx, valTfField, valGoField)
					case reflect.Ptr:

						// Get the Tf value object by getOf method
						TfV := getOf(ctx, valTfField)
						for j := 0; j < valGoField.Len(); j++ {
							if !TfV.Index(j).IsValid() {
								TfV.Index(j).Set(reflect.New(TfV.Type().Elem().Elem()))
							}
							if err := encodeToTerraformStruct(ctx, TfV.Index(j).Type(), TfV.Index(j), valGoField.Index(j).Type(), valGoField.Index(j)); err != nil {
								return err
							}
						}

						setOf(ctx, valTfField, TfV)
					}
					// End of Slice
				case reflect.Ptr:

					// Get the Tf value object by getOf method
					TfV := getOf(ctx, valTfField)
					if err := encodeToTerraformStruct(ctx, TfV.Type(), TfV, valGoField.Type(), valGoField); err != nil {
						return err
					}
					setOf(ctx, valTfField, TfV)
					// End of Ptr
				}
			}
		}
	}
	return nil
}

func decodeToStruct(ctx context.Context, rTfType reflect.Type, rTf reflect.Value, rGoType reflect.Type, rGo reflect.Value) error { //nolint:gocyclo
	if rTf.Kind() == reflect.Ptr && rTf.IsNil() {
		// TODO set rGo to nil
		return nil
	}

	if rGo.Kind() == reflect.Ptr && rGo.IsNil() {
		return nil
	}

	if rTfType.Kind() == reflect.Ptr {
		rTf = rTf.Elem()
		rTfType = rTfType.Elem()
	}

	if rGoType.Kind() == reflect.Ptr {
		rGo = rGo.Elem()
		rGoType = rGoType.Elem()
	}

	if rTfType.Kind() != reflect.Struct {
		return fmt.Errorf("%T is not a struct. Actual type is unsupported: %s", rTf.Interface(), rTfType)
	}

	if rGoType.Kind() != reflect.Struct {
		return fmt.Errorf("%T is not a struct. Actual type is unsupported: %s", rGo.Interface(), rGoType)
	}

	for i := 0; i < rTfType.NumField(); i++ {
		fieldTf := rTfType.Field(i)

		valTfField := rTf.Field(i)
		// if rGo.Kind() == reflect.Ptr {
		// 	rGo = rGo.Elem()
		// }

		if valTfField.IsZero() {
			continue
		}

		// Terraform attribute is null or unknown
		x := valTfField.MethodByName("IsKnown").Call([]reflect.Value{})
		if !x[0].Bool() {
			// TODO unset field in go struct
			continue
		}

		valGoField := rGo.FieldByName(fieldTf.Name)

		if valGoField.IsValid() {
			switch v := valTfField.Interface().(type) {
			case StringValue:
				// if the field in the goStruct is a pointer to a string
				if valGoField.Kind() == reflect.Ptr {
					valGoField.Set(reflect.ValueOf(v.GetPtr()))
				} else {
					valGoField.Set(reflect.ValueOf(v.Get()))
				}
			case BoolValue:
				// if the field in the goStruct is a pointer to a bool
				if valGoField.Kind() == reflect.Ptr {
					valGoField.Set(reflect.ValueOf(v.GetPtr()))
				} else {
					valGoField.Set(reflect.ValueOf(v.Get()))
				}
			case Int64Value:
				// In the goStruct the field maybe int,int8,int16,int32,int64,*int,*int8,*int16,*int32,*int64
				switch valGoField.Kind() {
				case reflect.Int:
					valGoField.Set(reflect.ValueOf(v.GetInt()))
				case reflect.Int8:
					valGoField.Set(reflect.ValueOf(v.GetInt8()))
				case reflect.Int16:
					valGoField.Set(reflect.ValueOf(v.GetInt16()))
				case reflect.Int32:
					valGoField.Set(reflect.ValueOf(v.GetInt32()))
				case reflect.Int64:
					valGoField.Set(reflect.ValueOf(v.GetInt64()))
				case reflect.Ptr:
					switch valGoField.Type().Elem().Kind() {
					case reflect.Int:
						valGoField.Set(reflect.ValueOf(v.GetIntPtr()))
					case reflect.Int8:
						valGoField.Set(reflect.ValueOf(v.GetInt8Ptr()))
					case reflect.Int16:
						valGoField.Set(reflect.ValueOf(v.GetInt16Ptr()))
					case reflect.Int32:
						valGoField.Set(reflect.ValueOf(v.GetInt32Ptr()))
					case reflect.Int64:
						valGoField.Set(reflect.ValueOf(v.GetInt64Ptr()))
					}
				}
			case Float64Value:
				// In the goStruct the field maybe float32,float64,*float32,*float64
				switch valGoField.Kind() {
				case reflect.Float32:
					valGoField.Set(reflect.ValueOf(v.GetFloat32()))
				case reflect.Float64:
					valGoField.Set(reflect.ValueOf(v.GetFloat64()))
				case reflect.Ptr:
					switch valGoField.Type().Elem().Kind() {
					case reflect.Float32:
						valGoField.Set(reflect.ValueOf(v.GetFloat32Ptr()))
					case reflect.Float64:
						valGoField.Set(reflect.ValueOf(v.GetFloat64Ptr()))
					}
				}
			default:

				oneOf := func(tt reflect.Kind) bool {
					switch tt {
					case reflect.Int64, reflect.String, reflect.Bool, reflect.Float64:
						return true
					}
					return false
				}

				// Detect if the object is not nil
				if valTfField.IsValid() && valTfField.IsZero() {
					break
				}

				// Default is not simple type. Require to get type in the goStruct to initialize the generic type in the tfStruct
				// Define if is a slice, map or a struct in the goStruct
				switch valGoField.Kind() {
				case reflect.Map:

					// rTfValue is a terraform struct decoded
					rTfValue := getOf(ctx, valTfField)

					if rTfValue.Kind() != reflect.Map {
						panic(fmt.Sprintf("field %s with type %s is not a map", fieldTf.Name, rTfValue.Type()))
					}

					// * Here catch the case of map[string]string, map[string]int64, map[string]float64
					// rTfValue == map[string]string -- rTfValue.Type().Elem().Kind() == string
					if rTfValue.Kind() == reflect.Map && oneOf(rTfValue.Type().Elem().Kind()) {
						valGoField.Set(rTfValue)
						break
					}

					// * Here catch the case of map[string]*structName
					// Init the map in the goStruct
					valGoField.Set(reflect.MakeMap(valGoField.Type()))

					// Loop over the map in the tfStruct
					for _, k := range rTfValue.MapKeys() {
						// Eq item := &structName{}
						item := reflect.New(valGoField.Type().Elem().Elem())
						// Set the pointer in the map
						valGoField.SetMapIndex(k, item)
						// Decode the tfStruct into goStruct
						if err := decodeToStruct(ctx, rTfValue.Type().Elem(), rTfValue.MapIndex(k), valGoField.Type().Elem(), valGoField.MapIndex(k)); err != nil {
							return err
						}
					}

				case reflect.Slice:
					// rTfValue is a terraform struct decoded
					rTfValue := getOf(ctx, valTfField)

					if rTfValue.Kind() != reflect.Slice {
						panic(fmt.Sprintf("field %s with type %s is not a slice", fieldTf.Name, rTfValue.Type()))
					}

					// * Here catch the case of []string, []int64, []bool, []float64
					// rTfValue == []string -- rTfValue.Type().Elem().Kind() == string
					if rTfValue.Kind() == reflect.Slice && oneOf(rTfValue.Type().Elem().Kind()) {
						valGoField.Set(rTfValue)
						break
					}

					// * Here catch the case of []*structName
					// Init the slice in the goStruct
					valGoField.Set(reflect.MakeSlice(valGoField.Type(), rTfValue.Len(), 2048))

					// Loop over the slice in the tfStruct
					for j := 0; j < rTfValue.Len(); j++ {
						// Eq item := &structName{}
						item := reflect.New(valGoField.Type().Elem().Elem())
						// Set the pointer in the slice
						valGoField.Index(j).Set(item)

						// Decode the tfStruct into goStruct
						if err := decodeToStruct(ctx, rTfValue.Type().Elem(), rTfValue.Index(j), valGoField.Type().Elem(), valGoField.Index(j)); err != nil {
							return err
						}
					}

				case reflect.Ptr:
					// rTfValue is a terraform struct decoded
					rTfValue := getOf(ctx, valTfField)

					item := reflect.New(valGoField.Type().Elem())
					valGoField.Set(item)

					if err := decodeToStruct(ctx, rTfValue.Type(), rTfValue, valGoField.Type(), valGoField); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func reflectStruct(s any) (reflect.Value, reflect.Type) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val, val.Type()
}

func getOf(ctx context.Context, tfV reflect.Value) reflect.Value {
	x := tfV.MethodByName("Get").Call([]reflect.Value{reflect.ValueOf(ctx)})

	values := x[0]
	diags, ok := x[1].Interface().(diag.Diagnostics)
	if !ok {
		panic(fmt.Sprintf("%T is not a diag.Diagnostics", x[1]))
	}
	if diags.HasError() {
		panic(fmt.Sprintf("%s", diags))
	}

	return values
}

// func isKnown(tfV reflect.Value) bool {
// 	x := tfV.MethodByName("IsKnown").Call([]reflect.Value{})
// 	return x[0].Bool()
// }

// func setNull(ctx context.Context, tfV reflect.Value) {
// 	tfV.Addr().MethodByName("SetNull").Call([]reflect.Value{reflect.ValueOf(ctx)})
// }

func set(tfV, goV reflect.Value) {
	if goV.Kind() == reflect.Ptr {
		tfV.Addr().MethodByName("SetPtr").Call([]reflect.Value{goV})
	} else {
		tfV.Addr().MethodByName("Set").Call([]reflect.Value{goV})
	}
}

func setOf(ctx context.Context, tfV, goV reflect.Value) {
	x := tfV.Addr().MethodByName("Set").Call([]reflect.Value{reflect.ValueOf(ctx), goV})
	diags, ok := x[0].Interface().(diag.Diagnostics)
	if !ok {
		panic(fmt.Sprintf("%T is not a diag.Diagnostics", x[1]))
	}
	if diags.HasError() {
		panic(fmt.Sprintf("%s", diags))
	}
}
