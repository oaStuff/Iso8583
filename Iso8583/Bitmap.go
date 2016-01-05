package Iso8583
import "com.aihe/Iso8583/Formatter"



type Bitmap struct {
	bits []bool
	formatter formatter.IFormatter
}


func NewBitmap(formatter formatter.IFormatter) *Bitmap  {
	return &Bitmap{formatter:formatter,bits:make([]bool,128)}
}

func NewBitmapWithDefaultFormatter() *Bitmap {
	return &Bitmap{formatter:formatter.Binary(),bits:make([]bool,128)}
}

func (bm *Bitmap) IsExtendedBitmap() bool  {
	return bm.IsFieldSet(1)
}

func (bm *Bitmap) PackedLength() int {
	if bm.IsExtendedBitmap() {
		return bm.formatter.GetPackedLength(32)
	}

	return bm.formatter.GetPackedLength(16)
}

func (bm *Bitmap) GetField(field int) bool {
	return bm.IsFieldSet(field)
}

func (bm *Bitmap) IsFieldSet(field int) bool {
	return bm.bits[field - 1]
}

func (bm *Bitmap) SetField(field int, value bool)  {
	bm.bits[field - 1] = value
	bm.bits[0] = false
	for i:= 64; i < 128; i++ {
		if bm.bits[i] {
			bm.bits[0] = true
			return
		}
	}
}

func (bm *Bitmap) ToMsg() []byte {

	var lengthOfBitmap int
	if bm.IsExtendedBitmap() {
		lengthOfBitmap = 16
	} else {
		lengthOfBitmap = 8
	}

	data := make([]byte, lengthOfBitmap)

	for i := 0; i < lengthOfBitmap; i++ {
		for j := 0; j < 8; j++ {
			if bm.bits[i * 8 + j] {
				data[i] = byte(data[i] | (128 / (1 << uint(j))))
			}
		}
	}

	if _, ok := bm.formatter.(*formatter.BinaryFormatter); ok {
		return data
	}

	binaryFormatter := formatter.Binary()
	bitmapString := binaryFormatter.GetString(data)
	retData, _ :=bm.formatter.GetBytes(bitmapString)
	return retData
}

func (bm *Bitmap) Unpack(msg []byte, offset int) int {

	lengthOfBitmap := bm.formatter.GetPackedLength(16)
	if _, ok := bm.formatter.(*formatter.BinaryFormatter); ok {
		if msg[offset] >= 128 {
			lengthOfBitmap += 8
		}
	} else {
		if msg[offset] >= 0x38 {
			lengthOfBitmap += 16
		}
	}

	bitmapData := msg[offset:offset + lengthOfBitmap]
	if _, ok := bm.formatter.(*formatter.BinaryFormatter); !ok {
		binaryFormatter := formatter.Binary()
		value := bm.formatter.GetString(bitmapData)
		bitmapData, _ = binaryFormatter.GetBytes(value)
	}

	bitmapDataLength := len(bitmapData)
	for i := 0; i < bitmapDataLength; i++ {
		for j := 0; j < 8; j++ {
			bm.bits[i * 8 + j] = (bitmapData[i] & (128 / (1 << uint(j)))) > 0
		}
	}

	return offset + lengthOfBitmap
}