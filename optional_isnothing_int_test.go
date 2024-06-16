package opt_test

import (
	"testing"

	"sourcecode.social/reiver/go-opt"
)

func TestOptional_IsNothing_int(t *testing.T) {
	tests := []struct{
		Optional opt.Optional[int]
		Expected bool
	}{
		{
			Optional: opt.Nothing[int](),
			Expected: true,
		},



		{
			Optional: opt.Something(-255),
			Expected: false,
		},

		{
			Optional: opt.Something(-5),
			Expected: false,
		},
		{
			Optional: opt.Something(-4),
			Expected: false,
		},
		{
			Optional: opt.Something(-3),
			Expected: false,
		},
		{
			Optional: opt.Something(-2),
			Expected: false,
		},
		{
			Optional: opt.Something(-1),
			Expected: false,
		},
		{
			Optional: opt.Something(0),
			Expected: false,
		},
		{
			Optional: opt.Something(1),
			Expected: false,
		},
		{
			Optional: opt.Something(2),
			Expected: false,
		},
		{
			Optional: opt.Something(3),
			Expected: false,
		},
		{
			Optional: opt.Something(4),
			Expected: false,
		},
		{
			Optional: opt.Something(5),
			Expected: false,
		},

		{
			Optional: opt.Something(255),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Optional.IsNothing()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value of is-nothing it not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("OPTIONAL: %#v", test.Optional)
			continue
		}
	}
}
