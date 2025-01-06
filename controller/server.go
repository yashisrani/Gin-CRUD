package controller

import (
	"fmt"

	"github.com/yashisrani/Go-Gin-API/store"
)

type Server struct {
	PostgressDB store.StoreOperation
}


func (s *Server) NewServer(pgdb store.Postgress)  {
	s.PostgressDB = &pgdb
	s.PostgressDB.NewStore()
	fmt.Printf("server = %v\n", s)
}

type ServerOperation interface {
	NewServer(pgdb store.Postgress)
}