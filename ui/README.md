# goyourcassandra

## Dev

- build and start backend

````bash
make install
make serve
````

- start dev server, it will proxy request in `/api` to `localhost:8088`

```
npm install
# Compiles and hot-reloads for development
npm run serve
# Compiles and minifies for production
npm run build
# Lints and fixes files
npm run lint
```

Example queries

````sql
-- NOTE: this seems only works with cassandra 2
select * from system.schema_keyspaces;
````

````sql
CREATE KEYSPACE app WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1};
CREATE TABLE app.tb1 (id int, s text,  PRIMARY KEY (id));
INSERT INTO app.tb1 (id, s) VALUES (1, 'apptb11');
INSERT INTO app.tb1 (id, s) VALUES (2, 'apptb12');
SELECT * FROM app.tb1
````

## UI

- use https://github.com/ElemeFE/element

````bash
npm i element-ui -S
````