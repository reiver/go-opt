package opt

import (
	"encoding"
	"fmt"
	"reflect"

	"github.com/reiver/go-json"

	"github.com/reiver/go-erorr"
)

var _ json.Marshaler = Nothing[bool]()
var _ json.Marshaler = Nothing[string]()

// MarshalJSON makes it so json.Marshaler is implemented.
func (receiver Optional[T]) MarshalJSON() ([]byte, error) {

	switch interface{}(receiver.value).(type) {
	case json.Marshaler, encoding.TextMarshaler, bool, int, int8, int16, int32, int64, string, uint, uint8, uint16, uint32, uint64:
		// these are OK.
	default:
		var reflectedType reflect.Type = reflect.TypeOf(receiver.value)
		if nil == reflectedType {
			return nil, errBadReflection
		}
		if reflect.Struct != reflectedType.Kind() {
			return nil, erorr.Errorf("opt: cannot marshal something of type %T into JSON because parameterized type is ‘%T’ is not supported", receiver, receiver.value)
		}
	}

	if receiver.IsNothing() {
		return nil, json.ErrEmpty(fmt.Sprintf("opt: cannot marshal opt.Nothing[%T]() into JSON", receiver.value))
	}

	return json.Marshal(receiver.value)
}
