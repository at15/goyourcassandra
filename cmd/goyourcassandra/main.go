package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gocql/gocql"
	dlog "github.com/dyweb/gommon/log"
)

var log = dlog.NewApplicationLogger()

type connPool struct {
	mu sync.Mutex

	conns map[string]*gocql.Session
}

func main() {
	fmt.Println("start http server")
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("pong"))
	})
	mux.HandleFunc("/api/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("api pong"))
	})
	logged := http.HandlerFunc(func(res http.ResponseWriter, r *http.Request) {
		tw := &TrackedWriter{w: res, status: 200}
		mux.ServeHTTP(tw, r)
		log.InfoF("http", dlog.Fields{
			dlog.Str("remote", r.RemoteAddr),
			dlog.Str("proto", r.Proto),
			dlog.Str("url", r.URL.String()),
			dlog.Int("size", tw.Size()),
			dlog.Int("status", tw.Status()),
		})
	})
	http.ListenAndServe(":8080", logged)
}
