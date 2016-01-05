package lengthFormatters

type ILengthFormatter interface {
	LengthOfLengthIndicator() int
	MaxLength() string
	Description() string
	GetLengthOfField([]byte,int) int
	Pack([]byte, int, int) int
	IsValidLength(int) bool
}