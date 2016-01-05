package fieldValidator

type HexFieldValidator struct  {
}

func NewHexFieldValidator() *HexFieldValidator {
	return &HexFieldValidator{}
}

func (hfv *HexFieldValidator) Description() string {
	return "Hex"
}

func (hfv *HexFieldValidator) IsValid(value string) bool {

	for _, c := range value {
		if (c < 48 || c > 57) && (c < 65 || c > 70) && (c < 97 || c > 102) {
			return false
		}
	}

	return true
}