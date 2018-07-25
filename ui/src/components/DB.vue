<template>
    <div class="DB">
        <h2>DB: {{ name }}</h2>
        <textarea cols="30" rows="10" placeholder="enter CQL" v-model="sql"></textarea>
        <button v-on:click="query">query</button>
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
            }
        },
        methods: {
            query: function () {
                if (this.sql === '') {
                    // eslint-disable-next-line
                    console.warn('empty sql')
                    return
                }
                axios.get('/api/ping').then(res => {
                    this.sql = res.data
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