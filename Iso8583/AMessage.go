package Iso8583

import (
	"bytes"
	"fmt"
	"errors"
)

type CreateFieldFunc func(int) IField

type AMessage struct {
	bitmap              *Bitmap
	MsgTemplate         *Template
	fields              map[int]IField
	CreateFieldCallback CreateFieldFunc
}

func NewAMessage(tmpl *Template) *AMessage {
	msg := &AMessage{MsgTemplate:tmpl,fields:make(map[int]IField),bitmap:NewBitmap(tmpl.BitmapFormatter)}
	msg.CreateFieldCallback = msg.CreateField

	return msg
}

func (msg *AMessage) PackedLength() int {

	length := msg.bitmap.PackedLength()
	for i := 2; i < 128; i++ {
		if msg.bitmap.IsFieldSet(i) {
			length += msg.fields[i].PackedLength()
		}
	}

	return length
}

func (msg *AMessage) ClearField(field int)  {
	msg.bitmap.SetField(field, false)
	delete(msg.fields,field)
}

func (msg *AMessage) IsFieldSet(field int) bool {
	return msg.bitmap.IsFieldSet(field)
}

func (msg *AMessage) ToMsg() []byte {

	packedLength := msg.PackedLength()
	data := make([]byte, packedLength)
	bmap := msg.bitmap.ToMsg()
	copy(data,bmap)
	offset := msg.bitmap.PackedLength()
	for i := 2; i < 128; i++ {
		if msg.bitmap.IsFieldSet(i) {
			field := msg.fields[i]
			copy(data[offset:],field.ToMsg())
			offset += field.PackedLength()
		}
	}

	return data
}

func (msg *AMessage) String() string {
	return msg.ToString("   ")
}

func (msg *AMessage) ToString(prefix string) string  {
	var buffer bytes.Buffer
	for i:= 2; i < 128; i++ {
		if msg.bitmap.IsFieldSet(i) {
			buffer.WriteString(msg.FieldsToString(i, prefix) + "\n")
		}
	}

	return buffer.String()
}

func (msg *AMessage) FieldsToString(field int, prefix string) string {
	return msg.fields[field].ToString(prefix)
}

func (msg *AMessage) CreateField(field int) IField {

	if _, ok := msg.MsgTemplate.template[field]; ok {
		return NewField(field, msg.MsgTemplate.template[field])
	}

	return nil
}

func (msg *AMessage) GetField(field int) (IField,error) {

	_, ok := msg.fields[field]
	if (!msg.bitmap.IsFieldSet(field)) || (! ok) {
		if msg.fields[field] = msg.CreateFieldCallback(field); msg.fields[field] != nil {
			msg.bitmap.SetField(field, true)
		} else {
			return nil, errors.New(fmt.Sprintf("Unable to create field number %d. Possibly because template does not have a defination for the field",
				field))
		}
	}

	return msg.fields[field],nil
}

func (msg *AMessage) Unpack(data []byte, startingOffset int) (int, error) {

	offset := msg.bitmap.Unpack(data,startingOffset)
	for i := 2; i < 128; i++ {
		if msg.bitmap.IsFieldSet(i) {
			field, err := msg.GetField(i)
			if err != nil {
				return 0, err
			}
			offset, err = field.Unpack(data, offset)
			if err != nil {
				return 0, err
			}
		}
	}

	return offset, nil
}

func (msg *AMessage) GetFieldValue(field int) string {

	if msg.bitmap.IsFieldSet(field) {
		return msg.fields[field].Value()
	}

	return ""
}

func (msg *AMessage) SetFieldValue(field int, value string) error  {

	if value == "" {
		msg.ClearField(field)
		return nil
	}

	fld, err := msg.GetField(field)
	if err != nil {
		return err
	}

	fld.SetValue(value)
	return nil
}

