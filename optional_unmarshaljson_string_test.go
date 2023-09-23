package opt_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-opt"
)

func TestOptional_UnmarshalJSON_string(t *testing.T) {

	tests := []struct{
		JSON string
		Expected opt.Optional[string]
	}{
		{
			JSON: `""`,
			Expected: opt.Something(""),
		},



		{
			JSON: `"apple"`,
			Expected: opt.Something("apple"),
		},
		{
			JSON: `"banana"`,
			Expected: opt.Something("banana"),
		},
		{
			JSON: `"cherry"`,
			Expected: opt.Something("cherry"),
		},



		{
			JSON: `"ONCE"`,
			Expected: opt.Something("ONCE"),
		},
		{
			JSON: `"TWICE"`,
			Expected: opt.Something("TWICE"),
		},
		{
			JSON: `"THRICE"`,
			Expected: opt.Something("THRICE"),
		},
		{
			JSON: `"FOURCE"`,
			Expected: opt.Something("FOURCE"),
		},



		{
			JSON: `"🙂"`,
			Expected: opt.Something("🙂"),
		},
		{
			JSON: `"😈"`,
			Expected: opt.Something("😈"),
		},
		{
			JSON: `"❤️"`,
			Expected: opt.Something("❤️"),
		},



		{
			JSON: `"٠١٢٣۴۵۶٧٨٩"`,
			Expected: opt.Something("٠١٢٣۴۵۶٧٨٩"),
		},



		{
			JSON: `"𐏑𐏓𐏕"`,
			Expected: opt.Something("𐏑𐏓𐏕"),
		},
	}

	for testNumber, test := range tests {

		var actual opt.Optional[string]

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

func TestOptional_UnmarshalJSON_string_fail(t *testing.T) {

	var op opt.Optional[string]

	var jason string = "12345"

	err := json.Unmarshal([]byte(jason), &op)
	if nil == err {
		t.Errorf("Expected an error but did not actually get one.")
		return
	}
}
