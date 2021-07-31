package ulid

import (
	"math/rand"
	"time"

	oklog "github.com/oklog/ulid/v2"
)

type ULIDGenerator struct {
	entropy *oklog.MonotonicEntropy
}

func (g *ULIDGenerator) New() oklog.ULID {
	return oklog.MustNew(oklog.Timestamp(time.Now()), g.entropy)
}

func (g *ULIDGenerator) String() string {
	return g.New().String()
}

func NewGenerator() *ULIDGenerator {
	entropy := oklog.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	return &ULIDGenerator{entropy: entropy}
}
