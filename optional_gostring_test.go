package opt_test

import (
	"sourcecode.social/reiver/go-opt"

	"fmt"

	"testing"
)

func TestOptional_GoString(t *testing.T) {

	tests := []struct{
		Value any
		Expected string
	}{
		{
			Value:                           "",
			Expected: `opt.Something[string]("")`,
		},
		{
			Value:                           "once twice thrice fource",
			Expected: `opt.Something[string]("once twice thrice fource")`,
		},
		{
			Value:                           "apple banana cherry",
			Expected: `opt.Something[string]("apple banana cherry")`,
		},



		{
			Value:                   uint8 (0x0),
			Expected: `opt.Something[uint8](0x0)`,
		},
		{
			Value:                   uint8 (0x1),
			Expected: `opt.Something[uint8](0x1)`,
		},
		{
			Value:                   uint8 (0x2),
			Expected: `opt.Something[uint8](0x2)`,
		},
		{
			Value:                   uint8 (0xfe),
			Expected: `opt.Something[uint8](0xfe)`,
		},
		{
			Value:                   uint8 (0xff),
			Expected: `opt.Something[uint8](0xff)`,
		},



		{
			Value:                   uint16 (0x0),
			Expected: `opt.Something[uint16](0x0)`,
		},
		{
			Value:                   uint16 (0x1),
			Expected: `opt.Something[uint16](0x1)`,
		},
		{
			Value:                   uint16 (0x2),
			Expected: `opt.Something[uint16](0x2)`,
		},
		{
			Value:                   uint16 (0xfe),
			Expected: `opt.Something[uint16](0xfe)`,
		},
		{
			Value:                   uint16 (0xff),
			Expected: `opt.Something[uint16](0xff)`,
		},
		{
			Value:                   uint16 (0x100),
			Expected: `opt.Something[uint16](0x100)`,
		},
		{
			Value:                   uint16 (0x101),
			Expected: `opt.Something[uint16](0x101)`,
		},
		{
			Value:                   uint16 (0x102),
			Expected: `opt.Something[uint16](0x102)`,
		},
		{
			Value:                   uint16 (0xfffe),
			Expected: `opt.Something[uint16](0xfffe)`,
		},
		{
			Value:                   uint16 (0xffff),
			Expected: `opt.Something[uint16](0xffff)`,
		},



		{
			Value:                   struct { A string; B int }{A:"joeblow",B:7},
			Expected: `opt.Something[struct { A string; B int }](struct { A string; B int }{A:"joeblow", B:7})`,
		},
	}

	for testNumber, test := range tests {

		op := opt.Something(test.Value)
		gostring := op.GoString()


		{
			expected := test.Expected
			actual   := gostring

			if expected != actual {
				t.Errorf("For test #%d, the actual go-string is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE-TYPE: %T", test.Value)
				t.Logf("VALUE: %#v", test.Value)
	//////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_GoString_nothing(t *testing.T) {

	tests := []struct{
		Optional fmt.GoStringer
		Expected string
	}{
		{
			Optional:  opt.Nothing[string](),
			Expected: `opt.Nothing[string]()`,
		},



		{
			Optional:  opt.Nothing[int8](),
			Expected: `opt.Nothing[int8]()`,
		},
		{
			Optional:  opt.Nothing[int16](),
			Expected: `opt.Nothing[int16]()`,
		},
		{
			Optional:  opt.Nothing[int32](),
			Expected: `opt.Nothing[int32]()`,
		},
		{
			Optional:  opt.Nothing[int64](),
			Expected: `opt.Nothing[int64]()`,
		},



		{
			Optional:  opt.Nothing[uint8](),
			Expected: `opt.Nothing[uint8]()`,
		},
		{
			Optional:  opt.Nothing[uint16](),
			Expected: `opt.Nothing[uint16]()`,
		},
		{
			Optional:  opt.Nothing[uint32](),
			Expected: `opt.Nothing[uint32]()`,
		},
		{
			Optional:  opt.Nothing[uint64](),
			Expected: `opt.Nothing[uint64]()`,
		},
	}

	for testNumber, test := range tests {

		expected := test.Expected
		actual   := test.Optional.GoString()

		if expected != actual {
			t.Errorf("For test #%d, the actual go-string value is not what was expected.", testNumber)
			t.Logf("EXPECTED GO-STRING: %q", expected)
			t.Logf("ACTUAL   GO-STRING: %q", actual)
			t.Logf("TYPE: %T", test.Optional)
	/////////////// CONTINUE
			continue
		}
	}
}
