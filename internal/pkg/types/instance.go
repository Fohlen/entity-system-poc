package types

import "reflect"

type Instance interface {
	InstanceType() *Type
	SetInstanceType(t Type)
	Value() interface{}
	SetValue(value interface{})
}

type InstanceImpl struct {
	UUIDBase
	Instance
	value interface{}
	instanceType *TypeImpl
}

// Getter for the instance type
func (instance *InstanceImpl) InstanceType() Type {
	return instance.instanceType
}

func (instance *InstanceImpl) SetInstanceType(t *TypeImpl) {
	instance.instanceType = t
}

// Getter for the value
func (instance *InstanceImpl) Value() interface{} {
	return instance.value
}

// Setter for the value
func (instance *InstanceImpl) SetValue(value interface{}) {
	if reflect.ValueOf(value).Kind() != instance.instanceType.ValueType() {
		panic("invalid type used when setting value")
	}
	instance.value = value
}
