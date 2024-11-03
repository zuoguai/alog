<template>
    <div >
        <BlankBG >

            <div class="row" >
                <div class="col-2"></div>
                <div class="col-8">
                    <div v-for="article in articleList" :key="article.id">
                        <router-link class="nav-link " style="text-decoration: none; color:black; font-size:medium; "
                        :to="{ name: 'article', query: { article_id: article.id, key: randomString(10) } }">
                       <span class="" style="color:skyblue"> {{timestampToTime( article.create_time * 1000)}}</span>
                       <span style="font-size: larger;"> {{ article.title }}</span>
                    </router-link>
                    <hr class="hr-solid-content">
                    </div>




                    <nav aria-label="Page navigation example">
                        <ul class="pagination  pagination-sm" style="float:right; border-radius: 0%;">
                            <li class="page-item" @click="click_page(-2)" ><a class="page-link" href="#">上一页</a></li>
                            <li @click="click_page(page.number)" :class="'page-item ' + page.is_active"
                                v-for="    page    in    showPage   " :key="page.number">
                                <a class="page-link" href="#">{{ page.number }}</a>
                            </li>
                            <li class="page-item" @click="click_page(-1)"><a class="page-link" href="#">下一页</a></li>
                        </ul>
                    </nav>
                </div>

       

            </div>
            
        </BlankBG>
    </div>

</template>

<script>

import BlankBG from "@/components/BlankBG.vue"
import store from "@/store";
import $ from "jquery"
import { ref } from "vue";

export default {
    name: "ArchivePage",
    components: {
        BlankBG,
    },
    data() {
        return {
            screenWidth: 1024
        }
    },
    mounted() {
        this.screenWidth = window.innerWidth;
        // console.log("屏幕宽度：" + this.screenWidth);
    },

    methods: {
        randomString(e) {
            e = e || 32;
            var t = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678",
                a = t.length,
                n = "";
            for (let i = 0; i < e; i++) n += t.charAt(Math.floor(Math.random() * a));
            return n
        },
    },

    setup() {


        const topTitle = ref("")
        const articleList = ref([]);
        let currentPage = 1;
        let defaultSize = 20
        let totalPage = 1;
        let showPage = ref([]);

        const toPageNum = ref("");

        const gotoPageNum = () => {
            if (toPageNum.value <= 0 || toPageNum.value > totalPage) {
                toPageNum.value = 1
                // alert("页码超出范围")
                return
            }
            if (toPageNum.value == currentPage) {
                return
            }
            pull_page(toPageNum.value)

        }

        const update_pages = () => {
            let max_pages = totalPage;
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

        const click_page = page => {
            if (page === -1) {
                page = currentPage + 1;

            } else if (page === -2) {
                page = currentPage - 1;
            }
            let max_pages = totalPage;
            if (page >= 1 && page <= max_pages) {
                pull_page(page);
            }
        }

        const totalCount = ref("")
        totalCount.value = 1
        const pull_page = page => {
            currentPage = page;


            $.ajax({
                url: store.state.BaseUrl + "/api/article/list?page=" + page + "&size="+defaultSize ,
                type: "get",
                headers: {
                    "Authorization":  store.state.user.token,
                },
                success(resp) {
                    

                    articleList.value = resp.data.list;
                    // console.log(resp.data.list)
                    // console.log(resp.data.count)
                    totalCount.value = resp.data.count;
                    totalPage = (resp.data.count + defaultSize) / defaultSize
                    update_pages();
                    if(page == 1){
                        topTitle.value = resp.data.list[0].title

                    }


                },

            })
            
        }

        const timestampToTime = (timestamp) => {
            timestamp = timestamp ? timestamp : null;
            let date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
            let Y = date.getFullYear() + '-';
            let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
            let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' ';
            // let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
            // let m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
            // let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
            return Y + M + D ;
        }





        pull_page(currentPage)




        return {
            // hotArticleList,
            click_page,
            articleList,
            showPage,
            timestampToTime,
            topTitle,
            pull_page,
            totalCount,
            gotoPageNum,
            toPageNum
        }


    }
}


</script>

<style scoped>

.hr-solid-content {
    /* color: #a2a9b6; */
    color: black;
    border: 1;
    font-size: 12px;
    padding: 0em 0;
    margin: 10px 0;
    position: relative;
}
.text-stroke {
    font-weight: bold;
      color: #cc2525;

      text-shadow:
        -1px -1px 0 rgb(255, 255, 255),
        1px -1px 0 rgb(255, 255, 255),
        -1px 1px 0 rgb(255, 255, 255),
        1px 1px 0 rgb(255, 255, 255); 
    }
/* img {
  opacity: 0;
}
img:hover{
    opacity: 1;
} */
.article-side-padding {
    /* background-color: white; */
    /* padding-left: 2vw; */
    /* border-radius: 0.3%; */
    position: sticky;
    top: 50px;
    font-family: "宋体";
    /* max-width: 20vw; */

}

.right-tag-padding {
    font-family: "宋体";
    /* background-color: rgba(255, 255, 255, 0.2); */

    color: white;
    font-size: medium;
    /* border: 1px solid rgb(216, 211, 192); */

    /* border-radius: 1%; */
    padding-right: 1vw;
    height: 15vh;
    overflow-y: auto;
    /* width: 17vw; */
}

.right-article-padding {
    /* background-color: rgba(255, 255, 255, 0.2); */

    /* border: 1px solid rgb(175, 255, 151); */

    font-size: large;
    max-height: 60vh;
    overflow-y: auto;
    /* position: fixed; */
}

.pre-text {
    background-color: rgba(0, 0, 0, 0);
    color: #ededed;
}

.opacity-for-img{
    opacity: 0.95;
}

.opacity-for-img:hover {
    opacity: 1;
}

/* 鼠标悬停出现下划线 */
.mid-text:hover {

    opacity: 1;
    

    color: black;
    background-color: rgba(255, 255, 255, 0.99)
}

.right-text:hover {

    border: 2px solid rgb(255, 255, 255);
}

.hr-wavy {
    border: 0;
    padding: 3px;
    background: repeating-linear-gradient(135deg, black 0px, #0358f7 1px, transparent 1px, transparent 6px);
}

div::-webkit-scrollbar {
    width: 8px;
    height: 1px;
}

div::-webkit-scrollbar-thumb {
    background: grey;
}

div::-webkit-scrollbar-track {
    background: #ededed;
    box-shadow: inset 0 0 5px white;
}
</style>