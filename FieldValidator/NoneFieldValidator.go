package fieldValidator


type NoneFieldValidator struct  {

}

func NewNoneFieldValidator() *NoneFieldValidator {
	return &NoneFieldValidator{}
}

func (nfv *NoneFieldValidator) Description() string {
	return "none"
}

func (nfv *NoneFieldValidator) IsValid(value string) bool {
	return true
}