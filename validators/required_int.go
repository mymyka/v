package validators

import (
	"errors"

	"github.com/mymyka/v/docs"
)

type reqiredIntRule struct{}

func RequiredInt() *reqiredIntRule {
	return &reqiredIntRule{}
}

func (r *reqiredIntRule) Validate(n *int) error {
	if *n == 0 {
		return errors.New("is required")
	}

	return nil
}

func (r *reqiredIntRule) Docs(p *docs.DocsProperty) {
	p.Required = true
}
