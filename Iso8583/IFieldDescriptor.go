package Iso8583
import (
	"com.aihe/Iso8583/Formatter"
	"com.aihe/Iso8583/LengthFormatters"
	"com.aihe/Iso8583/FieldValidator"
)

type IFieldDescriptor interface  {

	Adjuster() Adjuster
	Formatter() formatter.IFormatter
	LengthFormatter() lengthFormatters.ILengthFormatter
	Validator() fieldValidator.IFieldValidator
	IsComposite() bool
	CompositeTemplate() *Template
	Display(string,string,string) string
	GetPackedLength(string) int
	Pack(int, string) ([]byte,error)
	Unpack(int, []byte, int) (string, int, error)
}