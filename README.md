This is an implementation of the ISO8583 standard in GO.

#### Using default Iso8583 implementation

```go
func main() {
	// Creates a new Iso8583 message
	msg := NewIso8583()

    //Set various fields as required
	msg.SetFieldValue(4,"4000")
    msg.SetFieldValue(70,"677")

    //You can even set a sub field on field 127 (this conforms to the Postilion switch standard)
    msg.SetSubFieldValue(127,2,"999999999")
    msg.SetSubFieldValue(127,6,"p1")
    msg.SetSubFieldValue(127,34,"98")

    //You can get the raw bytes of the message
    data := msg.ToMsg()

    //Print out the message in the default formatted message style
    log.Printf("\n%s",msg.ToString(""))

    //retrieve values from the message
    log.Println("Data stored is ", msg.GetFieldValue(2))

    //View the hex dump of the message also
    log.Println("hex dump data is ")
    log.Printf("\n%v",hex.Dump(data))
}
```



#### Specifying your own template

```go
func main() {
	//Define the template

	tmplDef := TemplateDef{
        2 : AsciiVar(2,20,fieldValidator.N()),
        4 : AsciiNumeric(12),
        70 : AsciiAlphaNumeric(3),
    }

    //Create the template and use it in creating the message
    tmpl:= NewTemplate(tmplDef)
    msg := NewAMessage(tmpl)

    //Set various fields as required
    //You can only set values for the field you have defined in the template
	msg.SetFieldValue(4,"4000")
    msg.SetFieldValue(70,"677")

    //You can get the raw bytes of the message
    data := msg.ToMsg()

    //Print out the message in the default formatted message style
    log.Printf("\n%s",msg.ToString(""))

    //retrieve values from the message
    log.Println("Data stored is ", msg.GetFieldValue(2))

    //View the hex dump of the message also
    log.Println("hex dump data is ")
    log.Printf("\n%v",hex.Dump(data))
}
```

