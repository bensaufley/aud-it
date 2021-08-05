package genre

import "github.com/bensaufley/aud-it/internal/entities/base"

type Genre struct {
	*base.Entity
	Name string
}
