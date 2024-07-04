package opt_test

import (
	"github.com/reiver/go-opt"

	"testing"
)

func TestMap_stringToUint(t *testing.T) {

	tests := []struct{
		Optional opt.Optional[string]
		Expected opt.Optional[int]
	}{
		{
			Optional: opt.Nothing[string](),
			Expected: opt.Nothing[int](),
		},



		{
			Optional: opt.Something(""),
			Expected: opt.Something[int](0),
		},



		{
			Optional: opt.Something("once"),
			Expected: opt.Something[int](4),
		},
		{
			Optional: opt.Something("twice"),
			Expected: opt.Something[int](5),
		},
		{
			Optional: opt.Something("thrice"),
			Expected: opt.Something[int](6),
		},
		{
			Optional: opt.Something("fource"),
			Expected: opt.Something[int](6),
		},
	}

	for testNumber, test := range tests {

		fn := func(s string) int {
			return len(s)
		}

		mapped := opt.Map(test.Optional, fn)

		{
			expected := test.Expected
			actual   := mapped

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
