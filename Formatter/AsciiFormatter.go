package formatter

type AsciiFormatter struct  {

}


func NewAsciiFormatter() IFormatter  {
	return &AsciiFormatter{}
}


func (af *AsciiFormatter) GetBytes(value string) ([]byte, error) {
	return []byte(value), nil
}

func (af *AsciiFormatter) GetString(value []byte) (string) {
	return string(value)
}


func (af *AsciiFormatter) GetPackedLength(unpackedLength int) (int) {
	return unpackedLength
}