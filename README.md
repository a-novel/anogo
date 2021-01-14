# AnoGo

AnoGo is an util package for Go, used by most A-Novel repositories.

```cgo
go get github.com/a-novel/anogo
```

- [Map functions](#map-functions)
    - [ToMap functions](#tomap-functions)
- [Pointer functions](#pointer-functions)
    - [IsPtr](#isptr)
    - [IsSlicePtr](#issliceptr)
    
# Map functions

The following functions refer to map-like go types and their usage.

## ToMap functions

Perform conversions between map-like values.

| Function           | Arguments                         |
| :---               | :---                              |
| `ToMap()`          | v `interface{}`<br/>ptr `pointer` |
| `ToMapInterface()` | v `interface{}`                   |
| `ToMapString()`    | v `interface{}`                   |

Each function may return any of the following errors:

- `anogo.ErrIsNotPtr`
- `anogo.ErrIsNestedPtr`
- `anogo.ErrIsNotMappable`
- `anogo.ErrCannotMarshal`
- `anogo.ErrCannotUnmarshal`

## Flatten

Flatten a map-like object.

```go
flattened, err := anogo.Flatten(data)
```

Data should be either of type struct or map.

It may return any of the following errors:

- `anogo.ErrCannotMarshal`
- `anogo.ErrCannotUnmarshal`

# Pointer functions

The following functions refer to pointers and their usage.

## IsPtr

This function will return an error if the value passed as an argument doesn't represent a direct pointer to a defined
value. Nested pointers are also refused.

```go
err := anogo.IsPtr(value)
```

It may return any of the following errors:

- `anogo.ErrIsNotPtr`
- `anogo.ErrIsNestedPtr`

## IsSlicePtr

This function will return an error if the value passed as an argument doesn't represent a slice of pointers. Nested
pointers are also refused for slice type.

```go
err := anogo.IsPtr(value)
```

It may return any of the following errors:

- `anogo.ErrIsNotPtr`
- `anogo.ErrIsNestedPtr`
- `anogo.ErrIsNotSlice`

# License
2021, A-Novel [MIT License](https://github.com/a-novel/anogo/blob/master/LICENSE).