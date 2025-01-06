package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func StartServer(DB *sql.DB) *Server {
	router := LoadRouter(DB)
	server := &Server{
		Router: router,
	}

	return server
}
