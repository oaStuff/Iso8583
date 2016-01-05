package Iso8583
import "log"

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
	log.Println("in composite create field")
	return nil
}