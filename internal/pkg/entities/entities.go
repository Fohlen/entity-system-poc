package entities

import (
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"github.com/fohlen/entity-system/internal/pkg/types"
	"sync"
)

type Entity struct {
	attributes []attributes.Attribute
	types.Type
}

type EntityInstance struct {
	types.TypeInstance
}

// Singleton pattern
var manager *EntityInstance
var once sync.Once

func GetEntityInstanceManager() *EntityInstance {
	once.Do(func() {
		manager = &EntityInstance{}
	})
	return manager
}
