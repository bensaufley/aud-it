package list

import "github.com/bensaufley/aud-it/internal/entities/base"

type List struct {
	*base.Entity
	Name string
}
