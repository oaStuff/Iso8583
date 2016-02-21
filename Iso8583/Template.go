package Iso8583

import (
	"github.com/oaStuff/Iso8583/Formatter"
)

type TemplateDef map[int]IFieldDescriptor

type Template struct  {
	templateDefinition TemplateDef
	MsgTypeFormatter   formatter.IFormatter
	BitmapFormatter    formatter.IFormatter
}

func NewDefaultTemplate() *Template {
	return &Template{templateDefinition:make(TemplateDef),BitmapFormatter:formatter.Binary(),
						MsgTypeFormatter:formatter.Ascii()}
}

func NewTemplate(tmplateDef TemplateDef) *Template {
	return &Template{templateDefinition:tmplateDef,BitmapFormatter:formatter.Binary(), MsgTypeFormatter:formatter.Ascii()}
}

func (tmp *Template) AddFieldDescriptor(fieldNumber int, descriptor IFieldDescriptor)  {
	tmp.templateDefinition[fieldNumber] = descriptor
}