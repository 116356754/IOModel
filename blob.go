package IOValue

import (
	"encoding/json"
	"fmt"
)

// IOBlob is a binary large object which may have any type of byte data.
// It can be assigned to Value interface.
type IOBlob []byte

// Type returns TypeID of IOBlob. It's always TypeBlob.
func (b IOBlob) Type() TypeID {
	return TypeBlob
}

func (b IOBlob) asBool() (bool, error) {
	return false, castError(b.Type(), TypeBool)
}

func (b IOBlob) asInt() (int64, error) {
	return 0, castError(b.Type(), TypeInt)
}

func (b IOBlob) asFloat() (float64, error) {
	return 0, castError(b.Type(), TypeFloat)
}

func (b IOBlob) asString() (string, error) {
	return "", castError(b.Type(), TypeString)
}

func (b IOBlob) asBlob() ([]byte, error) {
	return b, nil
}

// Stringreturns JSON representation of a IOBlob. IOBlob is marshaled as a string.
func (b IOBlob) String() string {
	// the IOString return value is defined via the
	// default JSON serialization
	bytes, err := json.Marshal(b)
	if err != nil {
		return fmt.Sprintf("(unserializable blob: %v)", err)
	}
	return string(bytes)
}
