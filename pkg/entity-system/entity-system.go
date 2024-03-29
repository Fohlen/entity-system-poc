package entity_system

import (
	"fmt"
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"github.com/fohlen/entity-system/internal/pkg/entities"
	"github.com/fohlen/entity-system/pkg/visual-scripting"
	"reflect"
	"time"
)

type EntitySystem struct {
	AttributeManager attributes.Manager
	EntityManager entities.Manager
}

func NewEntitySystem() EntitySystem {
	return EntitySystem{}
}

func (entitySystem *EntitySystem) Init() {
	var attribute = attributes.NewAttribute("abc", reflect.String)
	var instance = attributes.NewInstance(&attribute, "abc")
	var instance2 = attributes.NewInstance(&attribute, "abcd")
	entitySystem.AttributeManager.AddAttribute(&attribute)
	entitySystem.AttributeManager.AddInstance(&instance)

	var entity = entities.NewEntity(&attribute)
	var entityInstance = entities.NewInstance(&entity)
	var entityInstance2 = entities.NewInstance(&entity)
	entitySystem.EntityManager.AddEntity(&entity)
	entitySystem.EntityManager.AddInstance(&entityInstance)
	entitySystem.EntityManager.AddInstance(&entityInstance2)

	for _, attributeUUID := range entitySystem.EntityManager.Types() {
		fmt.Printf("Entity type UUID: %v\n", attributeUUID)
	}

	for _, instanceUUID := range entitySystem.EntityManager.Instances() {
		fmt.Printf("Entity instance UUID: %v\n", instanceUUID)
	}

	var connector = visual_scripting.NewInputConnector(&instance, &instance2)
	go connector.SetValueLeft("de") // concurrently updates both values
}

func (entitySystem *EntitySystem) Run() {
	var startTime = time.Now()
	for {
		time.Sleep(2 * time.Second)
		fmt.Printf("Uptime: %v\n", time.Now().Sub(startTime))
	}
}
