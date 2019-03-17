package types

import "github.com/google/uuid"

type UUIDBase struct {
	uuid uuid.UUID
}

func (uuidBase *UUIDBase) UUID() uuid.UUID {
	if uuidBase.uuid == uuid.Nil {
		uuidBase.uuid = uuid.New()
	}
	return uuidBase.uuid
}

