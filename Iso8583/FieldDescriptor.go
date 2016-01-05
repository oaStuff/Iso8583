package Iso8583
import (
	"com.aihe/Iso8583/LengthFormatters"
	"com.aihe/Iso8583/FieldValidator"
	"com.aihe/Iso8583/Formatter"
	"errors"
	"com.aihe/Iso8583/Utils"
	"fmt"
	"log"
)

type FieldDescriptor struct  {
	lengthFormatter lengthFormatters.ILengthFormatter
	validator fieldValidator.IFieldValidator
	formatter formatter.IFormatter
	adjuster Adjuster
	isComposite bool
	compositeTemplate *Template
}

func (fd * FieldDescriptor) Adjuster() Adjuster {
	return fd.adjuster
}

func (fd *FieldDescriptor) Formatter() formatter.IFormatter {
	return fd.formatter
}

func (fd *FieldDescriptor) LengthFormatter() lengthFormatters.ILengthFormatter {
	return fd.lengthFormatter
}

func (fd *FieldDescriptor) Validator() fieldValidator.IFieldValidator {
	return fd.validator
}

func (fd *FieldDescriptor) IsComposite() bool {
	return fd.isComposite
}

func (fd *FieldDescriptor) CompositeTemplate() *Template {
	return fd.compositeTemplate
}

func NewFieldDescriptor(lenFormatter lengthFormatters.ILengthFormatter, vald fieldValidator.IFieldValidator,
							format formatter.IFormatter, adj Adjuster) (*FieldDescriptor, error)  {

	if _, ok1 := format.(*formatter.BinaryFormatter); ok1 {
		if _, ok2 := vald.(*fieldValidator.HexFieldValidator); !ok2 {
			return nil, errors.New("A Binary field must have a hex validator")
		}
	}

	if _, ok1 := format.(*formatter.BcdFormatter); ok1 {
		if _, ok2 := vald.(*fieldValidator.NumericFieldValidator); !ok2 {
			return nil, errors.New("A Binary field must have a hex validator")
		}
	}

	return &FieldDescriptor{lengthFormatter:lenFormatter, validator:vald, formatter:format, adjuster:adj, isComposite:false}, nil
}

func NewCompositeFieldDescriptor(lenFormatter lengthFormatters.ILengthFormatter, compTemplate *Template) *FieldDescriptor {
	return &FieldDescriptor{lengthFormatter:lenFormatter, compositeTemplate:compTemplate, isComposite:true}
}

func create (lengthFormatter lengthFormatters.ILengthFormatter, fieldValidator fieldValidator.IFieldValidator,
				formatter formatter.IFormatter, adjuster Adjuster) IFieldDescriptor {

	fd,err := NewFieldDescriptor(lengthFormatter,fieldValidator,formatter,adjuster)
	if err != nil {
		log.Println(err)
	}

	return fd
}

func AsciiAlphaNumeric(length int) IFieldDescriptor {
	setAdjuster := NewFuncAdjuster(nil, func(value string) string { return utils.PadRight(value,length,' ') })
	return create(lengthFormatters.NewFixedLengthFormatter(length),fieldValidator.Ansp(), formatter.Ascii(),setAdjuster)
}

func AsciiAmount(length int) IFieldDescriptor {
	setAdjuster := NewFuncAdjuster(nil, func(value string) string {return utils.PadLeft(value, length,'0') })
	return create(lengthFormatters.NewFixedLengthFormatter(length),fieldValidator.Rev87AmountValidator(),
					formatter.Ascii(),setAdjuster)
}

func AsciiFixed(packedLength int, validator fieldValidator.IFieldValidator) IFieldDescriptor {
	return create(lengthFormatters.NewFixedLengthFormatter(packedLength),validator,formatter.Ascii(),nil)
}

func AsciiVar(lengthIndicator int, maxLength int, validator fieldValidator.IFieldValidator) IFieldDescriptor {
	return create(lengthFormatters.NewDefaultVariableLengthFormatter(lengthIndicator,maxLength),validator,
					formatter.Ascii(),nil)
}

func AsciiLlCharacter(maxLength int) IFieldDescriptor {
	return AsciiVar(2,maxLength,fieldValidator.Ans())
}

func AsciiLlNumeric(maxLength int) IFieldDescriptor {
	return AsciiVar(2,maxLength,fieldValidator.N())
}

func AsciiLllBinary(packedLength int) IFieldDescriptor {
	return create(lengthFormatters.NewDefaultVariableLengthFormatter(3,packedLength),fieldValidator.Hex(),
					formatter.Binary(),nil)
}

func AsciiLllCharacter(maxLength int) IFieldDescriptor {
	return AsciiVar(3,maxLength,fieldValidator.Ans())
}

func AsciiLllNumeric(maxLength int) IFieldDescriptor {
	return AsciiVar(3,maxLength,fieldValidator.N())
}

