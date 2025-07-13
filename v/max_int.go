package v

import (
	"fmt"

	"github.com/mymyka/v/v/docs"
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

func (r *maxIntRule) Docs(p *docs.Property) {
	p.MaxValue = r.max
}
