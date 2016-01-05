package lengthFormatters
import "fmt"

type FixedLengthFormatter struct  {
	packedLength int
}

func NewFixedLengthFormatter(packedLen int) *FixedLengthFormatter {
	return &FixedLengthFormatter{packedLength:packedLen}
}

func (flf *FixedLengthFormatter) LengthOfLengthIndicator() int {
	return 0
}

func (flf *FixedLengthFormatter) MaxLength() string {
	return fmt.Sprintf("%d", flf.packedLength)
}

func (flf *FixedLengthFormatter) Description() string {
	return "Fixed"
}

func (flf *FixedLengthFormatter) GetLengthOfField(msg []byte, offset int) int {
	return flf.packedLength
}

func (flf *FixedLengthFormatter) Pack(msg []byte, length int, offset int) int {
	return offset
}

func (flf *FixedLengthFormatter) IsValidLength(packedLen int) bool {
	return flf.packedLength == packedLen
}
