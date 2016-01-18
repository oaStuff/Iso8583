package Iso8583
import (
	"testing"
	"log"
	"encoding/hex"
)

func TestAMessage(t *testing.T)  {



//	tmplDef := TemplateDef{
//		2:AsciiVar(2,20,fieldValidator.N()),
//		4:AsciiNumeric(12),
//		70:AsciiAlphaNumeric(3),
//	}
//	tmpl:= NewTemplate(tmplDef)
//	msg := NewAMessage(tmpl)
	msg := NewIso8583()
	if err := msg.SetFieldValue(2,"123456789012345678"); err != nil {
		t.Error(err)
	}
	msg.SetFieldValue(4,"4000")
	msg.SetFieldValue(70,"677")
	msg.SetSubFieldValue(127,2,"999999999")
	msg.SetSubFieldValue(127,6,"p1")
	msg.SetSubFieldValue(127,34,"98")
	data := msg.ToMsg()
	log.Printf("\n%s",msg.ToString(""))
	log.Println("packed len is ", msg.PackedLength())
	log.Println("Data stored is ", msg.GetFieldValue(2))
	log.Println("network data == ", (data))
	log.Println("hex dump data is ")
	log.Printf("\n%v",hex.Dump(data))
}
