<template>
    <div class="DB">
        <h2>DB: {{ name }}</h2>
        <textarea cols="30" rows="10" placeholder="enter CQL" v-model="sql"></textarea>
        <button v-on:click="query">query</button>
        <textarea cols="30" rows="10" disabled v-model="result"></textarea>
    </div>
</template>

<script>
    import axios from 'axios'

    export default {
        name: "DB",
        props: {
            name: String
        },
        data() {
            return {
                sql: '',
                result: '',
            }
        },
        methods: {
            query: function () {
                if (this.sql === '') {
                    // eslint-disable-next-line
                    console.warn('empty sql')
                    return
                }
                // axios.get('/api/ping').then(res => {
                //     this.sql = res.data
                // }).catch(e => {
                //     // eslint-disable-next-line
                //     console.warn(e)
                // })
                axios.post('/api/query', {
                    keyspace: 'system',
                    query: this.sql
                }).then(res => {
                    this.result = JSON.stringify(res.data)
                }, err => {
                    // TODO: it seems when server 500, err does not contain body?
                    console.warn(err)
                }).catch(e => {
                    // eslint-disable-next-line
                    console.warn(e)
                })
            }
        }
    }
</script>

<style scoped>

</style>