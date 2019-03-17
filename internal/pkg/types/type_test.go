package types

import (
	"reflect"
	"testing"
)

func TestTypeImpl_SetValueType(t *testing.T) {
	typeImpl := TypeImpl{}
	typeImpl.SetValueType(reflect.String)
	if typeImpl.valueType != reflect.String {
		t.Error("Incorrect value type")
	}
}

func TestTypeImpl_ValueType(t *testing.T) {
	typeImpl := TypeImpl{}
	if typeImpl.valueType != reflect.Invalid {
		t.Error("By default the value should be invalid")
	}
}