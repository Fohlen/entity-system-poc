package attributes

import (
	"errors"
	"github.com/fohlen/entity-system/internal/pkg/types"
	"github.com/google/uuid"
	"reflect"
	"sync"
)

// An attribute has a name and a type
type Attribute struct {
	uuid uuid.UUID
	attributeType reflect.Kind
	name string
	Name string
	types.Type
}

func (attribute *Attribute) UUID() uuid.UUID {
	if attribute.uuid == uuid.Nil {
		attribute.uuid = uuid.New()
	}
	return attribute.uuid
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
	attribute Attribute
	uuid uuid.UUID
	types.Instance
}

func (instance *Instance) UUID() uuid.UUID {
	if instance.uuid == uuid.Nil {
		instance.uuid = uuid.New()
	}
	return instance.uuid
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

func NewInstance(attribute Attribute, value interface{}) Instance {
	var instance = Instance{ attribute: attribute }
	instance.SetValue(value)
	return instance
}

type Manager struct {
	attributes []Attribute
	instances []Instance
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

