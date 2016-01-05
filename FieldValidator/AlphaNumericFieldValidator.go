package fieldValidator

type AlphaNumericFieldValidator struct {

}

func NewAlphaNumericFieldValidator() *AlphaNumericFieldValidator {
	return &AlphaNumericFieldValidator{}
}

func (anv *AlphaNumericFieldValidator) Description() string {
	return "an"
}

func (anv *AlphaNumericFieldValidator) IsValid(value string) bool {
	for _, c:= range value {
		if (c < 48 || c > 57) && (c < 65 || c > 90) && (c < 97 || c > 122) {
			return false;
		}
	}

	return true
}