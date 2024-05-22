<script>
    import { makeFetch } from "../assets/js/request_forgery.js";
    export default {
        props: {
            form_title: String,
            url: {
                type: String,
                required: true
            },
            dataset: {
                type: Object,
                required: true,
            } 
        },
        data() {
            return {
                requestData: {}
            }
        },
        methods: {
            async submitForm() {
                console.log("Form submitted:", this.requestData);
                try {
                    const response = await makeFetch(this.url, this.requestData);
                    console.log("Response from server: ", response);
                } catch (error) {
                    console.error("Error:", error);
                }
            },
            transferData(data) {
                var response = {};
                for (var key of Object.keys(data)) {
                    response[key] = "";
                }
                return response;
            }
        },
        created() {
            this.requestData = this.transferData(this.dataset);
        }
    }
</script>


<template>
    <div class="section">
        <div class="form d-flex flex-column align-items-center">
            <h3 class="form__title">{{ form_title }}</h3>
                <div v-for="(field, fieldName) in dataset" :key="fieldName">
                    <p class="form__input-label">{{ field.text }}</p>
                    <input :type="field.type" v-model="requestData[fieldName]" class="form__input-text" :placeholder="field.text">
                </div>
            <button @click="submitForm" class="form__submit">Submit</button>
        </div>
    </div>
</template>




<style scoped>
    .form__title {
        font-size: 35px;
        font-weight: lighter;
        font-style: italic;
        letter-spacing: 2px;
    }

    .form__submit {
        padding: 8px 25px; /* Устанавливаем отступы только сверху и снизу */
        border: 1.5px solid rgb(53, 57, 94);
        border-radius: 6px;
        font-size: 20px;
        margin: 0 auto;
        margin-bottom: 5px; /* добавляем отступ только снизу */
        box-sizing: border-box;
        transition: 500ms;
        background-color: rgb(147, 147, 179);
        color: rebeccapurple;

    }

    .form__submit:hover {
        background-color: #7d88aa;
    }

    .form__input-label {
        font-style: italic;
        font-size: 20px;
        margin-bottom: 10px;
        text-transform: capitalize;

    }

    .form * {
        margin: 5px;
    }
    .form__input-text {
        padding: 8px 25px; /* Устанавливаем отступы только сверху и снизу */
        border: 1.5px solid rgb(134, 136, 156);
        border-radius: 6px;
        font-size: 20px;
        margin: 0 auto;
        color: rebeccapurple;
        margin-bottom: 5px; /* добавляем отступ только снизу */
        box-sizing: border-box; /* Учитываем внутренние отступы в общей ширине элемента */
        background-color: #bbbfce;
    }


    .form {
        text-align: center;
    }  
</style>