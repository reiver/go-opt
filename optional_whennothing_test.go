package opt_test

import (
	"sourcecode.social/reiver/go-opt"

	"testing"
)

func TestOptional_WhenNothing(t *testing.T) {


	tests := []struct{
		Optional interface{WhenNothing(func())}
	}{
		{
			Optional: opt.Nothing[string](),
		},



		{
			Optional: opt.Nothing[int](),
		},
		{
			Optional: opt.Nothing[int8](),
		},
		{
			Optional: opt.Nothing[int16](),
		},
		{
			Optional: opt.Nothing[int32](),
		},
		{
			Optional: opt.Nothing[int64](),
		},



		{
			Optional: opt.Nothing[uint](),
		},
		{
			Optional: opt.Nothing[uint8](),
		},
		{
			Optional: opt.Nothing[uint16](),
		},
		{
			Optional: opt.Nothing[uint32](),
		},
		{
			Optional: opt.Nothing[uint64](),
		},
	}

	for testNumber, test := range tests {

		var worked bool = false

		test.Optional.WhenNothing(func(){

			worked = true
		})

		if !worked {
			t.Errorf("For test #%d, the call to the method did not seem work.", testNumber)
			t.Logf("WORKED: %t", worked)
			t.Logf("OPTIONAL: (%T) %#v", test.Optional, test.Optional)
	//////////////// CONTINUE
			continue
		}
	}
}

func TestOptional_WhenNothing_something(t *testing.T) {


	tests := []struct{
		Optional interface{WhenNothing(func())}
	}{
		{
			Optional: opt.Something[string]("once twice thrice fource"),
		},



		{
			Optional: opt.Something[int](-1),
		},
		{
			Optional: opt.Something[int8](-101),
		},
		{
			Optional: opt.Something[int16](-10101),
		},
		{
			Optional: opt.Something[int32](-1010101),
		},
		{
			Optional: opt.Something[int64](-101010101),
		},



		{
			Optional: opt.Something[uint](1),
		},
		{
			Optional: opt.Something[uint8](101),
		},
		{
			Optional: opt.Something[uint16](10101),
		},
		{
			Optional: opt.Something[uint32](1010101),
		},
		{
			Optional: opt.Something[uint64](101010101),
		},
	}

	for testNumber, test := range tests {

		var worked bool = false

		test.Optional.WhenNothing(func(){

			worked = true
		})

		if worked {
			t.Errorf("For test #%d, the call to the method worked, but it should not have.", testNumber)
			t.Logf("WORKED: %t", worked)
			t.Logf("OPTIONAL: (%T) %#v", test.Optional, test.Optional)
	//////////////// CONTINUE
			continue
		}
	}
}
