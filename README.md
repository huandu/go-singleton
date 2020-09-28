# go-singleton: Experimental singleton implementation in Go2

Per the article [The Next Step for Generics](https://blog.golang.org/generics-next-step) published by Go Authors, Go2 will support generic. This package is an experimental Go2 package to implement a type-safe singleton - a typical scenario in generic.

## Install

Since now (Jul 14, 2020), `go2go` doesn't support `go mod`. The only way (it seems) to import this package is to clone it manually to `$GO2PATH/src/github.com/huandu/go-generic` and then use it. Mind that the name of Go path environment variable is changed to `GO2PATH`.

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
