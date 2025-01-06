package main

import (
	"fmt"

	"github.com/yashisrani/Go-Gin-API/api"
	"github.com/yashisrani/Go-Gin-API/controller"
)

func main() {
	api := api.ApiRoutes{}
	api.StartApp(controller.Server{})
	fmt.Printf("main server = %v\n", api)
}