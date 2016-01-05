package lengthValidators

type ILengthValidator interface {
	IsValid(string) bool
}