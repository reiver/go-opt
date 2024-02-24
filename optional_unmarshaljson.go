package opt

import (
	"encoding/json"

	"sourcecode.social/reiver/go-erorr"
)

const errNilReceiver = erorr.Error("opt: nil receiver")

var _ json.Unmarshaler = new(Optional[bool])
var _ json.Unmarshaler = new(Optional[string])

// UnmarshalJSON makes it so json.Unmarshaler is implemented.
func (receiver *Optional[T]) UnmarshalJSON(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch interface{}(&receiver.value).(type) {
	case *bool, *int64, *uint64, *string, *json.Number, json.Unmarshaler:
		// these are OK.
	default:
		return erorr.Errorf("opt: cannot unmarshal into something of type %T from JSON because pointer to parameterized type is ‘%T’ rather than ‘*bool’, ‘*string’, or ‘json.Unmarshaler’", receiver, &receiver.value)
	}

	{
		var dst T

		err := json.Unmarshal(data, &dst)
		if nil != err {
			return err
		}

		*receiver = Something(dst)
	}

	return nil
}
