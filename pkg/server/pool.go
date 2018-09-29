package server

import (
	"sync"

	"github.com/dyweb/gommon/errors"
	"github.com/gocql/gocql"

	"github.com/at15/goyourcassandra/pkg/types"
)

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

func (pool *hostPool) Query(host, keyspace, query string) (*types.Result, error) {
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

func (pool *connPool) Query(keyspace string, query string) (*types.Result, error) {
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
	res := types.Result{
		Err:     err,
		Columns: cols,
		Rows:    rows,
	}
	if err != nil {
		return &res, errors.Wrap(err, "error when close iter")
	}
	return &res, nil
}
