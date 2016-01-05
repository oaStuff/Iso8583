package fieldValidator

type Rev87AmountFieldValidator struct {

}

func NewRev87AmountFieldValidator() *Rev87AmountFieldValidator {
	return &Rev87AmountFieldValidator{}
}

func (rafv *Rev87AmountFieldValidator) Description() string {
	return "amt"
}

func (rafv *Rev87AmountFieldValidator) IsValid(value string) bool {
	first := value[0]
	if (first != 'C') && (first != 'D') {
		return false
	}

	numeric := NewNumericFieldValidator()
	return numeric.IsValid(value[1:])
}