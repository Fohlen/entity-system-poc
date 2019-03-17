package entities

import (
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"github.com/fohlen/entity-system/internal/pkg/types"
	"github.com/google/uuid"
)

type Entity struct {
	attributes []*attributes.Attribute
	types.UUIDBase
}

// Return a new entity type
func NewEntity(attributes ...*attributes.Attribute) Entity {
	return Entity{ attributes: attributes }
}

type Instance struct {
	entity *Entity
	types.UUIDBase
}

func NewInstance(entity *Entity) Instance {
	return Instance{entity: entity}
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
