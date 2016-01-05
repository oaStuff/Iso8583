package fieldValidator


func Alpha() IFieldValidator {
	return NewAlphaFieldValidator()
}

func A() IFieldValidator {
	return Alpha()
}

func AlphaNumeric() IFieldValidator {
	return NewAlphaNumericFieldValidator()
}

func An() IFieldValidator {
	return AlphaNumeric()
}

func AlphaNumericAndSpace() IFieldValidator {
	return NewAlphaNumericAndSpaceFieldValidator()
}

func Ansp() IFieldValidator {
	return AlphaNumericAndSpace()
}

func AlphaNumericPrintable() IFieldValidator {
	return NewAlphaNumericPrintableFieldValidator()
}

func Anp() IFieldValidator {
	return AlphaNumericPrintable()
}

func AlphaNumericSpecial() IFieldValidator {
	return NewAlphaNumericSpecialFieldValidator()
}

func Ans() IFieldValidator {
	return AlphaNumericSpecial()
}

func Hex() IFieldValidator {
	return NewHexFieldValidator()
}

func None() IFieldValidator {
	return NewNoneFieldValidator()
}

func Numeric() IFieldValidator {
	return NewNumericFieldValidator()
}

func N() IFieldValidator {
	return Numeric()
}

func Rev87AmountValidator() IFieldValidator {
	return NewRev87AmountFieldValidator()
}