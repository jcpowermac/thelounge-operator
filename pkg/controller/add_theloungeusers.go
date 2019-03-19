package controller

import (
	"github.com/jcpowermac/thelounge-operator/pkg/controller/theloungeusers"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, theloungeusers.Add)
}
