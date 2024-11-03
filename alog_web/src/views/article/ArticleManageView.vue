<template>
    <BlankG style="opacity: 0.99;" v-if="screenWidth >= 1000">

        <div class="row" style="padding-top: 2vh;">

            <div class="col-1"></div>
            <div class="col-10">
                <nav class="navbar navbar-light bg-light" style="border: 1px solid rgb(255, 234, 151);">
                    <div class="container-fluid">
                        <a class="navbar-brand">文章管理</a>
                        <div class="d-flex">
                            <input v-model="hasTittle" class="form-control me-2" type="search" placeholder="查找文章标题"
                                aria-label="Search">
                            <button @click="search_article_byTittle()" class="btn btn-outline-success"
                                type="button">Search</button>
                        </div>
                    </div>
                </nav>
                <br>




                <div class="list-group">
                    <div v-for="article in articleList" :key="article.id">
                        <a href="#" class="list-group-item list-group-item-action "
                            style="border: 1px solid rgb(255, 234, 151);">
                            <div class="d-flex w-100 justify-content-between">
                                <h5 class="mb-1" style="color:black;font-weight: bolder;font-size: larger">{{ article.title }}</h5>
                                <small>创建时间： {{ timestampToTime(
                                    article.create_time * 1000) }}</small>

                            </div>
           
                            <p class="mb-1" style="color:black;font-size: medium">简介：{{ article.brief }}</p>





                            <span>

                                <!-- <button type="button" @click="confirm_delete_article(article.id)"
                                    class="btn btn-sm btn-danger" style="float: right; margin-left: 3px;">删除</button> -->
                                <button type="button" @click="modify_article(article.id)"
                                    class="btn btn-sm btn-warning" style="float: right; margin-left: 3px;">修改</button>
                                <button type="button" @click="draft_article(article.id)"
                                    class="btn btn-sm btn-success" style="float: right; margin-left: 3px;">草稿</button>
                                <button type="button" @click="show_article(article.id)"
                                    class="btn btn-sm btn-info" style="float: right; margin-left: 3px;">查看</button>

                            </span>



                            <span>
                                <text style="font-size: 12px; float: left;">
                                    热度：{{ article.view }}&nbsp;
                                </text>
                            </span>
                            <small>&nbsp;</small>

                        </a>
                    </div>

                </div>
                <br>
                <nav aria-label="Page navigation example">
                    <ul class="pagination  pagination-sm " style="float:right">
                        <li class="page-item" @click="click_page(-2, articleList)"><a class="page-link" href="#">上一页</a>
                        </li>
                        <li @click="click_page(page.number, articleList)" :class="'page-item ' + page.is_active"
                            v-for=" page in showPage" :key="page.number">
                            <a class="page-link" href="#">{{ page.number }}</a>
                        </li>

                        <li class="page-item" @click="click_page(-1, articleList)"><a class="page-link" href="#">下一页</a>
                        </li>
                    </ul>
                </nav>

            </div>

        </div>


    </BlankG>
    <div v-else>

        <div style="padding: 1vw;">
            <nav class="navbar navbar-light bg-light" style="border: 1px solid rgb(255, 234, 151);">
                <div class="container-fluid">

                    <a class="navbar-brand" >文章管理</a>

                    <div class="d-flex">
                        <input v-model="hasTittle" class="form-control me-2" type="search" placeholder="查找文章标题"
                            aria-label="Search">
                        <button @click="search_article_byTittle()" class="btn btn-outline-success"
                            type="button">Search</button>
                    </div>
                </div>
            </nav>
        </div>

        <div class="list-group" style="padding: 1vw;">
            <div v-for="article in articleList" :key="article.id">
                <a href="#" class="list-group-item list-group-item-action " style="border: 1px solid rgb(255, 234, 151);">
                    <div class="d-flex w-100 justify-content-between">
                        <h5 class="mb-1">{{ article.title }}</h5>
                        <p style="font-size: 10px;">{{ timestampToTime(article.create_time * 1000) }}</p>

                    </div>
                    <div>
                        <p class="mb-1" style="font-size: 12px;">简介：{{ article.brief }}</p>

                        
                    </div>
                    <br>
                    <span>

                        <!-- <button type="button" @click="confirm_delete_article(article.id)"
                            class="btn btn-sm btn-danger" style="float: right; margin-left: 3px;">删除</button> -->
                        <button type="button" @click="modify_article(article.id)"
                            class="btn btn-sm btn-warning" style="float: right; margin-left: 3px;">修改</button>
                        <button type="button" @click="draft_article(article.id)"
                            class="btn btn-sm btn-success" style="float: right; margin-left: 3px;">草稿</button>
                        <button type="button" @click="show_article(article.id)"
                            class="btn btn-sm btn-info" style="float: right; margin-left: 3px;">查看</button>

                    </span>


                    <p style="font-size: 10px;">热度：{{ article.view }}

                        <br>

                    </p>

                </a>
            </div>

        </div>
        <nav aria-label="Page navigation example">
            <ul class="pagination  pagination-sm " style="float:right">
                <li class="page-item" @click="click_page(-2, articleList)"><a class="page-link" href="#">上一页</a>
                </li>
                <li @click="click_page(page.number, articleList)" :class="'page-item ' + page.is_active"
                    v-for=" page in showPage" :key="page.number">
                    <a class="page-link" href="#">{{ page.number }}</a>
                </li>

                <li class="page-item" @click="click_page(-1, articleList)"><a class="page-link" href="#">下一页</a>
                </li>
            </ul>
        </nav>



    </div>
