package IOValue

// IONull corresponds to null in JSON. It can be assigned to Value interface.
// IONull is provided for IONull Object pattern and it should always be used
// instead of nil.
type IONull struct{}

// Type returns TypeID of IONull. It's always TypeNull.
func (n IONull) Type() TypeID {
	return TypeNull
}

func (n IONull) asBool() (bool, error) {
	return false, castError(n.Type(), TypeBool)
}

func (n IONull) asInt() (int64, error) {
	return 0, castError(n.Type(), TypeInt)
}

func (n IONull) asFloat() (float64, error) {
	return 0, castError(n.Type(), TypeFloat)
}

func (n IONull) asString() (string, error) {
	return "", castError(n.Type(), TypeString)
}

func (n IONull) asBlob() ([]byte, error) {
	return nil, castError(n.Type(), TypeBlob)
}

// MarshalJSON marshals IONull to JSON.
func (n IONull) MarshalJSON() ([]byte, error) {
	// the JSON serialization is defined via the IOString()
	// return value as defined below
	return []byte(n.String()), nil
}

// IOString returns JSON representation of a IONull.
func (n IONull) String() string {
	return "null"
}
