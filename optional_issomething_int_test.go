package opt_test

import (
	"testing"

	"github.com/reiver/go-opt"
)

func TestOptional_IsSomething_int(t *testing.T) {
	tests := []struct{
		Optional opt.Optional[int]
		Expected bool
	}{
		{
			Optional: opt.Nothing[int](),
			Expected: false,
		},



		{
			Optional: opt.Something(-255),
			Expected: true,
		},

		{
			Optional: opt.Something(-5),
			Expected: true,
		},
		{
			Optional: opt.Something(-4),
			Expected: true,
		},
		{
			Optional: opt.Something(-3),
			Expected: true,
		},
		{
			Optional: opt.Something(-2),
			Expected: true,
		},
		{
			Optional: opt.Something(-1),
			Expected: true,
		},
		{
			Optional: opt.Something(0),
			Expected: true,
		},
		{
			Optional: opt.Something(1),
			Expected: true,
		},
		{
			Optional: opt.Something(2),
			Expected: true,
		},
		{
			Optional: opt.Something(3),
			Expected: true,
		},
		{
			Optional: opt.Something(4),
			Expected: true,
		},
		{
			Optional: opt.Something(5),
			Expected: true,
		},

		{
			Optional: opt.Something(255),
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := test.Optional.IsSomething()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value of is-something it not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("OPTIONAL: %#v", test.Optional)
			continue
		}
	}
}
