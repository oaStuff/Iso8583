package fieldValidator

type AlphaNumericPrintableFieldValidator struct  {

}

func NewAlphaNumericPrintableFieldValidator() *AlphaNumericPrintableFieldValidator {
	return &AlphaNumericPrintableFieldValidator{}
}

func (anpv *AlphaNumericPrintableFieldValidator) Description() string {
	return "anp"
}

func (anpv *AlphaNumericPrintableFieldValidator) IsValid(value string) bool {
	for _, c:= range value {
		if c < 32 {
			return false
		}
	}

	return true
}