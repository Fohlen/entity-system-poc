package entities

import (
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"github.com/fohlen/entity-system/internal/pkg/types"
	"github.com/google/uuid"
	"sync"
)

type Entity struct {
	uuid uuid.UUID
	attributes []*attributes.Attribute
	types.UUIDBase
}

func (entity *Entity) UUID() uuid.UUID {
	if entity.uuid == uuid.Nil {
		entity.uuid = uuid.New()
	}
	return entity.uuid
}

// Return a new entity type
func NewEntity(attributes ...*attributes.Attribute) Entity {
	return Entity{ attributes: attributes }
}

type Instance struct {
	uuid uuid.UUID
	entity Entity
	types.UUIDBase
}

func (instance *Instance) UUID() uuid.UUID {
	if instance.uuid == uuid.Nil {
		instance.uuid = uuid.New()
	}
	return instance.uuid
}

type Manager struct {
	entities []*Entity
	instances []*Instance
}

func (manager *Manager) Types() []uuid.UUID {
	var uuidList = make([]uuid.UUID, len(manager.entities))
	for i, e := range manager.entities {
		uuidList[i] = e.UUID()
	}
	return uuidList
}

func (manager *Manager) Instances()  []uuid.UUID {
	var uuidList = make([]uuid.UUID, len(manager.instances))
	for i, instance := range manager.instances {
		uuidList[i] = instance.UUID()
	}
	return uuidList
}

// TODO: This could be made variadic
func (manager *Manager) AddEntity(entity *Entity) {
	manager.entities = append(manager.entities, entity)
}

// TODO: This could be made variadic
// TODO: Speed can be improved via hash map
func (manager *Manager) AddInstance(instance *Instance) {
	for _, attributeUUID := range manager.Types() {
		if attributeUUID == instance.entity.UUID() {
			manager.instances = append(manager.instances, instance)
		}
	}
}

// Singleton pattern
var manager *Manager
var once sync.Once

func GetEntityInstanceManager() *Manager {
	once.Do(func() {
		manager = &Manager{}
	})
	return manager
}
