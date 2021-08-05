package user

import (
	"github.com/bensaufley/aud-it/internal/entities/base"
)

type User struct {
	*base.Entity
	Username string
}
