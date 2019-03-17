package types

import (
	"github.com/google/uuid"
	"sync"
)

type DataType int

const (
	BOOL DataType = iota
	INT DataType = iota
	STRING DataType = iota
)

type Type struct {
	UUID uuid.UUID
	ValueType DataType
}

type Instance struct {
	UUID uuid.UUID
	InstanceType Type
	Value interface{}
}

type Manager struct {
	types []Type
	instances []Instance
}

func (manager *Manager) AddType(types ...Type) {
	for _, t := range types {
		manager.types = append(manager.types, t)
	}
}

func (manager *Manager) ListTypes() []uuid.UUID {
	var uuidList = make([]uuid.UUID, len(manager.types))
	for i, t := range manager.types {
		uuidList[i] = t.UUID
	}
	return uuidList
}

func (manager *Manager) AddInstance(instances ...Instance) {
	for _, i := range instances {
		manager.instances = append(manager.instances, i)
	}
}

func (manager *Manager) ListInstances() []uuid.UUID {
	var uuidList = make([]uuid.UUID, len(manager.types))
	for i, instance := range manager.instances {
		uuidList[i] = instance.UUID
	}
	return uuidList
}

// Singleton pattern
var manager *Manager
var once sync.Once

func GetManager() *Manager {
	once.Do(func() {
		manager = &Manager{}
	})
	return manager
}
