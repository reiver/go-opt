package opt

import (
	"fmt"
	"encoding/json"
)

type jsonMarshaler interface {
	bool | string
}

// MarshalJSON makes it so json.Marshaler is implemented.
func (receiver Optional[jsonMarshal]) MarshalJSON() ([]byte, error) {
	if !receiver.something {
		return nil, fmt.Errorf("cannot marshal opt.Nothing[%T]()", receiver.value)
	}

	return json.Marshal(receiver.value)
}
