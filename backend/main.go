package main

import (
	"1037Market/server"
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
	srv.Route()
	srv.ListenAndServe()
}
