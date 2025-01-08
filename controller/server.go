package controller

import (
	"fmt"

	"github.com/yashisrani/Go-Gin-API/store"
	"github.com/yashisrani/Go-Gin-API/utils"
)

type Server struct {
	PostgressDB store.StoreOperation
}

func (s *Server) NewServer(pgdb store.Postgress) {
	utils.SetLogger()
	utils.Logger.Infof("logger setup done")
	s.PostgressDB = &pgdb
	err:= s.PostgressDB.NewStore()
	if err != nil {
		utils.Logger.Errorf("error while creating Postgress connection")
	}else {
		utils.Logger.Infof("connection established")
	}
	fmt.Printf("server = %v\n", s)
}

type ServerOperation interface {
	NewServer(pgdb store.Postgress)
}
