package types

import (
	"github.com/google/uuid"
	"testing"
)

func TestUUIDBase_UUID(t *testing.T) {
	id := UUIDBase{}
	if id.UUID() == uuid.Nil {
		t.Error("Expected UUID to be not empty")
	}
}
