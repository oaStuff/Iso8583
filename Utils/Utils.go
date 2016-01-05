package utils


func PadLeft(str string, totalWidth int, padChar byte) string {
	n := totalWidth - len(str)
	for i := 0; i < n; i++ {
		str = string(padChar) + str
	}

	return str
}

func PadRight(str string, totalWidth int, padChar byte) string {
	n := totalWidth - len(str)
	for i := 0; i < n; i++ {
		str = str + string(padChar)
	}

	return str
}

func UnusedVariable(value interface{})  {

}