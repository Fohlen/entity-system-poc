package types

import (
	"github.com/google/uuid"
)

type Manager interface {
	Types() []uuid.UUID
	Instances() []uuid.UUID
}
