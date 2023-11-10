package server

import (
	"1037Market/server/api"
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
	s.POST("/api/user/register/email", api.RegisterEmail())
	s.POST("/api/user/register", api.Register())
	s.POST("/api/user/login", api.Login())
	s.POST("/api/user/info", api.UpdateUserInfo())
	s.GET("/api/user/info", api.GetUserInfo())
	s.POST("/api/image", api.UploadImage())
	s.GET("/api/image", api.DownloadImage())
	s.POST("/api/product", api.PublishProduct())
	s.POST("/api/update", api.UpdateProduct())
	s.POST("/api/comment", api.CreateComment())
	s.GET("/api/comment", api.QueryCommentList())
	s.GET("/api/comment/get", api.GetCommentById())
	s.GET("/api/product/get", api.GetProductById())
	s.GET("/api/product/query", api.GetProductListByKeyword())
	s.GET("api/product/student", api.GetProductListByStudentId())
	s.POST("/api/product/sold", api.SoldProduct())
	s.DELETE("/api/product", api.DeleteProduct())
	s.GET("/api/product/recommend", api.GetRecommendProductList())
	s.GET("/api/product/category", api.GetProductListByCategory())
	s.GET("/api/product/categories", api.GetCategoryList())
	s.GET("/api/subscribe", api.GetSubscribes())
	s.POST("/api/subscribe", api.AddSubscribe())
	s.DELETE("/api/subscribe", api.DeleteSubscribe())
	s.GET("/api/chat/session", api.GetSingleSessIdByStuIds())
	s.GET("/api/chat/sessions", api.GetSessIdListBySingleStuId())
	s.GET("/api/chat/userInfos", api.GetTwoStuInfosBySessId())
	s.GET("/api/chat/message", api.GetNewestMsgIdBySessId())
	s.GET("/api/chat/messages", api.GetNMsgIdsFromKthLastBySessId())
	s.GET("/api/chat/content", api.GetMsgInfoByMsgId())
	s.POST("/api/chat/send", api.SendMsg())
	s.GET("/api/chat/session/messages", api.GetMsgsInSession())
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
