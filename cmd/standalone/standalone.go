package main

import (
	"fmt"
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"github.com/fohlen/entity-system/internal/pkg/entities"
	"github.com/fohlen/entity-system/pkg/entity-system"
	"reflect"
)

func main()  {
	var system entity_system.EntitySystem
	var attribute = attributes.NewAttribute("abc", reflect.String)
	var instance = attributes.NewInstance(&attribute, "abc")
	system.AttributeManager.AddAttribute(&attribute)
	system.AttributeManager.AddInstance(&instance)

	var entity = entities.NewEntity(&attribute)
	var entityInstance = entities.NewInstance(&entity)
	system.EntityManager.AddEntity(&entity)
	system.EntityManager.AddInstance(&entityInstance)

	for _, attributeUUID := range system.EntityManager.Types() {
		fmt.Printf("Entity type UUID: %v\n", attributeUUID)
	}

	for _, instanceUUID := range system.EntityManager.Instances() {
		fmt.Printf("Entity instance UUID: %v\n", instanceUUID)
	}
}
