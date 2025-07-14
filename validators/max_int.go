package validators

import (
	"fmt"

	"github.com/mymyka/v/docs"
)

type maxIntRule struct {
	max int
}

func MaxInt(max int) *maxIntRule {
	return &maxIntRule{
		max: max,
	}
}

func (r *maxIntRule) Validate(n *int) error {
	if *n > r.max {
		return fmt.Errorf("mast be not more than %d", r.max)
	}

	return nil
}

func (r *maxIntRule) Docs(p *docs.DocsProperty) {
	p.MaxValue = r.max
}
