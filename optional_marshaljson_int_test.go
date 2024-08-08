package opt_test

import (
	"testing"

	"github.com/reiver/go-opt"
)

func TestOptional_MarshalJSON_int(t *testing.T) {

	tests := []struct{
		Value opt.Optional[int]
		Expected string
	}{
		{
			Value: opt.Something(-65536),
			Expected:           `-65536`,
		},
		{
			Value: opt.Something(-65535),
			Expected:           `-65535`,
		},



		{
			Value: opt.Something(-256),
			Expected:           `-256`,
		},
		{
			Value: opt.Something(-255),
			Expected:           `-255`,
		},



		{
			Value: opt.Something(-5),
			Expected:           `-5`,
		},
		{
			Value: opt.Something(-4),
			Expected:           `-4`,
		},
		{
			Value: opt.Something(-3),
			Expected:           `-3`,
		},
		{
			Value: opt.Something(-2),
			Expected:           `-2`,
		},
		{
			Value: opt.Something(-1),
			Expected:           `-1`,
		},
		{
			Value: opt.Something(0),
			Expected:           `0`,
		},
		{
			Value: opt.Something(1),
			Expected:           `1`,
		},
		{
			Value: opt.Something(2),
			Expected:           `2`,
		},
		{
			Value: opt.Something(3),
			Expected:           `3`,
		},
		{
			Value: opt.Something(4),
			Expected:           `4`,
		},
		{
			Value: opt.Something(5),
			Expected:           `5`,
		},



		{
			Value: opt.Something(255),
			Expected:           `255`,
		},
		{
			Value: opt.Something(256),
			Expected:           `256`,
		},



		{
			Value: opt.Something(65535),
			Expected:           `65535`,
		},
		{
			Value: opt.Something(65536),
			Expected:           `65536`,
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

func TestOptional_MarshalJSON_int_fail(t *testing.T) {

	var nothing opt.Optional[int]

	actualBytes, err := nothing.MarshalJSON()
	if nil == err {
		t.Error("Expected an error but did not actually get one.")
		t.Logf("ACTUAL: %q", actualBytes)
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
	if nil != actualBytes {
		t.Error("Expected not bytes but actually get some.")
		t.Logf("ACTUAL: %q", actualBytes)
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}
}
