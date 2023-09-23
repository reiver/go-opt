package opt

import (
	"fmt"
)

// Optional represents an optional value.
//
// In other programming languages this is know as: an option type, or a maybe type.
//
// For example:
//
//	var op opt.Optional[string] = opt.Something("once twice thrice fource")
//	
//	// ...
//	
//	value, found := op.Get()
//	if found {
//		fmt.Println("value:", value)
//	} else{
//		fmt.Println("nothing")
//	}
//
// Also, for example:
//
//	var op opt.Optional[uint8] = opt.Something[uint8](101)
//	
//	// ...
//	
//	value, found := op.Get()
//	if found {
//		fmt.Println("value:", value)
//	} else{
//		fmt.Println("nothing")
//	}
type Optional[T any] struct {
	value T
	something bool
}

// Nothing returns an optional-type with nothing in it.
//
// For example, here is an optional-type that can hold a string with nothing in it:
//
//	var op opt.Optional[string] = opt.Nothing[string]()
//
// Note that the default value for any optional-type is nothing.
// So the following code it equivalent to it:
//
//	var op opt.Optional[string]
//
// Also, for example, here is an optional-type that can hold a uint8 with nothing in it:
//
//	var op opt.Optional[uint8] = opt.Nothing[uint8]()
//
// Again, note that the default value for any optional-type is nothing.
// So the following code it equivalent to it:
//
//	var op opt.Optional[uint8]
func Nothing[T any]() Optional[T] {
	var nothing Optional[T]

	return nothing
}

// Something returns a optional-type with something in it.
//
// For example, here is an optional-type with the string "once twice thrice fource" in it:
//
//	var op opt.Optional[string] = opt.Something("once twice thrice fource")
//
// And, for example, here is an optional-type with the uint8 101 in it:
//
//
//	var op opt.Optional[uint8] = opt.Something(uint8(101))
func Something[T any](value T) Optional[T] {
	return Optional[T]{
		something:true,
		value:value,
	}
}

// Filter returns itself if it is holding something and ‘fn’ applied to its value returns true.
// Else it return nothing.
//
// For example:
//
//	fn := func(value int) bool {
//		return 0 == (value % 2)
//	}
//	
//	// ...
//	
//	var op1 opt.Optional[int] = opt.Something[int](10)
//	
//	op1 = op1.Filter(fn)
//	
//	// ...
//	
//	var op2 opt.Optional[int] = opt.Something[int](11)
//	
//	op2 = op2.Filter(fn)
//	
//	// ...
//	
//	var op3 opt.Optional[int] = opt.Nothing[int]()
//	
//	op3 = op3.Filter(fn)
func (receiver Optional[T]) Filter(fn func(T)bool) Optional[T] {
	if receiver.isnothing() {
		return Nothing[T]()
	}

	if !fn(receiver.value) {
		return Nothing[T]()
	}

	return receiver
}

func (receiver Optional[T]) isnothing() bool {
	return !receiver.something
}

// Get returns the value inside of the optional-type if it is holding something.
//
// Example usage:
//
//	var op opt.Optional[string]
//	
//	// ...
//	
//	value, found := op.Get()
//	
//	if found {
//		fmt.Println("VALUE:", value)
//	} else {
//		fmt.Println("nothing")
//	}
func (receiver Optional[T]) Get() (T, bool) {
	return receiver.value, receiver.something
}

// GoString makes it that when the fmt.Fprintf(), fmt.Printf(), and fmt.Sprintf() family of functions
// renders this type with the %#v verb, that it will be easier to understand.
//
// For example:
//
//	var op opt.Optional[string] = opt.Something("once twice thrice fource")
//	
//	// ...
//	
//	fmt.Printf("op = %#v", op)
//
//	// Output:
//	// op = opt.Something[string]("once twice thrice fource")
//
// Also, for example:
//
//	var op opt.Optional[uint8] = opt.Nothing[uint8]()
//	
//	// ...
//	
//	fmt.Printf("op = %#v", op)
//
//	// Output:
//	// op = opt.Nothing[uint8]()
func (receiver Optional[T]) GoString() string {
	if receiver.isnothing() {
		return fmt.Sprintf("opt.Nothing[%T]()", receiver.value)
	}

	return fmt.Sprintf("opt.Something[%T](%#v)", receiver.value, receiver.value)
}

// WhenNothing will call ‘fn’ when ‘receiver’ is holding nothing.
//
// It will not call ‘fn’ when ‘receier’ is hold something.
//
// For example:
//
//	var op opt.Optional = opt.Nothing[string]
//	
//	// ...
//
//	op.WhenNothing(func(){
//		//@TODO
//	})
func (receiver Optional[T]) WhenNothing(fn func()) {
	if receiver.isnothing() {
		fn()
	}
}

// WhenSomething will ‘fn’ when ‘receiver’ is holding something.
// The value that ‘receiver’  is holding will be passed as a parameter to the function ‘fn’.
//
// It will not call ‘fn’ when ‘receiver’ is hold nothing.
//
// For example:
//
//	var op opt.Optional = opt.Something("once twice thrice fource")
//	
//	// ...
//
//	op.WhenSomething(func(value string){
//		//@TODO
//	})
func (receiver Optional[T]) WhenSomething(fn func(T)) {
	if receiver.something {
		fn(receiver.value)
	}
}
