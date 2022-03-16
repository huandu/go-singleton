# go-singleton: Generic singleton factory

[![Build Status](https://github.com/huandu/go-singleton/workflows/Go/badge.svg)](https://github.com/huandu/go-singleton/actions)
[![GoDoc](https://godoc.org/github.com/huandu/go-singleton?status.svg)](https://pkg.go.dev/github.com/huandu/go-singleton)

Starting from go1.18, Go will officially support generic. This package is designed to embrace the new technic for convenience.

## Usage

Call `Singleton` with a given type name to get a pointer to this type.

```go
type MyType struct {
    field int
}

v1 := singleton.Singleton[MyType]()
v1.field = 123

v2 := singleton.Singleton[MyType]()
println(v2.field) // Output: 123
```
