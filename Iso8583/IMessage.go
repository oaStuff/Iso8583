package Iso8583

type IMessage interface  {
	ToMsg() []byte
	Unpack([]byte,int) (int, error)
}