func AsciiNumeric(length int) IFieldDescriptor {
	setAjuster := NewFuncAdjuster(nil, func(value string) string {return utils.PadLeft(value,length,'0') })
	return create(lengthFormatters.NewFixedLengthFormatter(length),fieldValidator.N(),formatter.Ascii(),setAjuster)
}

func BcdFixed(length int) IFieldDescriptor {
	return create(lengthFormatters.NewFixedLengthFormatter(length),fieldValidator.N(),formatter.Bcd(),nil)
}

func BcdVar(lengthIndicator int, maxLength int, lengthFormatter formatter.IFormatter) IFieldDescriptor {
	return create(lengthFormatters.NewVariableLengthFormatter(lengthIndicator,maxLength,lengthFormatter),
				fieldValidator.N(),formatter.Bcd(),nil)
}

func BinaryFixed(packedLength int) IFieldDescriptor {
	return create(lengthFormatters.NewFixedLengthFormatter(packedLength), fieldValidator.Hex(),formatter.Binary(),nil)
}

func BinaryVar(lengthIndicator int, maxLength int, lengthFormatter formatter.IFormatter) IFieldDescriptor {
	return create(lengthFormatters.NewVariableLengthFormatter(lengthIndicator,maxLength,lengthFormatter),
					fieldValidator.Hex(),formatter.Binary(),nil)
}

func CompositeField(lengthIndicator int, maxLength int, template *Template) IFieldDescriptor {
	return NewCompositeFieldDescriptor(lengthFormatters.NewDefaultVariableLengthFormatter(lengthIndicator,maxLength),
										template)
}

func (fd *FieldDescriptor) Display(prefix, fieldNumber, value string) string  {
	fieldValue := ""
	if value != "" {
		fieldValue = "[" + value + "]"
	}

	return fmt.Sprintf("%s[%-8s %-4s %6s %04d] %03s %s",prefix,fd.lengthFormatter.Description(),
				fd.validator.Description(),fd.lengthFormatter.MaxLength(),
				fd.formatter.GetPackedLength(len(value)), fieldNumber,fieldValue)
}

func (fd *FieldDescriptor) GetPackedLength(value string) int {
	return fd.lengthFormatter.LengthOfLengthIndicator() + fd.formatter.GetPackedLength(len(value))
}

func (fd *FieldDescriptor) Pack(fieldNumber int, value string) ([]byte,error) {

	if !fd.lengthFormatter.IsValidLength(fd.formatter.GetPackedLength(len(value))) {
		return nil, errors.New(fmt.Sprintf("The field length is not valid for field number %d",fieldNumber))
	}

	if !fd.validator.IsValid(value) {
		return nil, errors.New(fmt.Sprintf("Invalid value for field number %d",fieldNumber))
	}

	lenOfLenInd := fd.lengthFormatter.LengthOfLengthIndicator()
	lengthOfField := fd.formatter.GetPackedLength(len(value))
	field := make([]byte,lenOfLenInd + lengthOfField)
	fd.lengthFormatter.Pack(field,len(value),0)
	fieldData,_ := fd.formatter.GetBytes(value)
	copy(field[lenOfLenInd:],fieldData)

	return field, nil
}


func (fd *FieldDescriptor) Unpack(fieldNumber int, data []byte, offset int) (retval string, newOffset int, err error)  {

	lenOfLenInd := fd.lengthFormatter.LengthOfLengthIndicator()
	unpackedLengthOfField := fd.lengthFormatter.GetLengthOfField(data,offset)
	lengthOfField := unpackedLengthOfField
	if _, ok := fd.formatter.(*formatter.BcdFormatter); ok {
		lengthOfField = fd.formatter.GetPackedLength(lengthOfField)
	}

	fieldData := data[(offset + lenOfLenInd) : (offset + lenOfLenInd + lengthOfField)]
	newOffset = offset + lengthOfField + lenOfLenInd
	value := fd.formatter.GetString(fieldData)
	if !fd.validator.IsValid(value) {
		retval = ""
		err = errors.New(fmt.Sprintf("Invalid field format for field number %d",fieldNumber))
		return
	}

	length := len(value)
	_, ok1 := fd.formatter.(*formatter.BinaryFormatter)
	_, ok2 := fd.formatter.(*formatter.BcdFormatter)

	if ok1 || ok2 {
		length = fd.formatter.GetPackedLength(length)
		if unpackedLengthOfField % 2 != 0 {
			value = value[1:]
		}
	}

	_, ok1 = fd.lengthFormatter.(*lengthFormatters.VariableLengthFormatter)
	if ok1 && !fd.lengthFormatter.IsValidLength(length) {
		retval = ""
		err = errors.New(fmt.Sprintf("Field is too long for field number %d", fieldNumber))
		return
	}

	retval = value
	err = nil
	return
}