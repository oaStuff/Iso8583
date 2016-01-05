package fieldValidator

type NumericFieldValidator struct  {
}

func NewNumericFieldValidator() *NumericFieldValidator {
	return &NumericFieldValidator{}
}

func (nfv *NumericFieldValidator) Description() string {
	return "n"
}

func (nfv *NumericFieldValidator) IsValid(value string) bool {

	for _, c := range value {
		if c < 48 {
			return false
		}
		if c > 57 {
			return false
		}
	}

	return true
}