package types

import (
	"reflect"
	"testing"
)

func TestManagerImpl_Instances(t *testing.T) {

}

func TestManagerImpl_Types(t *testing.T) {
	typeInstance := TypeImpl{}
	typeInstance.SetValueType(reflect.String)
	managerImpl := ManagerImpl{}
	managerImpl.types = append(managerImpl.types, &typeInstance)
	if managerImpl.Types()[0] != typeInstance.UUID() {
		t.Error("Type UUID's are not identical")
	}
}
