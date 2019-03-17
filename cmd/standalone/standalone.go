package main

import (
	"fmt"
	"github.com/fohlen/entity-system/internal/pkg/attributes"
	"reflect"
)

func main()  {
	var attribute = attributes.NewAttribute("abc", reflect.String)
	var instance = attributes.NewInstance(attribute, "abc")
	fmt.Println(instance.Value())
}
