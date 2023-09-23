package opt

// Then applies the function ‘fn’ to the value inside of the optional-type ‘op’, if the optional-type ‘op’ is holding something, and returns the resulting  optional-type.
// If the optional-type ‘op’ is holding nothing, then Then also returns nothing.
//
// For example:
//
//	fn := func(s string) opt.Optional[byte] {
//		
//		if len(s) < 2 {
//			return opt.Nothing[byte]()
//		}
//		
//		return opt.Something[byte](s[1])
//	}
//	
//	var op opt.Optional[string] = opt.Something("Hello world!"")
//	
//	var result opt.Optional[byte] = opt.Then(op, fn)
//	
//	// result == opt.Something[byte]('e')
//	
//	// ...
//	
//	var op2 opt.Optional[string] = opt.Something[string]("X")
//	
//	var result2 opt.Optional[byte] = opt.Then(op, fn)
//	
//	// result2 == opt.Nothing[byte]()
//	
//	// ...
//	
//	var op2 opt.Optional[string] = opt.Nothing[string]()
//	
//	var result2 opt.Optional[byte] = opt.Then(op, fn)
//	
//	// result2 == opt.Nothing[byte]()
func Then[T1 any, T2 any](op Optional[T1], fn func(T1)Optional[T2]) Optional[T2] {
	if op.isnothing() {
		return Nothing[T2]()
	}

	return fn(op.value)
}
