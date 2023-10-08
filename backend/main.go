package main

import (
	"1037Market/server"
)

func main() {
	srv := server.DefaultServer()
	srv.Bind(":8080")
	srv.Cors()
	srv.Route()
	srv.ListenAndServe()
}
