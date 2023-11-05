package opt_test

import (
	"sourcecode.social/reiver/go-opt"

	"testing"
)

func TestOptional_GetElse_string(t *testing.T) {

	tests := []struct{
		Optional opt.Optional[string]
		Alternative string
		Expected string
	}{
		{
			Optional: opt.Nothing[string](),
			Alternative: "alternative",
			Expected:    "alternative",
		},
		{
			Optional: opt.Nothing[string](),
			Alternative: "",
			Expected:    "",
		},
		{
			Optional: opt.Nothing[string](),
			Alternative: "apple",
			Expected:    "apple",
		},
		{
			Optional: opt.Nothing[string](),
			Alternative: "banana",
			Expected:    "banana",
		},
		{
			Optional: opt.Nothing[string](),
			Alternative: "cherry",
			Expected:    "cherry",
		},
		{
			Optional: opt.Nothing[string](),
			Alternative: "ONCE TWICE THRICE FOURCE",
			Expected:    "ONCE TWICE THRICE FOURCE",
		},
		{
			Optional: opt.Nothing[string](),
			Alternative: "😈",
			Expected:    "😈",
		},



		{
			Optional: opt.Something("ONCE TWICE THRICE FOURCE"),
			Alternative: "😈",
			Expected:                "ONCE TWICE THRICE FOURCE",
		},



		{
			Optional: opt.Something("۰۱۲۳۴۵۶۷۸۹"),
			Alternative: "😈",
			Expected:               "۰۱۲۳۴۵۶۷۸۹",
		},
	}

	for testNumber, test := range tests {

		actual := test.Optional.GetElse(test.Alternative)

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("OPTIONAL: %#v", test.Optional)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_GetElse_int8(t *testing.T) {

	tests := []struct{
		Optional opt.Optional[int8]
		Alternative int8
		Expected int8
	}{
		{
			Optional: opt.Nothing[int8](),
			Alternative: -2,
			Expected:    -2,
		},
		{
			Optional: opt.Nothing[int8](),
			Alternative: -1,
			Expected:    -1,
		},
		{
			Optional: opt.Nothing[int8](),
			Alternative: 0,
			Expected:    0,
		},
		{
			Optional: opt.Nothing[int8](),
			Alternative: 1,
			Expected:    1,
		},
		{
			Optional: opt.Nothing[int8](),
			Alternative: 2,
			Expected:    2,
		},



		{
			Optional: opt.Something[int8](-127),
			Alternative: 42,
			Expected:                     -127,
		},
		{
			Optional: opt.Something[int8](0),
			Alternative: 42,
			Expected:                     0,
		},
		{
			Optional: opt.Something[int8](127),
			Alternative: 42,
			Expected:                     127,
		},
	}

	for testNumber, test := range tests {

		actual := test.Optional.GetElse(test.Alternative)

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("OPTIONAL: %#v", test.Optional)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_GetElse_uint8(t *testing.T) {

	tests := []struct{
		Optional opt.Optional[uint8]
		Alternative uint8
		Expected uint8
	}{
		{
			Optional: opt.Nothing[uint8](),
			Alternative: 0,
			Expected:    0,
		},
		{
			Optional: opt.Nothing[uint8](),
			Alternative: 1,
			Expected:    1,
		},
		{
			Optional: opt.Nothing[uint8](),
			Alternative: 2,
			Expected:    2,
		},



		{
			Optional: opt.Something[uint8](0),
			Alternative: 42,
			Expected:                     0,
		},
		{
			Optional: opt.Something[uint8](127),
			Alternative: 42,
			Expected:                     127,
		},
	}

	for testNumber, test := range tests {

		actual := test.Optional.GetElse(test.Alternative)

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED VALUE: %q", expected)
				t.Logf("ACTUAL   VALUE: %q", actual)
				t.Logf("OPTIONAL: %#v", test.Optional)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}
