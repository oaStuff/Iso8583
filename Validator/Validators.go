package validator

import (

	"github.com/oaStuff/Iso8583/FieldValidator"
)


func IsHex(value string) bool {
	v := &fieldValidator.HexFieldValidator{}
	return v.IsValid(value)
}