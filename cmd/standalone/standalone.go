package main

import (
	"fmt"
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"github.com/fohlen/entity-system/pkg/entity-system"
	"reflect"
)

func main()  {
	var system entity_system.EntitySystem
	var attribute = attributes.NewAttribute("abc", reflect.String)
	var instance = attributes.NewInstance(&attribute, "abc")
	system.AttributeManager.AddAttribute(&attribute)
	system.AttributeManager.AddInstance(&instance)

	for _, attributeUUID := range system.AttributeManager.Types() {
		fmt.Printf("Attribute UUID: %v\n", attributeUUID)
	}

	for _, instanceUUID := range system.AttributeManager.Instances() {
		fmt.Printf("Instance UUID: %v\n", instanceUUID)
	}
}
