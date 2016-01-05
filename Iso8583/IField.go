package Iso8583

type IField interface  {
	FieldNumber() int
	PackedLength() int
	Value() string
	SetValue(string)
	ToMsg() []byte
	Unpack([]byte, int) (int, error)
	ToString(string) string
}