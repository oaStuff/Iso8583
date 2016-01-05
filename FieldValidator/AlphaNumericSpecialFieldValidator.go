package fieldValidator

type AlphaNumericSpecialFieldValidator struct {

}

func NewAlphaNumericSpecialFieldValidator() *AlphaNumericSpecialFieldValidator {
	return &AlphaNumericSpecialFieldValidator{}
}

func (aspv *AlphaNumericSpecialFieldValidator) Description() string {
	return "ans"
}

func (anpv *AlphaNumericSpecialFieldValidator) IsValid(value string) bool {
	for _, c:= range value {
		if c < 32 {
			return false
		}
	}

	return true
}