package lengthValidators

type FixedLengthValidator struct {
	length int
}

func NewFixedLengthValidator(len int) *FixedLengthValidator {
	return &FixedLengthValidator{length:len}
}

func (flv *FixedLengthValidator) IsValid(value string) bool {
	return len(value) == flv.length
}