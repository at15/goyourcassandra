package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dyweb/gommon/errors"

	"github.com/at15/goyourcassandra/pkg/types"
)

type Server struct {
	cfg Config

	hosts *hostPool
	mux   *http.ServeMux
}

func New(cfg Config) (*Server, error) {
	srv := Server{
		cfg:   cfg,
		hosts: newHostPool(),
	}
	srv.routes()
	return &srv, nil
}

// TODO: might pass query in GET url as well, this makes life easier when copy and paste
func (srv *Server) handleQuery(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var q types.QueryRequest
	if err := json.Unmarshal(b, &q); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	qRes, err := srv.hosts.Query(q.Host, q.Keyspace, q.Query)
	if err != nil {
		write500(w, err)
		return
	}
	writeJSON(w, *qRes)
}

// http://localhost:8088/api/keyspace?host=localhost&keyspace=system
func (srv *Server) handleKeyspace(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	host := q.Get("host")
	if host == "" {
		write400(w, errors.New("host is required in query parameter"))
		return
	}
	ks := q.Get("keyspace")
	if ks == "" {
		write400(w, errors.New("keyspace is required in query parameter"))
		return
	}

	sess, err := srv.hosts.get(host).get(ks)
	if err != nil {
		write500(w, err)
		return
	}
	// TODO: it is returning a table called IndexInfo which seems does not exists ....
	meta, err := sess.KeyspaceMetadata(ks)
	if err != nil {
		write500(w, errors.Wrap(err, "error get keyspace meta using gocql"))
		return
	}
	// TODO: maybe allow use the unconverted version for debugging?
	converted := types.CopyKeyspaceMetadata(meta)
	writeJSON(w, converted)
}
