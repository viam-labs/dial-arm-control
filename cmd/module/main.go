package main

import (
	"dialarmcontrol"
	arm "go.viam.com/rdk/components/arm"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
)

func main() {
	// ModularMain can take multiple APIModel arguments, if your module implements multiple models.
	module.ModularMain(resource.APIModel{arm.API, dialarmcontrol.DialArmControl})
}
