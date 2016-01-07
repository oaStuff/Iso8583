package Iso8583

type CompositeMessage struct {
	AMessage
	CompositeFieldNumber int
}


func NewCompositeMessage(tmpl *Template, compFieldNumber int) *CompositeMessage {
	msg := &CompositeMessage{CompositeFieldNumber:compFieldNumber, AMessage : *NewAMessage(tmpl)}
	msg.CreateFieldCallback = msg.CreateField

	return msg
}

func (msg *CompositeMessage) CreateField(field int) IField {

	if _, ok := msg.MsgTemplate.template[field]; ok {
		return NewSubField(field,msg.MsgTemplate.template[field],msg.CompositeFieldNumber)
	}

	return nil
}