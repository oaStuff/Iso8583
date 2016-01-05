package Iso8583
import "com.aihe/Iso8583/Formatter"

type TemplateDef map[int]IFieldDescriptor

type Template struct  {
	template TemplateDef
	MsgTypeFormatter formatter.IFormatter
	BitmapFormatter formatter.IFormatter
}

func NewDefaultTemplate() *Template {
	return &Template{template:make(TemplateDef),BitmapFormatter:formatter.Binary(),
						MsgTypeFormatter:formatter.Ascii()}
}

func NewTemplate(tmplateDef TemplateDef) *Template {
	return &Template{template:tmplateDef,BitmapFormatter:formatter.Binary(), MsgTypeFormatter:formatter.Ascii()}
}

func (tmp *Template) AddFieldDescriptor(fieldNumber int, descriptor IFieldDescriptor)  {
	tmp.template[fieldNumber] = descriptor
}