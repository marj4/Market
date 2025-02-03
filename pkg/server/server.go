package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Router *gin.Engine
}

// Start server
func StartServer(DB *sql.DB, DB2 *redis.Client) *Server {
	router := LoadRouter(DB, DB2)
	server := &Server{
		Router: router,
	}
	return server
}
