package types

import (
	"github.com/google/uuid"
)

type Manager interface {
	Types() []uuid.UUID
	Instances() []uuid.UUID
}

type ManagerImpl struct {
	Manager
	types []*TypeImpl
	instances []*InstanceImpl
}

func (manager *ManagerImpl) Types() []uuid.UUID {
	var uuidList = make([]uuid.UUID, len(manager.types))
	for i, t := range manager.types {
		uuidList[i] = t.UUID()
	}
	return uuidList
}

func (manager *ManagerImpl) Instances() []uuid.UUID {
	var uuidList = make([]uuid.UUID, len(manager.instances))
	for i, instance := range manager.instances {
		uuidList[i] = instance.UUID()
	}
	return uuidList
}
