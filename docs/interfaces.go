package docs

type ValidationRule[T any] interface {
	Validate(*T) error
	Docs(*DocsProperty)
}

type ValidationBuilder interface {
	ForString(*string, ...ValidationRule[string])
	ForInt(*int, ...ValidationRule[int])
}

type Validatable interface {
	Validator(b ValidationBuilder)
}
