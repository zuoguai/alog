<template>
    <nav  class="navbar navbar-expand-lg navbar-fixed-top navbar-dark"
        style="background-color:rgba(0, 0, 0, 0.5);  font-size: larger;  font-family: '宋体'; ">
        <div class="container ">
            <router-link class="navbar-brand tttext" to="/">
                好好加油！{{ mouseX }}%
            </router-link>

            <div class="collapse navbar-collapse" id="navbarText">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">

                </ul>
                <ul class="navbar-nav">

                </ul>

                <ul v-if="$store.state.user.is_login" class="navbar-nav">
                    <li class="nav-item">
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link tttext" :to="{ name: 'home' }">首页</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link tttext" :to="{ name: 'archive' }">归档</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link tttext" :to="{ name: 'article_write' }">写作</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link tttext" :to="{ name: 'article_manage' }">管理</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link tttext" :to="{ name: 'article_draft' }">草稿</router-link>
                    </li>





                    <!-- <li class="nav-item dropdown">
                        <router-link class="nav-link dropdown-toggle tttext" to="" role="button"
                            data-bs-toggle="dropdown" aria-expanded="false">
                            写作
                        </router-link>
                        <ul class="dropdown-menu">
                            <li><router-link class="dropdown-item tttext"
                                    :to="{ name: 'article_write' }">编写文章</router-link></li>
                            <li>
                                <hr class="dropdown-divider">
                            </li>
                            <li><router-link class="dropdown-item tttext"
                                    :to="{ name: 'article_manage' }">文章管理</router-link></li>
                            <li>
                                <hr class="dropdown-divider">
                            </li>

                            <li><router-link class="dropdown-item tttext"
                                    :to="{ name: 'article_draft' }">草稿箱</router-link></li>
                        </ul>
                    </li> -->

                    <li class="nav-item dropdown">
                        <router-link class="nav-link dropdown-toggle tttext" to="" role="button"
                            data-bs-toggle="dropdown" aria-expanded="false">
                            {{ $store.state.user.username }}
                        </router-link>
                        <ul class="dropdown-menu">
                            <li>
                            <a href="https://zuoguai.github.io/AboutMe" class="dropdown-item tttext" target="_blank">AboutMe</a>

                            </li>
                            <li>
                                <a @click="confirm_user_logout" class="dropdown-item tttext" href="">注销</a>
                            </li>
                        </ul>
                    </li>
                </ul>
                <ul v-else class="navbar-nav">


                    <li class="nav-item">
                        <router-link class="nav-link tttext" :to="{ name: 'home' }">首页</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link tttext" :to="{ name: 'archive' }">归档</router-link>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link tttext" href="https://zuoguai.github.io/AboutMe"  target="_blank">作怪</a>
                    </li>




                    <!-- <li class="nav-item dropdown">


                            <router-link class="nav-link dropdown-toggle tttext" to="" role="button"
                            data-bs-toggle="dropdown" aria-expanded="false">

                            作怪
                        </router-link>
                        
                        
                            <ul class="dropdown-menu tttext"
                            style="background-color:rgba(255, 255, 255, 0.9); max-height: 200px; overflow-y: scroll;">
                            <li>

                                <router-link class="dropdown-item tttext" to="" @click="show_label_article('日常', 1)">日常</router-link>
                            </li>
                            <li>
                            <a href="https://zuoguai.github.io/AboutMe" class="dropdown-item tttext" target="_blank">AboutMe</a>

                            </li>

                        </ul>


                    </li> -->

                </ul>

            </div>
        </div>
    </nav>

</template>


