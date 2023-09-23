package opt

import (
	"fmt"
	"encoding/json"
)

var _ json.Marshaler = Nothing[bool]()
var _ json.Marshaler = Nothing[string]()

// MarshalJSON makes it so json.Marshaler is implemented.
func (receiver Optional[T]) MarshalJSON() ([]byte, error) {
	switch interface{}(receiver.value).(type) {
	case bool, string:
		// these are OK.
	default:
		return nil, fmt.Errorf("cannot marshal something of type %T into JSON because parameterized type is ‘%T’ rather than ‘bool’ or ‘string’", receiver, receiver.value)
	}

	if receiver.isnothing() {
		return nil, fmt.Errorf("cannot marshal opt.Nothing[%T]() into JSON", receiver.value)
	}

	return json.Marshal(receiver.value)
}
