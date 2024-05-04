<script>
    import { makeFetch } from "../assets/js/request_forgery.js";
    export default {
        props: {
            list_title: String,
            url: {
                type: String,
                required: true
            },
        },
        data() {
            return {
                requestData: {}
            }
        },
        methods: {
            async Get(url) {
                return await makeFetch(url);;
            }
        },
        async created() {
            this.requestData = await this.Get(this.url);
        }
    }
</script>


<template>
    <div class="section">
        <div class="list d-flex flex-column align-items-center">
            <h3 class="list__title">{{ list_title }}</h3>
                <div v-for="(field, fieldName) in requestData" class="list__block d-flex" :key="fieldName">
                    <p class="list__label">{{fieldName}}</p>
                    <p class="list__data">{{ field }}</p>
                </div>
        </div>
    </div>
</template>




<style scoped>
    .list__title {
        font-size: 35px;
        font-weight: lighter;
        font-style: italic;
        letter-spacing: 2px;
    }

    .list * {
        margin: 5px;
    }

    .list {
        text-align: center;
    }  
</style>