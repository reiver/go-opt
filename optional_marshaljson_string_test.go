package opt_test

import (
	"testing"

	"sourcecode.social/reiver/go-opt"
)

func TestOptional_MarshalJSON_string(t *testing.T) {

	tests := []struct{
		Value opt.Optional[string]
		Expected string
	}{
		{
			Value: opt.Something(""),
			Expected: `""`,
		},



		{
			Value: opt.Something("apple"),
			Expected: `"apple"`,
		},
		{
			Value: opt.Something("banana"),
			Expected: `"banana"`,
		},
		{
			Value: opt.Something("cherry"),
			Expected: `"cherry"`,
		},



		{
			Value: opt.Something("ONCE"),
			Expected: `"ONCE"`,
		},
		{
			Value: opt.Something("TWICE"),
			Expected: `"TWICE"`,
		},
		{
			Value: opt.Something("THRICE"),
			Expected: `"THRICE"`,
		},
		{
			Value: opt.Something("FOURCE"),
			Expected: `"FOURCE"`,
		},



		{
			Value: opt.Something("🙂"),
			Expected: `"🙂"`,
		},
		{
			Value: opt.Something("😈"),
			Expected: `"😈"`,
		},
		{
			Value: opt.Something("❤️"),
			Expected: `"❤️"`,
		},



		{
			Value: opt.Something("٠١٢٣۴۵۶٧٨٩"),
			Expected: `"٠١٢٣۴۵۶٧٨٩"`,
		},



		{
			Value: opt.Something("𐏑𐏓𐏕"),
			Expected: `"𐏑𐏓𐏕"`,
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
