package Iso8583
import (
	"testing"
	"log"
	"encoding/hex"
	"fmt"
	"com.aihe/Iso8583/Formatter"
)

func TestAMessage(t *testing.T)  {

	d := formatter.BcdFormatter{}
	fmt.Println(d.GetPackedLength(11))

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
	msg.SetFieldValue(70,"AB1")
	data := msg.ToMsg()
	log.Printf("\n%s",msg.ToString(""))
	log.Println("packed len is ", msg.PackedLength())
	log.Println("Data stored is ", msg.GetFieldValue(2))
	log.Println("network data == ", (data))
	log.Println("hex dump data is ")
	log.Printf("\n%v",hex.Dump(data))
}
