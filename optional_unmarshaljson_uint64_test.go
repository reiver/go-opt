package opt_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-opt"
)

func TestOptional_UnmarshalJSON_uint64(t *testing.T) {

	tests := []struct{
		JSON string
		Expected opt.Optional[uint64]
	}{
		{
			JSON: `0`,
			Expected: opt.Something(uint64(0)),
		},
		{
			JSON: `1`,
			Expected: opt.Something(uint64(1)),
		},


		{
			JSON: `255`,
			Expected: opt.Something(uint64(255)),
		},
	}

	for testNumber, test := range tests {

		var actual opt.Optional[uint64]

		err := json.Unmarshal([]byte(test.JSON), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON: %q", test.JSON)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value of the optional type is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON: %q", test.JSON)
				continue
			}
		}
	}
}

func TestOptional_UnmarshalJSON_uint64_fail(t *testing.T) {

	var op opt.Optional[uint64]

	var jason string = `false`

	err := json.Unmarshal([]byte(jason), &op)
	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}
}
