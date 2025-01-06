package api

import (
	"github.com/yashisrani/Go-Gin-API/controller"
	"github.com/yashisrani/Go-Gin-API/store"
)

type ApiRoutes struct {
	Server controller.ServerOperation
}

func (api *ApiRoutes) StartApp(server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})
}
