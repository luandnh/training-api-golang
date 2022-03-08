package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	e := gin.New()
	server := &Server{}
	e.Use(CORSMiddleware())
	e.Use(server.allowOptionsMethod())
	server.Engine = e
	return server
}

func (s *Server) allowOptionsMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
		c.Next()
	}
}

func (s *Server) Run() {
	s.Engine.Run()
}
