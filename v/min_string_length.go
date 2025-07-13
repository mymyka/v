package v

import (
	"fmt"

	"github.com/mymyka/v/v/docs"
)

type minLengthRule struct {
	minLength int
}

func MinLength(l int) *minLengthRule {
	return &minLengthRule{
		minLength: l,
	}
}

func (r *minLengthRule) Validate(s *string) error {
	if len(*s) < r.minLength {
		return fmt.Errorf("min lenght should be %d", r.minLength)
	}

	return nil
}

func (r *minLengthRule) Docs(p *docs.Property) {
	p.MinLength = r.minLength
}
