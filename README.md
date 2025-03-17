<div align="center">
    <a href="https://github.com/orange-cloudavenue/terraform-plugin-framework-supertypes/releases/latest">
      <img alt="Latest release" src="https://img.shields.io/github/v/release/orange-cloudavenue/terraform-plugin-framework-supertypes?style=for-the-badge&logo=starship&color=C9CBFF&logoColor=D9E0EE&labelColor=302D41&include_prerelease&sort=semver" />
    </a>
    <a href="https://github.com/orange-cloudavenue/terraform-plugin-framework-supertypes/pulse">
      <img alt="Last commit" src="https://img.shields.io/github/last-commit/orange-cloudavenue/terraform-plugin-framework-supertypes?style=for-the-badge&logo=starship&color=8bd5ca&logoColor=D9E0EE&labelColor=302D41"/>
    </a>
    <a href="https://github.com/orange-cloudavenue/terraform-plugin-framework-supertypes/stargazers">
      <img alt="Stars" src="https://img.shields.io/github/stars/orange-cloudavenue/terraform-plugin-framework-supertypes?style=for-the-badge&logo=starship&color=c69ff5&logoColor=D9E0EE&labelColor=302D41" />
    </a>
    <a href="https://github.com/orange-cloudavenue/terraform-plugin-framework-supertypes/issues">
      <img alt="Issues" src="https://img.shields.io/github/issues/orange-cloudavenue/terraform-plugin-framework-supertypes?style=for-the-badge&logo=bilibili&color=F5E0DC&logoColor=D9E0EE&labelColor=302D41" />
    </a>
</div>

# terraform-plugin-framework-supertypes

This repository contains the custom types for the Terraform Plugin Framework.

*Why supertypes?*

The supertypes are custom types that are used to extend the functionality of the Terraform Plugin Framework. They are used to define the custom types that are used in the Terraform Plugin Framework.

Custom types available in the supertypes:

* **Simple types**
  * `StringType` - A custom type that is used to define the string type.
  * `Int32Type` - A custom type that is used to define the integer type.
  * `Int64Type` - A custom type that is used to define the integer type.
  * `NumberType` - A custom type that is used to define the number type.
  * `BoolType` - A custom type that is used to define the boolean type.
  * `Float32Type` - A custom type that is used to define the float type.
  * `Float64Type` - A custom type that is used to define the float type.
* **Array types**
  * `MapType` - A custom type that is used to define the map type.
  * `ListType` - A custom type that is used to define the list type.
  * `SetType` - A custom type that is used to define the set type.
  * `MapTypeOf` - A custom type that is used to define the map type. This type use golang generics to construct automatically the map.
  * `ListTypeOf` - A custom type that is used to define the list type. This type use golang generics to construct automatically the list.
  * `SetTypeOf` - A custom type that is used to define the set type. This type use golang generics to construct automatically the set.
* **Object types**
  * `SingleNestedObjectTypeOf` - A custom type that is used to define the nested object value of the single type. This type use golang generics to construct automatically the nested object.
* **Object array types**
  * `MapNestedType` - A custom type that is used to define the nested map type.
  * `ListNestedType` - A custom type that is used to define the nested list type.
  * `SetNestedType` - A custom type that is used to define the nested set type.
  * `MapNestedObjectTypeOf` - A custom type that is used to define the nested object value of the map type. This type use golang generics to construct automatically the nested object.
  * `ListNestedObjectTypeOf` - A custom type that is used to define the nested object value of the list type. This type use golang generics to construct automatically the nested object.
  * `SetNestedObjectTypeOf` - A custom type that is used to define the nested object value of the set type. This type use golang generics to construct automatically the nested object.

Principal methods available for each custom type:

* `Get()` - A function that is used to get the value.
* `GetPtr()` - A function that is used to get the value as a pointer.
* `Set()` - A function that is used to set the value.
* `SetPtr()` - A function that is used to set the value as a pointer.
* `SetNull()` - A function that is used to set the value to null.
* `SetUnknown()` - A function that is used to set the value to unknown.
* `IsKnown()` - A function that is used to check if the value is known.

Special methods available for the `Int32Type` and `Int64Type`:

* `SetInt()` - A function that is used to set `int` into the type (`int32` or `int64`).
* `SetInt8()` - A function that is used to set `int8` into the type (`int32` or `int64`).
* `SetInt16()` - A function that is used to set `int16` into the type (`int32` or `int64`).
* `SetInt32()` - A function that is used to set `int32` into the type (`int32` or `int64`).
* `SetInt64()` - A function that is used to set `int64` into the type (`int32` or `int64`).
* `SetIntPtr()` - A function that is used to set `int` pointer into the type (`int32` or `int64`).
* `SetInt8Ptr()` - A function that is used to set `int8` pointer into the type (`int32` or `int64`).
* `SetInt16Ptr()` - A function that is used to set `int16` pointer into the type (`int32` or `int64`).
* `SetInt32Ptr()` - A function that is used to set `int32` pointer into the type (`int32` or `int64`).
* `SetInt64Ptr()` - A function that is used to set `int64` pointer into the type (`int32` or `int64`).
* `GetInt()` - A function that is used to get the value as an `int`.
* `GetInt8()` - A function that is used to get the value as an `int8`.
* `GetInt16()` - A function that is used to get the value as an `int16`.
* `GetInt32()` - A function that is used to get the value as an `int32`.
* `GetInt64()` - A function that is used to get the value as an `int64`.
* `GetIntPtr()` - A function that is used to get the value as an `int` pointer.
* `GetInt8Ptr()` - A function that is used to get the value as an `int8` pointer.
* `GetInt16Ptr()` - A function that is used to get the value as an `int16` pointer.
* `GetInt32Ptr()` - A function that is used to get the value as an `int32` pointer.
* `GetInt64Ptr()` - A function that is used to get the value as an `int64` pointer.

Special methods available for the all `*TypeOf` types:

* `MustGet(context.Context)` - A function that is used to get the value. If the value get return an error, the function panic.
* `DiagsGet(context.Context, diag.Diagnostics)` - A function that is used to get the value. If the value get return an error, the function add the error to the diagnostics.
* `MustSet(context.Context, T)` - A function that is used to set the value. If the value set return an error, the function panic.
* `DiagsSet(context.Context, diag.Diagnostics, T)` - A function that is used to set the value. If the value set return an error, the function add the error to the diagnostics.

## Installation

For installing the supertypes, you can use the `go get` command:

```sh
go get github.com/orange-cloudavenue/terraform-plugin-framework-supertypes@latest
```

## Documentation

documentation is in progress.
