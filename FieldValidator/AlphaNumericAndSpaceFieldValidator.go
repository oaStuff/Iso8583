package fieldValidator

type AlphaNumericAndSpaceFieldValidator struct  {

}

func NewAlphaNumericAndSpaceFieldValidator() *AlphaNumericAndSpaceFieldValidator {
	return &AlphaNumericAndSpaceFieldValidator{}
}

func (ansv *AlphaNumericAndSpaceFieldValidator) Description() string {
	return "ansp"
}

func (ansv *AlphaNumericAndSpaceFieldValidator) IsValid(value string) bool {

	for _, c := range value {
		if (c < 32 || c > 32) && (c < 48 || c > 57) && (c < 65 || c > 90) && (c < 97 || c > 122) {
			return false
		}
	}

	return true
}