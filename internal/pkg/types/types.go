package types

import (
	"sync"
	"github.com/fohlen/entity-system/internal/pkg/container"
	"github.com/google/uuid"
)

type Type struct {
	uuid uuid.UUID
	container container.Container
}

type TypeInstance struct {
	uuid.UUID
	types []Type
}

func (instance *TypeInstance) AddTypeInstance(t Type) {
	instance.types = append(instance.types, t)
}

func (instance *TypeInstance) ListTypeInstances() []uuid.UUID {
	UUIDList := make([]uuid.UUID, len(instance.types))
	for i, t := range instance.types {
		UUIDList[i] = t.uuid
	}
	return UUIDList
}

func (instance *TypeInstance) DeleteTypeInstance(uuid uuid.UUID) {
	for i, _ := range instance.types {
		if instance.types[i].uuid == uuid {
			instance.types = append(instance.types[:i], instance.types[i+1:]...)
			break
		}
	}
}

type TypeInstanceManager struct {
	instance TypeInstance
}

// Singleton pattern
var manager *TypeInstanceManager
var once sync.Once

func GetTypeInstanceManager() *TypeInstanceManager {
	once.Do(func() {
		manager = &TypeInstanceManager{}
	})
	return manager
}
