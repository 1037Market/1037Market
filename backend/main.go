package main

import (
	"1037Market/server"
	"1037Market/server/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	srv := server.DefaultServer()
	srv.Bind(":7301")
	srv.Cors()
	srv.RegisterMiddleWare(middleware.UserCookieCheck())
	srv.RegisterMiddleWare(middleware.RequestValidationCheck())
	srv.Route()
	srv.ListenAndServe()
}
