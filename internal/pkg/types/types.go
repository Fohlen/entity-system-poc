package types

import (
	"github.com/google/uuid"
	"reflect"
)

type UUIDBase struct {
	uuid uuid.UUID
}

func (uuidBase *UUIDBase) UUID() uuid.UUID {
	if uuidBase.uuid == uuid.Nil {
		uuidBase.uuid = uuid.New()
	}
	return uuidBase.uuid
}

type Type interface {
	ValueType() reflect.Kind
}

type Instance interface {
	InstanceType() Type
	Value() interface{}
}

type Manager interface {
	Types() []uuid.UUID
	Instances() []uuid.UUID
}
