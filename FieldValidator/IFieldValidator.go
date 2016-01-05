package fieldValidator

type IFieldValidator interface {
	Description() string
	IsValid(string) bool
}