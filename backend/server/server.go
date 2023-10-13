package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	router *gin.Engine
	srv    *http.Server
}

func DefaultServer() *Server {
	return &Server{
		router: gin.Default(),
		srv:    nil,
	}
}

func (s *Server) Bind(port string) {
	s.srv = &http.Server{
		Addr:    port,
		Handler: s.router,
	}
}

func (s *Server) Cors() {
	s.router.Use(cors())
}

func (s *Server) ListenAndServe() {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT)
	<-quit
	log.Println("Shutdown Server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s\n", err)
	}
	log.Println("Server Exiting")
}

func (s *Server) GET(relativePath string, handler gin.HandlerFunc) {
	s.router.GET(relativePath, handler)
}

func (s *Server) POST(relativePath string, handler gin.HandlerFunc) {
	s.router.POST(relativePath, handler)
}

func (s *Server) PUT(relativePath string, handler gin.HandlerFunc) {
	s.router.PUT(relativePath, handler)
}

func (s *Server) DELETE(relativePath string, handler gin.HandlerFunc) {
	s.router.DELETE(relativePath, handler)
}

func (s *Server) Route() {
	s.GET("/", helloWorld())
	s.POST("/api/user/register/email", registerEmail())
	s.POST("/api/user/register", register())
	s.POST("/api/user/login", login())
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}