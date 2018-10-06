package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dyweb/gommon/errors"

	"github.com/at15/goyourcassandra/pkg/types"
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

// TODO: might pass query in get url as well, this makes life easier when copy and paste
func (srv *Server) handleQuery(res http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	var q types.QueryRequest
	if err := json.Unmarshal(b, &q); err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	qRes, err := srv.hosts.Query(q.Host, q.Keyspace, q.Query)
	if err != nil {
		write500(res, err)
		return
	}
	writeJSON(res, *qRes)
}

// http://localhost:8088/api/keyspace?host=localhost&keyspace=system
func (srv *Server) handleKeyspace(res http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	host := q.Get("host")
	if host == "" {
		write400(res, errors.New("host is required in query parameter"))
		return
	}
	ks := q.Get("keyspace")
	if ks == "" {
		write400(res, errors.New("keyspace is required in query parameter"))
		return
	}

	sess, err := srv.hosts.get(host).get(ks)
	if err != nil {
		write500(res, err)
		return
	}
	// TODO: it is returning a table called IndexInfo which seems does not exists ....
	meta, err := sess.KeyspaceMetadata(ks)
	if err != nil {
		write500(res, errors.Wrap(err, "error get keyspace meta using gocql"))
		return
	}
	// TODO: maybe allow use the unconverted version for debugging?
	converted := types.CopyKeyspaceMetadata(meta)
	writeJSON(res, converted)
}
