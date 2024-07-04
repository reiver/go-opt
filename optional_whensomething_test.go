package opt_test

import (
	"github.com/reiver/go-opt"

	"testing"
)

func TestOptional_WhenSomething_string(t *testing.T) {


	tests := []struct{
		Optional interface{WhenSomething(func(string))}
		Expected                              string
	}{
		{
			Optional: opt.Something[string](""),
			Expected:                       "",
		},
		{
			Optional: opt.Something[string]("once twice thrice fource"),
			Expected:                       "once twice thrice fource",
		},
		{
			Optional: opt.Something[string]("apple banana cherry"),
			Expected:                       "apple banana cherry",
		},
		{
			Optional: opt.Something[string]("😈"),
			Expected:                       "😈",
		},
		{
			Optional: opt.Something[string]("۰۱۲۳۴۵۶۷۸۹"),
			Expected:                       "۰۱۲۳۴۵۶۷۸۹",
		},
	}

	for testNumber, test := range tests {

		var worked bool = false
		var value string = "-<([-RaNdOmNeSs])>-"

		test.Optional.WhenSomething(func(v string){

			worked = true
			value  = v
		})

		if !worked {
			t.Errorf("For test #%d, the call to the method did not seem work.", testNumber)
			t.Logf("WORKED: %t", worked)
			t.Logf("OPTIONAL: (%T) %#v", test.Optional, test.Optional)
	/////////////// CONTINUE
			continue
		}

		{
			expected := test.Expected
			actual   := value

			if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("EXPECTED: %q", actual)
			t.Logf("WORKED: %t", worked)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_WhenSomething_stringNothing(t *testing.T) {

	var worked bool = false

	opt.Nothing[string]().WhenSomething(func(v string){

		worked = true
	})

	if worked {
		t.Errorf("The call to the method worked, but it should not have.")
		t.Logf("WORKED: %t", worked)
//////////////// RETURN
		return
	}
}

func TestOptional_WhenSomething_int8(t *testing.T) {


	tests := []struct{
		Optional interface{WhenSomething(func(int8))}
		Expected                              int8
	}{
		{
			Optional: opt.Something[int8](-127),
			Expected:                     -127,
		},
		{
			Optional: opt.Something[int8](-126),
			Expected:                     -126,
		},
		{
			Optional: opt.Something[int8](-125),
			Expected:                     -125,
		},
		{
			Optional: opt.Something[int8](-124),
			Expected:                     -124,
		},
		{
			Optional: opt.Something[int8](-123),
			Expected:                     -123,
		},
		{
			Optional: opt.Something[int8](-122),
			Expected:                     -122,
		},
		{
			Optional: opt.Something[int8](-121),
			Expected:                     -121,
		},
		{
			Optional: opt.Something[int8](-120),
			Expected:                     -120,
		},
		{
			Optional: opt.Something[int8](-9),
			Expected:                     -9,
		},
		{
			Optional: opt.Something[int8](-8),
			Expected:                     -8,
		},
		{
			Optional: opt.Something[int8](-7),
			Expected:                     -7,
		},
		{
			Optional: opt.Something[int8](-6),
			Expected:                     -6,
		},
		{
			Optional: opt.Something[int8](-5),
			Expected:                     -5,
		},
		{
			Optional: opt.Something[int8](-4),
			Expected:                     -4,
		},
		{
			Optional: opt.Something[int8](-3),
			Expected:                     -3,
		},
		{
			Optional: opt.Something[int8](-2),
			Expected:                     -2,
		},
		{
			Optional: opt.Something[int8](-1),
			Expected:                     -1,
		},
		{
			Optional: opt.Something[int8](0),
			Expected:                     0,
		},
		{
			Optional: opt.Something[int8](1),
			Expected:                     1,
		},
		{
			Optional: opt.Something[int8](2),
			Expected:                     2,
		},
		{
			Optional: opt.Something[int8](3),
			Expected:                     3,
		},
		{
			Optional: opt.Something[int8](4),
			Expected:                     4,
		},
		{
			Optional: opt.Something[int8](5),
			Expected:                     5,
		},
		{
			Optional: opt.Something[int8](6),
			Expected:                     6,
		},
		{
			Optional: opt.Something[int8](7),
			Expected:                     7,
		},
		{
			Optional: opt.Something[int8](8),
			Expected:                     8,
		},
		{
			Optional: opt.Something[int8](9),
			Expected:                     9,
		},
		{
			Optional: opt.Something[int8](120),
			Expected:                     120,
		},
		{
			Optional: opt.Something[int8](121),
			Expected:                     121,
		},
		{
			Optional: opt.Something[int8](122),
			Expected:                     122,
		},
		{
			Optional: opt.Something[int8](123),
			Expected:                     123,
		},
		{
			Optional: opt.Something[int8](124),
			Expected:                     124,
		},
		{
			Optional: opt.Something[int8](125),
			Expected:                     125,
		},
		{
			Optional: opt.Something[int8](126),
			Expected:                     126,
		},
		{
			Optional: opt.Something[int8](127),
			Expected:                     127,
		},
	}

	for testNumber, test := range tests {

		var worked bool = false
		var value int8 = -101

		test.Optional.WhenSomething(func(v int8){

			worked = true
			value = v
		})

		if !worked {
			t.Errorf("For test #%d, the call to the method did not seem work.", testNumber)
			t.Logf("WORKED: %t", worked)
			t.Logf("OPTIONAL: (%T) %#v", test.Optional, test.Optional)
	//////////////// CONTINUE
			continue
		}

		{
			expected := test.Expected
			actual   := value

			if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("EXPECTED: %d", actual)
			t.Logf("WORKED: %t", worked)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_WhenSomething_int8Nothing(t *testing.T) {

	var worked bool = false

	opt.Nothing[int8]().WhenSomething(func(v int8){

		worked = true
	})

	if worked {
		t.Errorf("The call to the method worked, but it should not have.")
		t.Logf("WORKED: %t", worked)
//////////////// RETURN
		return
	}
}

func TestOptional_WhenSomething_uint8(t *testing.T) {


	tests := []struct{
		Optional interface{WhenSomething(func(uint8))}
		Expected                              uint8
	}{
		{
			Optional: opt.Something[uint8](0),
			Expected:                      0,
		},
		{
			Optional: opt.Something[uint8](1),
			Expected:                      1,
		},
		{
			Optional: opt.Something[uint8](2),
			Expected:                      2,
		},
		{
			Optional: opt.Something[uint8](3),
			Expected:                      3,
		},
		{
			Optional: opt.Something[uint8](4),
			Expected:                      4,
		},
		{
			Optional: opt.Something[uint8](5),
			Expected:                      5,
		},
		{
			Optional: opt.Something[uint8](6),
			Expected:                      6,
		},
		{
			Optional: opt.Something[uint8](7),
			Expected:                      7,
		},
		{
			Optional: opt.Something[uint8](8),
			Expected:                      8,
		},
		{
			Optional: opt.Something[uint8](9),
			Expected:                      9,
		},
		{
			Optional: opt.Something[uint8](120),
			Expected:                      120,
		},
		{
			Optional: opt.Something[uint8](121),
			Expected:                      121,
		},
		{
			Optional: opt.Something[uint8](122),
			Expected:                      122,
		},
		{
			Optional: opt.Something[uint8](123),
			Expected:                      123,
		},
		{
			Optional: opt.Something[uint8](124),
			Expected:                      124,
		},
		{
			Optional: opt.Something[uint8](125),
			Expected:                      125,
		},
		{
			Optional: opt.Something[uint8](126),
			Expected:                      126,
		},
		{
			Optional: opt.Something[uint8](127),
			Expected:                      127,
		},
		{
			Optional: opt.Something[uint8](250),
			Expected:                      250,
		},
		{
			Optional: opt.Something[uint8](251),
			Expected:                      251,
		},
		{
			Optional: opt.Something[uint8](252),
			Expected:                      252,
		},
		{
			Optional: opt.Something[uint8](253),
			Expected:                      253,
		},
		{
			Optional: opt.Something[uint8](254),
			Expected:                      254,
		},
		{
			Optional: opt.Something[uint8](255),
			Expected:                      255,
		},
	}

	for testNumber, test := range tests {

		var worked bool = false
		var value uint8 = 101

		test.Optional.WhenSomething(func(v uint8){

			worked = true
			value = v
		})

		if !worked {
			t.Errorf("For test #%d, the call to the method did not seem work.", testNumber)
			t.Logf("WORKED: %t", worked)
			t.Logf("OPTIONAL: (%T) %#v", test.Optional, test.Optional)
	//////////////// CONTINUE
			continue
		}

		{
			expected := test.Expected
			actual   := value

			if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("EXPECTED: %d", actual)
			t.Logf("WORKED: %t", worked)
	/////////////////////// CONTINUE
				continue
			}
		}
	}
}

func TestOptional_WhenSomething_uint8Nothing(t *testing.T) {

	var worked bool = false

	opt.Nothing[uint8]().WhenSomething(func(v uint8){

		worked = true
	})

	if worked {
		t.Errorf("The call to the method worked, but it should not have.")
		t.Logf("WORKED: %t", worked)
//////////////// RETURN
		return
	}
}
