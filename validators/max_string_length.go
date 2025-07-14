package validators

import (
	"fmt"

	"github.com/mymyka/v/docs"
)

type maxLengthRule struct {
	maxLength int
}

func MaxLength(l int) *maxLengthRule {
	return &maxLengthRule{
		maxLength: l,
	}
}

func (r *maxLengthRule) Validate(s *string) error {
	if len(*s) > r.maxLength {
		return fmt.Errorf("max lenght should be not more than %d", r.maxLength)
	}

	return nil
}

func (r *maxLengthRule) Docs(p *docs.DocsProperty) {
	p.MaxLength = r.maxLength
}
