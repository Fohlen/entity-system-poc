package visual_scripting

import (
	"github.com/fohlen/entity-system/internal/pkg/attributes"
)

type InputConnector struct {
	channel chan interface{}
	attributeLeft *attributes.Instance
	attributeRight *attributes.Instance
}

func NewInputConnector(attributeLeft *attributes.Instance, attributeRight *attributes.Instance) InputConnector {
	channel := make(chan interface{})
	return InputConnector{
		channel: channel,
		attributeLeft: attributeLeft,
		attributeRight: attributeRight,
	}
}

// TODO: This can actually be executed inside the scope of attribute.Instance.SetValue() to make it reactive
func (connector *InputConnector) SetValueLeft(v interface{}) {
	connector.attributeLeft.SetValue(v)
	connector.channel <- connector.attributeLeft.Value()
	connector.attributeRight.SetValue(<-connector.channel)
}
