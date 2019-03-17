package attributes

import (
	"github.com/fohlen/entity-system/internal/pkg/types"
	"github.com/google/uuid"
	"reflect"
)

// An attribute has a name and a type
type Attribute struct {
	types.TypeImpl
	types.UUIDBase
	name string
}

// TODO: Refactor with injector
func NewAttribute(name string, attributeType reflect.Kind) Attribute {
	attribute := Attribute{}
	attribute.name = name
	attribute.SetValueType(attributeType)
	return attribute
}

// An instance is of type attribute and has a value
type Instance struct {
	types.InstanceImpl
	types.UUIDBase
	instanceType Attribute
}

// TODO: Refactor with injector
func NewInstance(attribute *Attribute, value interface{}) Instance {
	instance := Instance{}
	instance.SetInstanceType(attribute.TypeImpl)
	instance.SetValue(value)
	return instance
}

// Implements a singleton based manager
type Manager struct {
	attributes []*Attribute
	instances []*Instance
}
func (manager *Manager) Types() []uuid.UUID {
	var uuidList = make([]uuid.UUID, len(manager.attributes))
	for i, a := range manager.attributes {
		uuidList[i] = a.UUID()
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
func (manager *Manager) AddAttribute(attribute *Attribute) {
	manager.attributes = append(manager.attributes, attribute)
}

// TODO: This could be made variadic
// TODO: Speed can be improved via hash map
func (manager *Manager) AddInstance(instance *Instance) {
	for _, attributeUUID := range manager.Types() {
		if attributeUUID == instance.UUID() {
			manager.instances = append(manager.instances, instance)
		}
	}
}
