package fieldValidator

type AlphaFieldValidator struct {

}

func NewAlphaFieldValidator() *AlphaFieldValidator {
	return &AlphaFieldValidator{}
}

func (afv *AlphaFieldValidator) Description() string {
	return "a"
}

func (afv * AlphaFieldValidator) IsValid(value string) bool {

	for _, c := range value {
		if (c < 65 || c > 90) && (c < 97 || c > 122) {
			return false
		}
	}

	return true
}