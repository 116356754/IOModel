package IOModel

import (
	"encoding/json"
	"fmt"
)

// IOString is a string. It can be assigned to Value interface.
type IOString string

// Type returns TypeID of IOString. It's always TypeString.
func (s IOString) Type() TypeID {
	return TypeString
}

func (s IOString) asBool() (bool, error) {
	return false, castError(s.Type(), TypeBool)
}

func (s IOString) asInt() (int64, error) {
	return 0, castError(s.Type(), TypeInt)
}

func (s IOString) asFloat() (float64, error) {
	return 0, castError(s.Type(), TypeFloat)
}

func (s IOString) asString() (string, error) {
	return string(s), nil
}

func (s IOString) asBlob() ([]byte, error) {
	return nil, castError(s.Type(), TypeBlob)
}

// IOString returns JSON representation of a IOString. A string "a" will be marshaled
// as `"a"` (double quotes are included), not `a`. To obtain a plain string
// without double quotes, use ToString function.
func (s IOString) String() string {
	// the IOString return value is defined via the
	// default JSON serialization
	bytes, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("(unserializable string: %v)", err)
	}
	return string(bytes)
}