<script>
import store from '@/store';
import router from '@/router';
// import { ref } from 'vue';
// import $ from "jquery"
export default {
    name: "NavBar",
    components: {
    },
    props: {
    mouseX: String,
    mouseY: String,
  },
    data() {
        return {
        }
    },
    mounted() {

    },
    methods: {

        alert_msg(msg) {
            alert(msg)
        },

    },

    setup() {

        const confirm_user_logout = () => {
            let flag;
            flag = confirm("是否确认注销，将退出登录")
            if (flag == true) {
                userLogout()
            } else {
                return
            }
        }

        const userLogout = () => {

            store.dispatch("logout", {
                success() {
                    router.push({ name: "login" })
                },
                error() {

                }
            }
            )


        }

        // const articleList = ref([]);
        // let currentPage = 1;
        // let totalPage = 1;
        // let showPage = ref([]);

        // const hasTittle = ref("")



        // // 检测登录
        if (!store.state.user.is_login) {
            const a_log_token = localStorage.getItem("a_log_token");
            if (a_log_token != null && a_log_token != "" && !store.state.user.is_login) {
                store.commit("updateToken", a_log_token);
            }
        }

      






        // const show_article = (articleId) => {
        //     hasTittle.value = ""
        //     router.push({
        //         path: "/article",
        //         query: {
        //             authorId: store.state.user.id,
        //             articleId: articleId
        //         }
        //     })
        // }

        // const update_pages = () => {
        //     let max_pages = totalPage;
        //     console.log("最大页数" + max_pages)

        //     let new_page = [];
        //     for (let i = currentPage - 2; i < currentPage + 2; i++) {
        //         if (i >= 1 && i <= max_pages) {
        //             new_page.push({
        //                 number: i,
        //                 is_active: i === currentPage ? "active" : "",
        //             })
        //         }
        //     }

        //     showPage.value = new_page;

        // }

        // const click_page = (page) => {
        //     if (page === -1) {
        //         page = currentPage + 1;

        //     } else if (page === -2) {
        //         page = currentPage - 1;
        //     }
        //     let max_pages = totalPage;
        //     if (page >= 1 && page <= max_pages) {
        //         console.log("获取页面：" + page)
        //         pull_page(page);
        //     }
        // }

        // const pull_page = (page) => {
        //     if (hasTittle.value == null || hasTittle.value == "") return;
        //     currentPage = page;

        //     $.ajax({
        //         url: store.state.BaseUrl + "/api/article/list/condition",
        //         type: "post",
        //         headers: {
        //             Authorization: "Bearer " + store.state.user.token,
        //         },
        //         data: {
        //             hasTittle: hasTittle.value,
        //             pageNum: page


        //         },
        //         success(resp) {
        //             articleList.value = resp.content.articleList;
        //             console.log(resp);
        //             totalPage = resp.content.totalPage;
        //             update_pages();
        //             console.log("拉取 " + hasTittle.value + " 文章列表成功")
        //         },
        //     })
        // }
        return {

        //     show_article,
        //     pull_page,
            confirm_user_logout,
        //     articleList,
        //     click_page,
        //     showPage,
        //     hasTittle,
        }
    }
}


</script>


<style scoped>
/* nav a.router-link-exact-active {
  color: #42b983;
} */
 
.progress-container {
  width: 100%;
  background-color: rgba(0, 0, 0, 0);
  border-radius: 4px;
  overflow: hidden;
}
 
.progress-bar {
  height: 2px;
  background-color: #ffffff;
  transition: width 1s ease-in-out;
  border-radius: 4px;
  width: 0; /* 初始宽度为0 */
}
nav:hover {
    background-color: rgb(231, 31, 31);
}

li {
    color: white;
}

.tttext {
    border-bottom: 2px solid #000000;
}

.tttext:hover {
    /* border-bottom: 3px solid #000000; */
    /* border: 5px solid rgb(231,231,231); */
    /* border-radius: 3%; */
    opacity: 1;
    /* border: 1px solid rgb(255, 255, 255); */
    border-bottom: 2px solid #ffffff;

    font-weight: 900;
    color: white;



}


ul::-webkit-scrollbar {
    width: 8px;
    height: 1px;
}

ul::-webkit-scrollbar-thumb {
    background: grey;
}

ul::-webkit-scrollbar-track {
    background: #ededed;
    box-shadow: inset 0 0 5px white;
}
</style>