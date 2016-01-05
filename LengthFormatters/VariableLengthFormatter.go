package lengthFormatters
import (
	"com.aihe/Iso8583/Formatter"
	"strings"
	"strconv"
	"com.aihe/Iso8583/Utils"
)

type VariableLengthFormatter struct  {
	lengthFormatter formatter.IFormatter
	lengthIndicator int
	maxLength int
	lenOfLenIndicator int
}

func NewVariableLengthFormatter(lenIndicator int, maxLen int, lenFormatter formatter.IFormatter) *VariableLengthFormatter {

	vlFormatter := &VariableLengthFormatter{lengthIndicator:lenIndicator,maxLength:maxLen,lengthFormatter:lenFormatter}
	vlFormatter.lenOfLenIndicator = vlFormatter.lengthFormatter.GetPackedLength(lenIndicator)
	return vlFormatter
}

func NewDefaultVariableLengthFormatter(lenIndicator int, maxLen int) *VariableLengthFormatter {
	return NewVariableLengthFormatter(lenIndicator,maxLen,formatter.NewAsciiFormatter())
}

func (vlf *VariableLengthFormatter) LengthOfLengthIndicator() int {
	return vlf.lenOfLenIndicator
}

func (vlf *VariableLengthFormatter) MaxLength() string {
	return ".." + strconv.Itoa(vlf.maxLength)
}

func (vlf *VariableLengthFormatter) Description() string {
	return strings.Repeat("L",vlf.lengthIndicator) + "Var"
}

func (vlf *VariableLengthFormatter) GetLengthOfField(msg []byte, offset int) int {
	lenStr := vlf.lengthFormatter.GetString(msg[offset:(offset + vlf.lenOfLenIndicator)])
	len, _ := strconv.Atoi(lenStr)
	return len
}

func (vlf *VariableLengthFormatter) Pack(msg []byte, length int, offset int) int {
	lengthStr := utils.PadLeft(strconv.Itoa(length),vlf.lenOfLenIndicator,'0')
	header, _ := vlf.lengthFormatter.GetBytes(lengthStr)
	copy(msg[offset:],header)

	return offset + vlf.lenOfLenIndicator
}

func (vlf *VariableLengthFormatter) IsValidLength(packedLen int) bool {
	return packedLen <= vlf.maxLength
}
