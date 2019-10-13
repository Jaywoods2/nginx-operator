package controller

import (
	"github.com/Jaywoods/nginx-operator/pkg/controller/nginxservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nginxservice.Add)
}
