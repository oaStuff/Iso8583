package Iso8583
import (
	"github.com/oaStuff/Iso8583/Formatter"
	"github.com/oaStuff/Iso8583/LengthFormatters"
	"github.com/oaStuff/Iso8583/FieldValidator"
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