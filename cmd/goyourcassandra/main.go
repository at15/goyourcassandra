package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/gocql/gocql"
)

var log = dlog.NewApplicationLogger()

type connPool struct {
	mu sync.RWMutex

	conns map[string]*gocql.Session
}

type Request struct {
	Keyspace string `json:"keyspace"`
	Query    string `json:"query"`
}

type Result struct {
	Err     error
	Columns []gocql.ColumnInfo
	Rows    []map[string]interface{}
}

func newConnPool() *connPool {
	return &connPool{
		conns: make(map[string]*gocql.Session),
	}
}

func (pool *connPool) Query(keyspace string, query string) (*Result, error) {
	pool.mu.RLock()
	session, ok := pool.conns[keyspace]
	pool.mu.RUnlock()
	if !ok {
		cluster := gocql.NewCluster("localhost")
		cluster.Keyspace = keyspace
		tSession, err := cluster.CreateSession()
		if err != nil {
			return nil, errors.Wrapf(err, "error connect use keyspace %s", keyspace)
		}
		pool.mu.Lock()
		pool.conns[keyspace] = tSession
		pool.mu.Unlock()
		session = tSession
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
	err := iter.Close()
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

func main() {
	fmt.Println("start http server")
	pool := newConnPool()
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("pong"))
	})
	mux.HandleFunc("/api/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("api pong"))
	})
	mux.HandleFunc("/api/query", func(res http.ResponseWriter, req *http.Request) {
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
		qRes, err := pool.Query(q.Keyspace, q.Query)
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
