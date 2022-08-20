# go-opt

Package **opt** implements an **optional-type**, for the Go programming language.

In other programming languages, an **optional-type** might be know as: a **option type**, or a **maybe type**.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-opt

[![GoDoc](https://godoc.org/github.com/reiver/go-opt?status.svg)](https://godoc.org/github.com/reiver/go-opt)

## Examples

Here is an example **optional-type** that can hold a string:
```go
var op opt.Optional[string] // the default value is nothing.

// ...

if opt.Nothing[string]() == op {
	fmt.Println("nothing")
}

// ...

op = opt.Something("Hello world! 👾")

// ...

switch op {
case op.Nothing[string]():
	//@TODO
case op.Something("apple"):
	//@TODO
case op.Something("banana"):
	//@TODO
case op.Something("cherry"):
	//@TODO
default:
	//@TODO
}

// ...

value, found := op.Get()
if found {
	fmt.Println("VALUE:", value)
} else {
	fmt.Println("nothing")
}
```
