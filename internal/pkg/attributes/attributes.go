package attributes

import (
	"github.com/fohlen/entity-system/internal/pkg/types"
	"sync"
)

type Attribute struct {
	types.Type
}

type AttributeInstance struct {
	types.TypeInstance
}

type AttributeInstanceManager struct {
	instance AttributeInstance
}

// Singleton pattern
var manager *AttributeInstanceManager
var once sync.Once

func GetAttributeInstanceManager() *AttributeInstanceManager {
	once.Do(func() {
		manager = &AttributeInstanceManager{}
	})
	return manager
}
