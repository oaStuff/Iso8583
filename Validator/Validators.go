package validator
import "com.aihe/Iso8583/FieldValidator"


func IsHex(value string) bool {
	v := &fieldValidator.HexFieldValidator{}
	return v.IsValid(value)
}