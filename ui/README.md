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