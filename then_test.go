package opt_test

import (
	"sourcecode.social/reiver/go-opt"

	"testing"
)

func TestThen_stringToUint(t *testing.T) {

	tests := []struct{
		Optional opt.Optional[string]
		Expected opt.Optional[byte]
	}{
		{
			Optional: opt.Nothing[string](),
			Expected: opt.Nothing[byte](),
		},



		{
			Optional: opt.Something(""),
			Expected: opt.Nothing[byte](),
		},



		{
			Optional: opt.Something("A"),
			Expected: opt.Nothing[byte](),
		},
		{
			Optional: opt.Something("B"),
			Expected: opt.Nothing[byte](),
		},
		{
			Optional: opt.Something("C"),
			Expected: opt.Nothing[byte](),
		},



		{
			Optional: opt.Something("once"),
			Expected: opt.Something[byte]('n'),
		},
		{
			Optional: opt.Something("twice"),
			Expected: opt.Something[byte]('w'),
		},
		{
			Optional: opt.Something("thrice"),
			Expected: opt.Something[byte]('h'),
		},
		{
			Optional: opt.Something("fource"),
			Expected: opt.Something[byte]('o'),
		},
	}

	for testNumber, test := range tests {

		fn := func(s string) opt.Optional[byte] {
			if len(s) < 2 {
				return opt.Nothing[byte]()
			}

			return opt.Something[byte](s[1])
		}

		thened := opt.Then(test.Optional, fn)

		{
			expected := test.Expected
			actual   := thened

			if expected != actual  {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: (%T) %#v", expected, expected)
				t.Logf("ACTUAL:   (%T) %#v", actual, actual)
				t.Logf("OPTIONAL: (%T), %#v", test.Optional, test.Optional)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}
