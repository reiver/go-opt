package opt

// Map applies the function ‘fn’ to the value inside of the optional-type ‘op’, if the optional-type ‘op’ is holding something, and returns it as a new optional-type.
// If the optional-type ‘op’ is holding nothing, then Map also returns nothing.
//
// For example:
//
//	var op opt.Optional[string] = opt.Something("Hello world!")
//	
//	var result opt.Optional[string] = opt.Map(op, strings.ToUpper)
//	
//	// result == opt.Something[string]("HELLO WORLD!")
//	
//	// ...
//	
//	var op2 opt.Optional[string] = opt.Nothing[string]()
//	
//	var result2 opt.Optional[string] = opt.Map(op, strings.ToUpper)
//	
//	// result2 == opt.Nothing[string]()
//
// Or also, for example:
//
//	fn := func(s string) int {
//		return len(s)
//	}
//	
//	var op opt.Optional[string] = opt.Something("Hello world!")
//	
//	var result opt.Optional[int] = opt.Map(op, fn)
//	
//	// result == opt.Something[int](12)
//	
//	// ...
//	
//	var op2 opt.Optional[string] = opt.Nothing[string]()
//	
//	var result2 opt.Optional[int] = opt.Map(op, fn)
//	
//	// result2 == opt.Nothing[int]()
func Map[T1 any, T2 any](op Optional[T1], fn func(T1)T2) Optional[T2] {
	if op.IsNothing() {
		return Nothing[T2]()
	}

	return Something(fn(op.value))
}
