package types

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInstanceImpl_SetInstanceType(t *testing.T) {
	instanceImpl := InstanceImpl{}
	typeImpl := TypeImpl{}
	typeImpl.SetValueType(reflect.Int)
	instanceImpl.SetInstanceType(&typeImpl)
}

func TestInstanceImpl_SetValue(t *testing.T) {
	instanceImpl := InstanceImpl{}
	typeImpl := TypeImpl{}
	typeImpl.SetValueType(reflect.Int)
	instanceImpl.SetInstanceType(&typeImpl)
	instanceImpl.SetValue(3)
	assert.Panics(t, func() {
		instanceImpl.SetValue("a")
	})
}

func TestInstanceImpl_Value(t *testing.T) {
	instanceImpl := InstanceImpl{}
	typeImpl := TypeImpl{}
	typeImpl.SetValueType(reflect.Int)
	instanceImpl.SetInstanceType(&typeImpl)
	instanceImpl.SetValue(3)
	if instanceImpl.Value() != 3 {
		t.Error("Setting the value incorrectly")
	}
}
