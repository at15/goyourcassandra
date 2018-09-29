package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/at15/goyourcassandra/pkg/types"
	dlog "github.com/dyweb/gommon/log"
)

type Server struct {
	hosts *hostPool
	mux   *http.ServeMux
}

func New() (*Server, error) {
	srv := Server{
		hosts: newHostPool(),
	}
	srv.routes()
	return &srv, nil
}

func (srv *Server) handleQuery(res http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	var q types.Request
	if err := json.Unmarshal(b, &q); err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	qRes, err := srv.hosts.Query(q.Host, q.Keyspace, q.Query)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	bRes, err := json.Marshal(*qRes)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	res.Write(bRes)
}

func (srv *Server) Handler() http.Handler {
	return srv.mux
}

func (srv *Server) HandlerWithLogger() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, r *http.Request) {
		tw := &TrackedWriter{w: res, status: 200}
		srv.mux.ServeHTTP(tw, r)
		log.InfoF("http", dlog.Fields{
			dlog.Str("remote", r.RemoteAddr),
			dlog.Str("proto", r.Proto),
			dlog.Str("url", r.URL.String()),
			dlog.Int("size", tw.Size()),
			dlog.Int("status", tw.Status()),
		})
	})
}
