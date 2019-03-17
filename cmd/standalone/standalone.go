package main

import (
	"fmt"
	"github.com/fohlen/entity-system/internal/pkg/types"
	"github.com/google/uuid"
)

func main()  {
	var t = types.Type{ UUID: uuid.New(), ValueType: types.BOOL }
	var t2 = types.Type{ UUID: uuid.New(), ValueType: types.STRING }

	var i = types.Instance{ UUID: uuid.New(), InstanceType: t, Value: false }

	var manager = types.GetManager()
	manager.AddType(t, t2)
	manager.AddInstance(i)

	for _, v := range manager.ListTypes() {
		fmt.Printf("Type UUID: %v\n", v)
	}

	for _, v := range manager.ListInstances() {
		fmt.Printf("Type instance UUID: %v\n", v)
	}
}