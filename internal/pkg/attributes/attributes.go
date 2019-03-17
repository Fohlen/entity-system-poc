package attributes

import (
	"errors"
	"github.com/fohlen/entity-system/internal/pkg/types"
	"github.com/google/uuid"
	"reflect"
)

// An attribute has a name and a type
type Attribute struct {
	attributeType reflect.Kind
	name string
	types.Type
	types.UUIDBase
}

func (attribute *Attribute) ValueType() reflect.Kind {
	return attribute.attributeType
}

func NewAttribute(name string, attributeType reflect.Kind) Attribute {
	var attribute = Attribute{
		name: name,
		attributeType: attributeType,
	}
	return attribute
}

// An instance is of type attribute and has a value
type Instance struct {
	value interface{}
	attribute *Attribute
	types.Instance
	types.UUIDBase
}

// Set the value of an instance
func (instance *Instance) SetValue(i interface{}) {
	if reflect.ValueOf(i).Kind() != instance.attribute.attributeType {
		panic(errors.New("invalid type used"))
	}
	instance.value = i
}

func (instance *Instance) Value() interface{} {
	return instance.value
}

func NewInstance(attribute *Attribute, value interface{}) Instance {
	var instance = Instance{ attribute: attribute }
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
		if attributeUUID == instance.attribute.UUID() {
			manager.instances = append(manager.instances, instance)
		}
	}
}
