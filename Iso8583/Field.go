package Iso8583
import (
	"fmt"
	"strconv"
	"github.com/oaStuff/Iso8583/Utils"
	"github.com/oaStuff/Iso8583/FieldValidator"
)

type Field struct {
	fieldDescriptor IFieldDescriptor
	fieldValue string
	isComposite bool
	isSubField bool
	parentFieldNumber int
	fieldNumber int
	compositeValue *CompositeMessage
}


func NewField(fldNumber int, fldDescriptor IFieldDescriptor) *Field {

	field := &Field{fieldNumber:fldNumber,fieldDescriptor:fldDescriptor, isComposite:fldDescriptor.IsComposite(),isSubField:false}
	if field.isComposite {
		field.compositeValue = NewCompositeMessage(fldDescriptor.CompositeTemplate(),fldNumber)
	}

	return field
}

func NewSubField(fldNumber int, fldDescriptor IFieldDescriptor, parentNum int) *Field  {

	field := &Field{parentFieldNumber:parentNum,isSubField:true,fieldNumber:fldNumber,fieldDescriptor:fldDescriptor,
					isComposite:fldDescriptor.IsComposite()}
	if field.isComposite {
		field.compositeValue = NewCompositeMessage(fldDescriptor.CompositeTemplate(),fldNumber)
	}

	return field
}

func (fld *Field) FieldNumber() int {
	return fld.fieldNumber
}

func (fld *Field) PackedLength() int {
	if fld.isComposite {
		return fld.fieldDescriptor.LengthFormatter().LengthOfLengthIndicator() + fld.compositeValue.PackedLength()
	}

	return fld.fieldDescriptor.GetPackedLength(fld.Value())
}

func (fld *Field) Value() string {

	adjuster := fld.fieldDescriptor.Adjuster()
	if  adjuster == nil {
		return fld.fieldValue
	}

	return adjuster.Get(fld.fieldValue)
}

func (fld *Field) SubFieldValue(subField int) string {

	if !fld.isComposite {
		panic("Can not set a sub field value on a non composite field")
	}

	return fld.compositeValue.GetFieldValue(subField)
}

func (fld *Field) SetValue(value string)  {
	adjuster := fld.fieldDescriptor.Adjuster()
	if adjuster == nil {
		fld.fieldValue = value
	} else {
		fld.fieldValue = adjuster.Set(value)
	}

	if !fld.fieldDescriptor.Validator().IsValid(fld.fieldValue) {
		panic(fmt.Sprintf("the value '%s' is invalid for field [No. %d] expected format is '%s'", fld.fieldValue,fld.fieldNumber, fld.fieldDescriptor.Validator().Description()))
	}

	if !fld.fieldDescriptor.LengthFormatter().IsValidLength(fld.fieldDescriptor.Formatter().GetPackedLength(len(fld.fieldValue))) {
		panic(fmt.Sprintf("The field length is not valid for field number [%d]",fld.fieldNumber))
	}
}

func (fld *Field) SetSubFieldValue(subField int, value string)  {

	if !fld.isComposite {
		panic("Can not set a sub field value on a non composite field")
	}

	fld.compositeValue.SetFieldValue(subField,value)
}

func (fld *Field) ToMsg() []byte {

	if fld.isComposite {
		lenOfLenInd := fld.fieldDescriptor.LengthFormatter().LengthOfLengthIndicator()
		lengthOfField := fld.compositeValue.PackedLength()
		field := make([]byte, lenOfLenInd + lengthOfField)
		fld.fieldDescriptor.LengthFormatter().Pack(field,lengthOfField,0)
		fieldData := fld.compositeValue.ToMsg()
		copy(field[lenOfLenInd:],fieldData)

		return field
	}

	field,_ := fld.fieldDescriptor.Pack(fld.FieldNumber(),fld.Value())

	return field
}

func (fld *Field) Unpack(msg []byte, offset int) (int, error) {

	var newOffset int
	var err error
	var val string

	if fld.isComposite {
		lenOfLenInd := fld.fieldDescriptor.LengthFormatter().LengthOfLengthIndicator()
		unpackedLengthOfField := fld.fieldDescriptor.LengthFormatter().GetLengthOfField(msg,offset)
		utils.UnusedVariable(unpackedLengthOfField)
		newOffset, err = fld.compositeValue.Unpack(msg,offset + lenOfLenInd)
		if err != nil {
			return 0, err
		}
	} else {
		val, newOffset, err = fld.fieldDescriptor.Unpack(fld.fieldNumber,msg,offset)
		if err != nil {
			return 0, err
		}
		fld.SetValue(val)
	}

	return newOffset, nil
}

func (fld *Field) ToString(prefix string) string {
	if fld.isComposite {
		return fld.compositeValue.ToString(prefix)
	}

	if fld.isSubField {
		return fld.fieldDescriptor.Display(prefix,fmt.Sprintf("%d.%d",fld.parentFieldNumber,fld.fieldNumber),fld.Value())
	}

	return fld.fieldDescriptor.Display(prefix,strconv.Itoa(fld.fieldNumber),fld.Value())
}

func AsciiFixedField(fieldNumber int, packedLength int, validator fieldValidator.IFieldValidator) IField {
	return NewField(fieldNumber,AsciiFixed(packedLength,validator))
}

func AsciiVarField(fieldNumber int, lengthIndicator int, maxLength int, validator fieldValidator.IFieldValidator) IField {
	return NewField(fieldNumber,AsciiVar(lengthIndicator,maxLength,validator))
}

func BinFixedField(fieldNumber int, packedLength int) IField {
	return NewField(fieldNumber,BinaryFixed(packedLength))
}