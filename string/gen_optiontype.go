package optstring

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	errNilReceiver  = errors.New("Nil Receiver")
	errNone         = errors.New("optstring.None()")
	errNoneNullable = errors.New("optstring.NoneNullable()")
	errNull         = errors.New("optstring.Null()")
)

func (receiver *NullableType) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == src {
		*receiver = Null()
		return nil
	}

	switch t := src.(type) {
	case NullableType:
		*receiver = t
		return nil
	case Type:
		switch t {
		case None():
			*receiver = NoneNullable()
		default:
			datum, err := t.String()
			if nil != err {
				return fmt.Errorf("Problem unwrapping %T: (%T) %v", t, err, err)
			}
			*receiver = SomeNullable(datum)
		}
		return nil
	case string:
		*receiver = SomeNullable(t)
		return nil
	case []byte:
		s := string(t)
		*receiver = SomeNullable(s)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}

func (receiver NullableType) String() (string, error) {
	if NoneNullable() == receiver {
		return "", errNoneNullable
	}
	if Null() == receiver {
		return "", errNull
	}

	return receiver.value, nil
}

type NullableType struct {
	loaded bool
	null   bool
	value  string
}

func (receiver NullableType) MarshalJSON() ([]byte, error) {
	if NoneNullable() == receiver {
		return nil, errNoneNullable
	}
	if Null() == receiver {
		return json.Marshal(nil)
	}

	return json.Marshal(receiver.value)
}

func (receiver *NullableType) UnmarshalJSON(b []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 0 == bytes.Compare(b, []byte{'n','u','l','l'}) {
		*receiver = Null()
		return nil
	}

	var target string

	if err := json.Unmarshal(b, &target); nil != err {
		return err
	}

	*receiver = SomeNullable(target)

	return nil
}

func (receiver NullableType) Value() (driver.Value, error) {
	if NoneNullable() == receiver {
		return nil, errNoneNullable
	}
	if Null() == receiver {
		return nil, nil
	}

	return receiver.value, nil
}

func NoneNullable() NullableType {
	return NullableType{}
}

func SomeNullable(value string) NullableType {
	return NullableType{
		value:  value,
		loaded: true,
	}
}

func Null() NullableType {
	return NullableType{
		null:   true,
		loaded: true,
	}
}

func (receiver *Type) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch t := src.(type) {
	case Type:
		*receiver = t
		return nil
	case string:
		*receiver = Some(t)
		return nil
	case []byte:
		s := string(t)
		*receiver = Some(s)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}

func (receiver Type) String() (string, error) {
	if None() == receiver {
		return "", errNone
	}

	return receiver.value, nil
}

type Type struct {
	loaded bool
	value  string
}

func (receiver Type) MarshalJSON() ([]byte, error) {
	if None() == receiver {
		return nil, errNone
	}

	return json.Marshal(receiver.value)
}

func (receiver *Type) UnmarshalJSON(b []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 0 == bytes.Compare(b, []byte{'n','u','l','l'}) {
		return fmt.Errorf("Cannot unmarshal %q into %T.", string(b), *receiver)
	}

	var target string

	if err := json.Unmarshal(b, &target); nil != err {
		return err
	}

	*receiver = Some(target)

	return nil
}

func (receiver Type) Value() (driver.Value, error) {
	if None() == receiver {
		return nil, errNone
	}

	return receiver.value, nil
}

func None() Type {
	return Type{}
}

func Some(value string) Type {
	return Type{
		value:  value,
		loaded: true,
	}
}
