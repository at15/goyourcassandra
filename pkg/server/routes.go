package server

import (
	"fmt"
	"net/http"

	dlog "github.com/dyweb/gommon/log"
)

func (srv *Server) routes() {
	// TODO: need to serve static assets, can use noodle
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("pong"))
	})
	mux.HandleFunc("/api/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("api pong"))
	})
	mux.HandleFunc("/api/bookmark", srv.handleGetBookmarks)
	mux.HandleFunc("/api/query", srv.handleQuery)
	mux.HandleFunc("/api/keyspace", srv.handleKeyspace)
	srv.mux = mux
}

func (srv *Server) Handler() http.Handler {
	return srv.mux
}

func (srv *Server) HandlerWithLogger() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, r *http.Request) {
		tw := &TrackedWriter{w: res, status: 200}
		srv.mux.ServeHTTP(tw, r)
		log.InfoF(fmt.Sprintf("%d %s %s", tw.Status(), r.Method, r.URL.Path), dlog.Fields{
			dlog.Str("remote", r.RemoteAddr),
			dlog.Str("proto", r.Proto),
			dlog.Str("url", r.URL.String()),
			dlog.Int("size", tw.Size()),
			dlog.Int("status", tw.Status()),
		})
	})
}
