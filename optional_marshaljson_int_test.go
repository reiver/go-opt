package opt_test

import (
	"testing"

	"github.com/reiver/go-opt"
)

func TestOptional_MarshalJSON_int(t *testing.T) {

	tests := []struct{
		Value opt.Optional[int]
	}{
		{
			Value: opt.Something(-65536),
		},
		{
			Value: opt.Something(-65535),
		},



		{
			Value: opt.Something(-256),
		},
		{
			Value: opt.Something(-255),
		},



		{
			Value: opt.Something(-5),
		},
		{
			Value: opt.Something(-4),
		},
		{
			Value: opt.Something(-3),
		},
		{
			Value: opt.Something(-2),
		},
		{
			Value: opt.Something(-1),
		},
		{
			Value: opt.Something(0),
		},
		{
			Value: opt.Something(1),
		},
		{
			Value: opt.Something(2),
		},
		{
			Value: opt.Something(3),
		},
		{
			Value: opt.Something(4),
		},
		{
			Value: opt.Something(5),
		},



		{
			Value: opt.Something(255),
		},
		{
			Value: opt.Something(256),
		},



		{
			Value: opt.Something(65535),
		},
		{
			Value: opt.Something(65536),
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := test.Value.MarshalJSON()
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one." , testNumber)
			t.Logf("ACTUAL BYTES: %q", actualBytes)
			t.Logf("ERROR: (%T) %s", err, err)
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
