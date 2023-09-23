package opt_test

import (
	"testing"

	"sourcecode.social/reiver/go-opt"
)

func TestOptional_MarshalJSON_bool(t *testing.T) {

	tests := []struct{
		Value opt.Optional[bool]
		Expected string
	}{
		{
			Value: opt.Something(false),
			Expected: "false",
		},
		{
			Value: opt.Something(true),
			Expected: "true",
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		actual := string(actualBytes)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value for the JSON marshaling is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
