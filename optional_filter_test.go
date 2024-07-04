package opt_test

import (
	"github.com/reiver/go-opt"

	"testing"
)

func TestOptional_Filter_int(t *testing.T) {

	tests := []struct{
		Optional opt.Optional[int]
		Expected opt.Optional[int]
	}{
		{
			Optional: opt.Nothing[int](),
			Expected: opt.Nothing[int](),
		},



		{
			Optional: opt.Something[int](-2),
			Expected: opt.Something[int](-2),
		},
		{
			Optional: opt.Something[int](-1),
			Expected: opt.Nothing[int](),
		},
		{
			Optional: opt.Something[int](0),
			Expected: opt.Something[int](0),
		},
		{
			Optional: opt.Something[int](1),
			Expected: opt.Nothing[int](),
		},
		{
			Optional: opt.Something[int](2),
			Expected: opt.Something[int](2),
		},
	}

	for testNumber, test := range tests {

		fn := func(value int) bool {
			return 0 == (value % 2)
		}

		actual   := test.Optional.Filter(fn)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
	/////////////// CONTINUE
			continue
		}
	}
}
