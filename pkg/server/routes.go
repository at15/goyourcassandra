package server

import (
	"net/http"
)

func (srv *Server) routes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("pong"))
	})
	mux.HandleFunc("/api/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("api pong"))
	})
	mux.HandleFunc("/api/query", srv.handleQuery)
	srv.mux = mux
}
