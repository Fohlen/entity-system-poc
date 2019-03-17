package types

import "reflect"

type Type interface {
	ValueType() reflect.Kind
	SetValueType(kind reflect.Kind)
}

type TypeImpl struct {
	UUIDBase
	Type
	valueType reflect.Kind
}

// Getter for the ValueType
func (t *TypeImpl) ValueType() reflect.Kind  {
	return t.valueType
}

// Setter for the ValueType
// Never invoke this on an already initialized type as it will panic
func (t *TypeImpl) SetValueType(kind reflect.Kind) {
	if t.valueType != reflect.Invalid {
		panic("Cannot change existing type ")
	}
	t.valueType = kind
}
