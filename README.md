# goyourcassandra

Web based Cassandra database browser, named after phpmyadmin

## Roadmap

- [ ] test out gocql
  - [ ] list keyspace
  - [ ] switch between keyspace
  - [ ] desc table etc.
  - [ ] permission operation
  - [ ] get generic result, UI don't know what type of data it is expecting
- [ ] UI
  - [ ] check how auto complete is implemented
  - [ ] pick a framework, might just use jquery
- [ ] support handling JSON

## Usage

````bash
make install
goyourcassandra
````

## Dev

````bash
# start cassandra
make run-c2
# compile and start api server
make reload
# start ui server
cd ui && make serve
# visit http://localhost:8080
````

## Related

- [sosedoff/pgweb](https://github.com/sosedoff/pgweb) Web-based PostgreSQL database browser written in Go.
- [avalanche123/cassandra-web](https://github.com/avalanche123/cassandra-web) 
- [Kindrat/cassandra-client](https://github.com/Kindrat/cassandra-client) Java based desktop UI

## About

The name GoYourCassandra (走你卡珊德拉) is named after phpmyadmin, the well known web based management UI for MySQL written in PHP.
