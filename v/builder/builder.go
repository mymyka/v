package builder

import "fmt"

type ValidationBuilderImpl struct {
	stringFields []Field[string]
	intFields    []Field[int]
}

type Field[T any] struct {
	field *T
	rules []ValidationRule[T]
}

func (b *ValidationBuilderImpl) ForString(s *string, rules ...ValidationRule[string]) {
	b.stringFields = append(b.stringFields, Field[string]{
		field: s,
		rules: rules,
	})
}

func (b *ValidationBuilderImpl) ForInt(s *int, rules ...ValidationRule[int]) {
	b.intFields = append(b.intFields, Field[int]{
		field: s,
		rules: rules,
	})
}

type ValidationError map[string]error

func (v ValidationError) Error() string {
	if len(v) == 0 {
		return "validation error"
	}

	var result string
	for field, err := range v {
		if result != "" {
			result += "; "
		}
		result += field + ": " + err.Error()
	}
	return result
}

func (b *ValidationBuilderImpl) Validate() error {
	var validationErrors ValidationError

	for i, field := range b.stringFields {
		for _, rule := range field.rules {
			if err := rule.Validate(field.field); err != nil {
				if validationErrors == nil {
					validationErrors = make(ValidationError)
				}
				fieldName := fmt.Sprintf("string_field_%d", i)
				validationErrors[fieldName] = err
				break
			}
		}
	}

	for i, field := range b.intFields {
		for _, rule := range field.rules {
			if err := rule.Validate(field.field); err != nil {
				if validationErrors == nil {
					validationErrors = make(ValidationError)
				}
				fieldName := fmt.Sprintf("int_field_%d", i)
				validationErrors[fieldName] = err
				break
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}
	return nil
}
