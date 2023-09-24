package opt

import (
	"fmt"
	"encoding/json"
)

var _ json.Unmarshaler = new(Optional[bool])
var _ json.Unmarshaler = new(Optional[string])

// UnmarshalJSON makes it so json.Unmarshaler is implemented.
func (receiver *Optional[T]) UnmarshalJSON(data []byte) error {
	switch interface{}(receiver.value).(type) {
	case bool, string, json.Unmarshaler:
		// these are OK.
	default:
		return fmt.Errorf("cannot unmarshal into something of type %T from JSON because parameterized type is ‘%T’ rather than ‘bool’, ‘string’, or ‘json.Unmarshaler’", receiver, receiver.value)
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
