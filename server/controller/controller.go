package controller

import (
	"github.com/gorilla/mux"
	"github.com/5imili/reboot/server/service"
	"github.com/5imili/reboot/pkg/dao"
)

type Controller interface{
	Register(router *mux.Router)
}

type Options struct{
	Service service.Operation
	DB	dao.Storage
}
