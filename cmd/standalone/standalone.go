package main

import (
	"github.com/fohlen/entity-system/pkg/entity-system"
)

func main()  {
	var system = entity_system.NewEntitySystem()
	system.Init()
	system.Run()
}
