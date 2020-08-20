package IOModel

import (
	"encoding/json"
	"fmt"
)

// IOInt is an integer. It can be assigned to Value interface. A value more than
// 2^53 - 1 cannot exactly be marshaled to JSON because some languages like
// JavaScript or Lua only has float as a numeric type.
type IOInt int64

// Type returns TypeID of IOInt. It's always TypeInt.
func (i IOInt) Type() TypeID {
	return TypeInt
}

func (i IOInt) asBool() (bool, error) {
	return false, castError(i.Type(), TypeBool)
}

func (i IOInt) asInt() (int64, error) {
	return int64(i), nil
}

func (i IOInt) asFloat() (float64, error) {
	return 0, castError(i.Type(), TypeFloat)
}

func (i IOInt) asString() (string, error) {
	return "", castError(i.Type(), TypeString)
}

func (i IOInt) asBlob() ([]byte, error) {
	return nil, castError(i.Type(), TypeBlob)
}

// IOString returns JSON representation of an IOInt.
func (i IOInt) String() string {
	// the IOString return value is defined via the
	// default JSON serialization
	bytes, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf("(unserializable int: %v)", err)
	}
	return string(bytes)
}
