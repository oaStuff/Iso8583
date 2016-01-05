package main
import (
	"fmt"
	"com.aihe/Iso8583/Iso8583"
)

func main()  {
	fmt.Println("starting")
	dd := Iso8583.NewBitmap(nil)
	fmt.Println(dd)
}
