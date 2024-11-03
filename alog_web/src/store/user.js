import $ from 'jquery'
import store from '.';
// import { all } from 'core-js/fn/promise';

export default ({
    state: {
        id: 1,
        username: "作怪",
        token: "",
        is_login: false,
        is_admin: false,
    },
    getters: {
    },
    mutations: {
        updateUser(state, user) {
            state.id = user.user_id;
            state.username = user.username;
            state.toke = user.token;
            state.is_login = true;
            state.is_admin = true;
        },

        updateToken(state, token) {
            state.token = token;
            state.is_login = true;
        },

        logout(state) {
            state.id = "";
            state.username = "作怪";
            state.token = "";
            state.is_login = false;
            state.is_admin = false;
        }



    },
    actions: {
        
        login(context, data) {
            $.ajax({
                url: store.state.BaseUrl + "/api/login",
                type: "POST",
                data: {
                    username: data.username,
                    password: data.password
                },
                success(resp) {
                    if (resp.msg == "ok") {
                        console.log(resp)
                        //保存token并持久化
                        localStorage.setItem("a_log_token", resp.data.token);
                        context.commit("updateToken", resp.data.Authorization);
                        context.commit("updateUser", resp.data);
                        //登录成功回调
                        data.success()

                    } else {
                        data.error(resp)
                        //登录失败，保持登出状态
                        context.commit("logout");
                    }

                },
                error() {
                    data.error("服务器未响应")
                    //登录失败，保持登出状态
                    context.commit("logout");
                }
            })
        },

        logout(context, data) {
            alert("退出成功")
            localStorage.removeItem("a_log_token")
            context.commit("logout");
            data.success()
        }

    },
    modules: {

    }
})
