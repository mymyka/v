package builder

import "github.com/mymyka/v/v/docs"

type ValidationRule[T any] interface {
	Validate(*T) error
	Docs(*docs.Property)
}

type ValidationBuilder interface {
	ForString(*string, ...ValidationRule[string])
	ForInt(*int, ...ValidationRule[int])
}
