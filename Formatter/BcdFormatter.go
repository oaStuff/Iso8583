package formatter
import (
	"encoding/hex"
	"github.com/oaStuff/Iso8583/Utils"
)

type BcdFormatter struct  {

}


func NewBcdFormatter() IFormatter {
	return &BcdFormatter{}
}

func (bf *BcdFormatter) GetBytes(value string) ([]byte, error) {

	if len(value) % 2 == 1 {
		value = utils.PadLeft(value,len(value)+1,'0')
	}

	chars := []byte(value)
	charLen := len(chars) / 2
	bytes := make([]byte,charLen)
	for i := 0; i < charLen; i++ {
		highNibble := chars[2 * i]
		lowNibble := chars[2 * i + 1]
		bytes[i] = byte(byte(highNibble << 4) | lowNibble)
	}

	return bytes, nil
}

func (bf *BcdFormatter) GetString(value []byte) (string) {
	return hex.EncodeToString(value)
}


func (bf *BcdFormatter) GetPackedLength(unpackedLength int) (int) {

	tmp := int(unpackedLength / 2)
	if tmp % 2 != 0 {
		tmp++
	}

	return tmp
}