</template>
  
<script>
import BlankG from '@/components/BlankBG.vue';
import $ from 'jquery';
import { useStore } from 'vuex';
import { ref } from 'vue';
import router from "@/router";
export default {
    components: {
        BlankG,
    },
    data() {
        return {
            screenWidth: 1024,
        }
    },
    mounted() {
        this.screenWidth = window.innerWidth
    },
    setup() {
        const store = useStore();

        const articleList = ref([])
        // const PAGE = 6;
        let currentPage = 1;
        let totalPage = 1;
        let showPage = ref([]);
        const hasTittle = ref("")



        const draft_article = (article_id) => {
            $.ajax({
                //url格式：article/list/页码/用户id/是否需要草稿/
                url: store.state.BaseUrl + "/api/admin" ,
                type: "put",
                headers: {
                    Authorization:  store.state.user.token,
                },
                data: {
                    "id":article_id,
                    "status":2,
                },

                success(resp) {
                    if (resp.msg == 'ok') {
                        alert("加入草稿箱成功")
                        router.push({
                        path: "/article-manage",
                    })
                    } else {
                        alert("加入草稿箱失败")
                    }
                    


                },
                error() {
                    alert("删除失败")
                },


            })
            
            

        }

        const show_article = (article_id) => {
            router.push({
                path: "/article",
                query: {
                    article_id: article_id
                }
            })
        }

        const modify_article = (article_id) => {
            router.push({
                path: "/article-modify",
                query: {
                    // authorId: store.state.user.id,
                    article_id: article_id
                }
            })
        }

        const confirm_delete_article = (article_id) => {
            let flag;
            flag = confirm("是否确认永久删除，无法恢复")
            if (flag == true) {
                delete_article(article_id);
            }
        }

        const delete_article = (article_id) => {
            $.ajax({
                //url格式：article/list/页码/用户id/是否需要草稿/
                url: store.state.BaseUrl + "/api/admin" ,
                type: "put",
                headers: {
                    Authorization:  store.state.user.token,
                },
                data: {
                    "id":article_id,
                    "status":3,
                },

                success(resp) {
                    if (resp.msg == 'ok') {
                        alert("加入删除成功")
                    } else {
                        alert("加入删除失败")
                    }


                },
                error() {
                    alert("删除失败")
                },


            })
            router.push({
                path: "/article-manage",
            })
            

        }


        const update_pages = () => {
            // let max_pages = parseInt(Math.ceil(totalPage / PAGE));
            let max_pages = totalPage;
            console.log("最大页数" + max_pages)

            let new_page = [];
            for (let i = currentPage - 2; i < currentPage + 2; i++) {
                if (i >= 1 && i <= max_pages) {
                    new_page.push({
                        number: i,
                        is_active: i === currentPage ? "active" : "",
                    })
                }
            }

            showPage.value = new_page;

        }

        const click_page = (page) => {
            if (page === -1) {
                page = currentPage + 1;

            } else if (page === -2) {
                page = currentPage - 1;
            }
            let max_pages = totalPage;
            if (page >= 1 && page <= max_pages) {
                console.log("获取页面：" + page)
                pull_page(page);
            }
        }


        const totalCount = ref("");
        let defaultSize = 10
        const pull_page = (page) => {
            currentPage = page;
                $.ajax({
                    url: store.state.BaseUrl + "/api/article/list?page=" + page + "&size="+defaultSize +"&status=1" ,
                    type: "get",
                    headers: {
                    "Authorization":  store.state.user.token,
                },
                    success(resp) {

                        articleList.value = resp.data.list;
                        console.log(resp);
                        totalPage = parseInt((resp.data.count + defaultSize) / defaultSize);
                        totalCount.value = resp.data.count;
                        update_pages();
                        console.log("拉取文章列表成功")
                    },

                })
        }
        const search_article_byTittle = () => {

        }






        const timestampToTime = (timestamp) => {
            timestamp = timestamp ? timestamp : null;
            let date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
            let Y = date.getFullYear() + '-';
            let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
            let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' ';
            let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
            let m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
            let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
            return Y + M + D + h + m + s;
        }

        pull_page(currentPage)
        return {
            articleList,
            click_page,
            showPage,
            show_article,
            delete_article,
            modify_article,
            draft_article,
            confirm_delete_article,
            search_article_byTittle,
            hasTittle,
            timestampToTime,
            totalCount
        }





    }



};

</script>

<style scoped>
.article-list-padding {
    height: 50vh;
    width: 55vw;
    position: fixed;
    overflow-y: scroll;
}
</style>