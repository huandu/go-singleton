# go-singleton: Singleton using generic technic

Starting from go1.18, Go will officially support generic. This package is designed to embrace the new technic for convenience.

Go1.18 (or later) is required to compile this package.

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
