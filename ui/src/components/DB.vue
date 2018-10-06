<template>
    <div class="DB">
        <div>
            <!-- TODO: totally forgot how to layout a form -->
            <h2>DB: {{ name }}</h2>
            <label>host</label>
            <input type="text" v-model="host">
            <label>keyspace</label>
            <input type="text" v-model="keyspace">
            <textarea cols="30" rows="10" placeholder="enter CQL" v-model="sql"></textarea>
            <button v-on:click="query">query</button>
            <textarea cols="30" rows="10" disabled v-model="result"></textarea>
        </div>
        <div>
            <!--https://vuejs.org/v2/guide/list.html-->
            <table>
                <thead>
                <tr>
                    <!-- NOTE: it's value, key ... -->
                    <td v-for="(col, index) in cols" :key="index">
                        {{col}}
                    </td>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(row, index) in rows" :key="index">
                    <td v-for="(col, index) of cols" :key="index">
                        {{row[col]}}
                    </td>
                </tr>
                </tbody>
            </table>
        </div>
        <div>
            <h2>Keyspace</h2>
            <button v-on:click="fetchKeyspace">Fetch</button>
        </div>
        <div>
            <div v-for="(tbl, name) in keyspaceTables" :key="name">
                {{name}}
                <table>
                    <thead>
                    <tr>
                        <td>name</td>
                        <td>type</td>
                    </tr>
                    </thead>
                    <tbody>
                    <tr v-for="(colName, index) in tbl.orderedColumns" :key="index">
                        <td>{{colName}}</td>
                        <td>{{tbl.columns[colName]}}</td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";

export default {
  name: "DB",
  props: {
    name: String
  },
  data() {
    return {
      host: "localhost",
      keyspace: "system",
      sql: "",
      result: "",
      cols: [],
      rows: [],
      keyspaceTables: {}
    };
  },
  methods: {
    query: function() {
      if (this.sql === "") {
        // eslint-disable-next-line
        console.warn("empty sql");
        return;
      }
      axios
        .post("/api/query", {
          host: this.host,
          keyspace: this.keyspace,
          query: this.sql
        })
        .then(
          res => {
            this.result = JSON.stringify(res.data);
            // TODO: check if there is error
            // TODO: anyway to give it type? ... now I miss typescript ...
            let data = res.data;
            let cols = [];
            if (data.columns != null) {
              for (let v of data.columns) {
                // eslint-disable-next-line
                console.log(v.name);
                cols.push(v.name);
              }
            }
            this.cols = cols;
            // eslint-disable-next-line
            console.log(this.cols);
            let rows = [];
            if (data.rows != null) {
              for (let v of data.rows) {
                // eslint-disable-next-line
                console.log(v);
                rows.push(v);
              }
            }
            this.rows = rows;
            // eslint-disable-next-line
            console.log(this.rows);
          },
          err => {
            // TODO: it seems when server 500, err does not contain body?
            // eslint-disable-next-line
            console.warn(err);
          }
        )
        .catch(e => {
          // eslint-disable-next-line
          console.warn(e);
        });
    },
    fetchKeyspace: function() {
      axios
        .get("/api/keyspace", {
          params: {
            host: this.host,
            keyspace: this.keyspace
          }
        })
        .then(
          res => {
            let decoded = res.data;
            // TODO: check if there is error
            // TODO: anyway to give it type? ... now I miss typescript ...
            // eslint-disable-next-line
            console.log(decoded);
            this.keyspaceTables = decoded.tables;
          },
          err => {
            // TODO: it seems when server 500, err does not contain body?
            // eslint-disable-next-line
            console.warn(err);
          }
        )
        .catch(e => {
          // eslint-disable-next-line
          console.warn(e);
        });
    }
  }
};
</script>

<style scoped>
</style>
