package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/gocql/gocql"
	"github.com/spf13/cobra"
)

var log = dlog.NewApplicationLogger()

type Request struct {
	Host     string `json:"host"`
	Keyspace string `json:"keyspace"`
	Query    string `json:"query"`
}

type Result struct {
	Err     error
	Columns []gocql.ColumnInfo
	Rows    []map[string]interface{}
}

type hostPool struct {
	mu sync.RWMutex

	hosts map[string]*connPool
}

func newHostPool() *hostPool {
	return &hostPool{
		hosts: make(map[string]*connPool),
	}
}

func (pool *hostPool) get(host string) *connPool {
	pool.mu.RLock()
	h, ok := pool.hosts[host]
	pool.mu.RUnlock()
	if !ok {
		log.Infof("creating new connection pool for host %s", host)
		pool.mu.Lock()
		h = newConnPool(host)
		pool.hosts[host] = h
		pool.mu.Unlock()
	}
	return h
}

func (pool *hostPool) Query(host, keyspace, query string) (*Result, error) {
	p := pool.get(host)
	return p.Query(keyspace, query)
}

type connPool struct {
	mu sync.RWMutex

	host  string // TODO: might change to []string because gocql supports round robin
	conns map[string]*gocql.Session
}

func newConnPool(host string) *connPool {
	return &connPool{
		host:  host,
		conns: make(map[string]*gocql.Session),
	}
}

func (pool *connPool) get(keyspace string) (*gocql.Session, error) {
	pool.mu.RLock()
	session, ok := pool.conns[keyspace]
	pool.mu.RUnlock()
	if !ok {
		log.Infof("creating new session for host %s keyspace %s", pool.host, keyspace)
		cluster := gocql.NewCluster(pool.host)
		cluster.Keyspace = keyspace
		// TODO: either allow user to pass it ... or config it in backend
		cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
		tSession, err := cluster.CreateSession()
		if err != nil {
			log.Errorf("error creating new session for host %s keyspace %s: %s", pool.host, keyspace, err)
			return nil, errors.Wrapf(err, "error connect use keyspace %s", keyspace)
		}
		pool.mu.Lock()
		pool.conns[keyspace] = tSession
		pool.mu.Unlock()
		session = tSession
		log.Infof("created new session for host %s keyspace %s", pool.host, keyspace)
	}
	return session, nil
}

func (pool *connPool) Query(keyspace string, query string) (*Result, error) {
	session, err := pool.get(keyspace)
	if err != nil {
		return nil, err
	}
	iter := session.Query(query).Iter()
	//err := iter.Scanner().Err()
	//if err != nil {
	//	return errors.Wrap(err, "error scan")
	//}
	cols := iter.Columns()
	// TODO: need to convert TypeInfo
	//cols[0].TypeInfo.Type()
	rows := make([]map[string]interface{}, 0)
	for {
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			break
		}
		rows = append(rows, row)
	}
	err = iter.Close()
	res := Result{
		Err:     err,
		Columns: cols,
		Rows:    rows,
	}
	if err != nil {
		return &res, errors.Wrap(err, "error when close iter")
	}
	return &res, nil
}

type server struct {
	hosts *hostPool
}

func (srv *server) handleQuery(res http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}
	var q Request
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

func main() {
	port := 8088
	cmd := cobra.Command{
		Use: "goyourcassandra",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("start http server")
			srv := server{
				hosts: newHostPool(),
			}
			mux := http.NewServeMux()
			mux.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte("pong"))
			})
			mux.HandleFunc("/api/ping", func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte("api pong"))
			})
			mux.HandleFunc("/api/query", srv.handleQuery)
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
			addr := ":" + strconv.Itoa(port)
			log.Infof("listen on %s", addr)
			http.ListenAndServe(addr, logged)
		},
	}
	cmd.Flags().IntVar(&port, "port", 8088, "port to listen to")
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
