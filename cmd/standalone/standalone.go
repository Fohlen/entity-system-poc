package main

import (
	"fmt"
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"reflect"
)

func main()  {
	var attribute = attributes.NewAttribute("abc", reflect.String)
	var instance = attributes.NewInstance(&attribute, "abc")
	var manager = attributes.GetManager()
	manager.AddAttribute(&attribute)
	manager.AddInstance(&instance)

	for _, attributeUUID := range manager.Types() {
		fmt.Printf("Attribute UUID: %v\n", attributeUUID)
	}

	for _, instanceUUID := range manager.Instances() {
		fmt.Printf("Instance UUID: %v\n", instanceUUID)
	}
}
