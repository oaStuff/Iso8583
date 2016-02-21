package Iso8583
import (
	"testing"
	"encoding/hex"
	"github.com/oaStuff/Iso8583/Formatter"
)

func TestBitmapDefaultFormatter(t *testing.T)  {
	bitmap := NewBitmapWithDefaultFormatter()
	if _, ok := bitmap.formatter.(*formatter.BinaryFormatter); !ok {
		t.Error("Default formatter ought to be BinaryFormatter")
	}
}

func TestExtendedBitmap(t *testing.T)  {
	Bitmap := NewBitmapWithDefaultFormatter()
	Bitmap.SetField(100,true)
	if !Bitmap.IsExtendedBitmap() {
		t.Error("Extended bitmap wrong")
	}
}

func TestUnpacking(t *testing.T)  {
	data, _ := hex.DecodeString("F3000000000000000000000000000001FAD1000820")
	Bitmap := NewBitmap(formatter.Binary())
	nextOffset := Bitmap.Unpack(data,0)
	if !Bitmap.IsExtendedBitmap() {
		t.Error("Bitmap unpacking should have produced an extended bitmap")
	}

	if nextOffset < Bitmap.formatter.GetPackedLength(16) {
		t.Error("Bitmap unpacking returned an invalid offset")
	}
}
