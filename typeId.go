package IOModel

// TypeID is an ID of a type. A unique value is assigned to each type.
type TypeID int

//目前只支持如下6种数据类型，后期可能会有Date等类型的加入
const (
	typeUnknown TypeID = iota
	// TypeNull is a TypeID of IONull.
	TypeNull
	// TypeBool is a TypeID of IOBool.
	TypeBool
	// TypeInt is a TypeID of IOInt.
	TypeInt
	// TypeUInt is a TypeID of IOInt.
	TypeFloat
	// TypeFloat is a TypeID of Double.
	TypeString
	// TypeBlob is a TypeID of IOBlob.
	TypeBlob
)

func (t TypeID) String() string {
	switch t {
	case TypeNull:
		return "null"
	case TypeBool:
		return "bool"
	case TypeInt:
		return "int"
	case TypeFloat:
		return "float"
	case TypeString:
		return "string"
	case TypeBlob:
		return "blob"
	default:
		return "unknown"
	}
}
