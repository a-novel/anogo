# AnoGo

AnoGo is an util package for Go, used by most A-Novel repositories.

```cgo
go get github.com/a-novel/anogo
```

# Pointer functions

The following functions refer to pointers and their usage.

## IsPtr

This function will return an error if the value passed as an argument doesn't represent a direct pointer to a defined
value. Nested pointers are also refused.

```go
err := anogo.IsPtr(value)
```

# License
2021, A-Novel [MIT License](https://github.com/a-novel/anogo/blob/master/LICENSE).