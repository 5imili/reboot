package controller

import (
	"github.com/gorilla/mux"
	"github.com/5imili/reboot/server/service"
)

type Controller interface{
	Register(router *mux.Router)
}

type Options struct{
	Service service.Operation
	//DB
}
