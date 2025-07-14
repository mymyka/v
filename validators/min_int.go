package validators

import (
	"fmt"

	"github.com/mymyka/v/docs"
)

type minIntRule struct {
	min int
}

func MinInt(min int) *minIntRule {
	return &minIntRule{
		min: min,
	}
}

func (r *minIntRule) Validate(n *int) error {
	if *n < r.min {
		return fmt.Errorf("must me more than %d", r.min)
	}

	return nil
}

func (r *minIntRule) Docs(p *docs.DocsProperty) {
	p.MaxValue = r.min
}
