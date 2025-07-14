package validators

import (
	"fmt"

	"github.com/mymyka/v/docs"
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

func (r *minLengthRule) Docs(p *docs.DocsProperty) {
	p.MinLength = r.minLength
}
