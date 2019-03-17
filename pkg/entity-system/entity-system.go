package entity_system

import (
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"github.com/fohlen/entity-system/internal/pkg/entities"
)

type EntitySystem struct {
	AttributeManager attributes.Manager
	EntityManager entities.Manager
}

func NewEntitySystem() EntitySystem {
	return EntitySystem{}
}
