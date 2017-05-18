/*
Package opt has a number of sub-packages provides alternatives to Go's built-in types (bool, float64, int64, string, time.Time)
that also let you express a "lack of a value"; for the Go programming language; these are similar to "option types" or "maybe types"
in some other programming languages.

Example:

	var s optstring.NullableType
	
	// ...
	
	switch s {
	case optstring.NoneNullable():
		//@TODO
	case optstring.Null():
		//@TODO
	case optstring.SomeNullable("Hello world!"):
		//@TODO
	case optstring.SomeNullable("apple banana cherry"):
		//@TODO
	default:
		//@TODO
	}

Example:

	var s optstring.Type
	
	// ...
	
	switch s {
	case optstring.None():
		//@TODO
	case optstring.Some("Hello world!"):
		//@TODO
	case optstring.Some("apple banana cherry"):
		//@TODO
	default:
		//@TODO
	}
*/
package opt

