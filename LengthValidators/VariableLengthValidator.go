package lengthValidators

type VariableLengthValidator struct  {
	maxLength int
	minLength int
}

func NewVariableLengthValidator(minimumLength, maximumLength int) *VariableLengthValidator {
	return &VariableLengthValidator{minLength:minimumLength, maxLength:maximumLength}
}

func (vlv *VariableLengthValidator) IsValid(value string) bool {
	return (len(value) >= vlv.minLength) && (len(value) <= vlv.maxLength)
}