package v

import (
	"errors"

	"github.com/mymyka/v/v/docs"
)

type reqiredStringRule struct{}

func RequiredString() *reqiredStringRule {
	return &reqiredStringRule{}
}

func (r *reqiredStringRule) Validate(s *string) error {
	if *s == "" {
		return errors.New("is required")
	}

	return nil
}

func (r *reqiredStringRule) Docs(p *docs.Property) {
	p.Required = true
}
