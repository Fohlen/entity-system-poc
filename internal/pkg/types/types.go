package types

import (
	"github.com/google/uuid"
	"reflect"
)

// UUID interface. Every type that composes of UUIDBase should have a unique UUID
type UUIDBase interface {
	UUID() uuid.UUID
}

type Type interface {
	UUIDBase
	ValueType() reflect.Kind
}

type Instance interface {
	UUIDBase
	InstanceType() Type
	Value() interface{}
}

type Manager interface {
	Types() []Type
	Instances() []Instance
}
