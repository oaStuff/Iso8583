package formatter

type IFormatter interface {

	GetBytes(value string) ([]byte, error)
	GetString([]byte) (string)
	GetPackedLength(unpackedLength int) (int)
}