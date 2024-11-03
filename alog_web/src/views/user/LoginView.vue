<template>
    <BlankBG v-if="screenWidth >= 600">
        <div class="row justify-content-md-center">
            <div class="col-3">
                <form @submit.prevent="login">
                    <div class="mb-3">
                        <label for="username" class="form-label">用户名:</label>
                        <input v-model="username" type="text" class="form-control" id="username" placeholder="请输入用户名">
                    </div>
                    <div class="mb-3">
                        <label for="password" class="form-label">密码:</label>
                        <input v-model="password" type="password" class="form-control" id="password" placeholder="请输入密码">
                    </div>
                    <div class="error-message">{{ error_message }}</div>
                    <button type="submit" class="btn btn-primary">登录</button>
                </form>
            </div>
        </div>
    </BlankBG>
    <div v-else>
        <form @submit.prevent="login">
            <div class="mb-3">
                <label for="username" class="form-label">用户名:</label>
                <input v-model="username" type="text" class="form-control" id="username" placeholder="请输入用户名">
            </div>
            <div class="mb-3">
                <label for="password" class="form-label">密码:</label>
                <input v-model="password" type="password" class="form-control" id="password" placeholder="请输入密码">
            </div>
            <div class="error-message">{{ error_message }}</div>
            <button type="submit" class="btn btn-primary">登录</button>
        </form>
    </div>
</template>
    
<script>

import BlankBG from "@/components/BlankBG.vue"
import { ref } from "vue";
import router from "@/router";
import store from "@/store";

export default {
    name: "LoginPage",
    components: {
        BlankBG,
    },

    data() {
        return {
            screenWidth: 1024,
        }
    },

    mounted() {
        // console.log("登录页面挂起")
        this.screenWidth = window.innerWidth

        const a_log_token = localStorage.getItem("a_log_token");
        if (a_log_token != null && a_log_token != "" && !store.state.user.is_login) {

            store.commit("updateToken", a_log_token);
            
            router.push({ name: "home" })
                
            
        }

    },

    setup() {
        let username = ref('');
        let password = ref('');
        let error_message = ref('');


        const login = () => {
            error_message.value = "";
            store.dispatch("login", {
                username: username.value,
                password: password.value,
                success() {
                    console.log("登录成功")
                    router.push({ name: "home" })
                },
                error(resp) {
                    console.log(resp)
                    error_message.value = resp
                }
            })

        }




        return {
            username,
            password,
            error_message,
            login
        }
    }
}
</script>
    
<style scoped>
button {
    width: 100%;
}

div.error-message {
    color: red;
}
</style>