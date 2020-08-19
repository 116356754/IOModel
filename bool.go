package IOValue

import (
	"fmt"
)

// IOBool is a boolean value. It can be assigned to Value interface.
type IOBool bool

const (
	// True is a constant having true value of IOBool type.
	True IOBool = true

	// False is a constant having false value of IOBool type.
	False IOBool = false
)

// Type returns TypeID of IOBool. It's always TypeBool.
func (b IOBool) Type() TypeID {
	return TypeBool
}

func (b IOBool) asBool() (bool, error) {
	return bool(b), nil
}

func (b IOBool) asInt() (int64, error) {
	return 0, castError(b.Type(), TypeInt)
}

func (b IOBool) asFloat() (float64, error) {
	return 0, castError(b.Type(), TypeFloat)
}

func (b IOBool) asString() (string, error) {
	return "", castError(b.Type(), TypeString)
}

func (b IOBool) asBlob() ([]byte, error) {
	return nil, castError(b.Type(), TypeBlob)
}

// IOString returns JSON representation of a IOBool, which is true or false.
func (b IOBool) String() string {
	return fmt.Sprintf("%#v", b)
}
