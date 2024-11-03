<template>
    <BlankG v-if="screenWidth >= 1000" style="opacity: 0.99;">



        <div class="row" style="padding-top: 5vh;">
            <div class="col-3">
                <div class="set-article-padding">
                    <div class="mb-3">
                        <label for="titleControlInput" class="form-label">标题：</label>
                        <input type="text" class="form-control" id="titleControlInput" v-model="title"
                            placeholder="标题不能小于2字符">
                    </div>
                    <div class="mb-3">
                        <label for="FormControlTextarea" class="form-label">摘要：</label>
                        <textarea class="form-control" id="FormControlTextarea" rows="3" v-model="brief"
                            placeholder="写文章简介"></textarea>
                    </div>

                   

                    <div class="mb-3">
                        <div v-if="isLoading" class="spinner-border text-primary col-sm-2" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                        <button v-else class="btn btn-bg btn-info" style="width: 100%;"
                            @click="upload_article(article_id, title, brief, text, articleTemp)">发布</button>

                    </div>
                    <div class="mb-3">

                    </div>
                </div>
            </div>
            <div class="col-9">

                <div class="write-article-padding">
                    <label class="form-label">文章编写：</label>
                    <v-md-editor v-model="text" height="70vh" mode="editable" default-show-toc="true"
                        placeholder="在这里编写您的文章"
                        left-toolbar="undo redo clear | emoji | h bold italic strikethrough quote | ul ol table hr | link image code "></v-md-editor>

                </div>


            </div>

        </div>





    </BlankG>
    <div v-else>
        <div class="mb-3">
            <label for="titleControlInput" class="form-label">标题：</label>
            <input type="text" class="form-control" id="titleControlInput" v-model="title" placeholder="标题不能小于2字符">
        </div>
        <div class="write-article-padding">
            <label class="form-label">文章编写：</label>
            <v-md-editor v-model="text" height="50vh" mode="edit" placeholder="在这里编写您的文章"
                left-toolbar="undo redo clear | emoji | h bold italic strikethrough quote | ul ol table hr | link image code "></v-md-editor>
        </div>

        <div class="set-article-padding">

            <div class="mb-3 ">
                <label for="formFile" class=" col-form-label">文章封面（可无）：</label>
                <input class="form-control" type="file" accept="image/gif,image/jpeg,image/jpg,image/png,image/svg"
                    placeholder="选择文件" id="articleBgFile">
            </div>
            <div class="mb-3">
                <label for="FormControlTextarea" class="form-label">摘要：</label>
                <textarea class="form-control" id="FormControlTextarea" rows="3" v-model="brief"
                    placeholder="写文章简介"></textarea>
            </div>



            <div class="mb-3">
                <div v-if="isLoading" class="spinner-border text-primary col-sm-2" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>
                <button v-else class="btn btn-bg btn-info" style="width: 100%;"
                    @click="upload_article(article_id, title, brief, text, articleTemp)">发布</button>

            </div>
            <div class="mb-3">
            </div>
        </div>
        <br>
        <br>
    </div>
</template>

<script>
import BlankG from '@/components/BlankBG.vue';
import { useStore } from 'vuex';
import $ from "jquery"
// import router from '@/router';

import { ref } from 'vue';
import router from '@/router';

export default {
    components: {
        BlankG,
    },
    data() {
        return {
            title: "",
            text: '',
            brief: "",
            article_id: 0,
            articleTemp: "",
            screenWidth: 1024,
            authorId: 0,

        };
    },

    mounted() {
        const that = this
        const article_id = this.$route.query.article_id

        that.article_id = this.$route.query.article_id
        this.screenWidth = window.innerWidth
        const store = useStore();



        $.ajax({
            url: store.state.BaseUrl + "/api/article/" + article_id,
            type: "GET",
            headers: {
                    "Authorization":  store.state.user.token,
                },
            success(resp) {

                if (resp.msg == 'ok') {
                    that.authorId = resp.data.list[0].id
                    that.text = resp.data.list[0].content
                    that.title = resp.data.list[0].title
                    that.brief = resp.data.list[0].brief
                    that.articleTemp = resp.data.list[0]

                }
                else {
                    alert("获取文章失败")
                }


            },
            error() {
                alert("获取文章失败,请检查网络连接")

            }

        })

    




    },
    methods: {

    },

    setup() {
        const store = useStore();


        const isLoading = ref("")
        isLoading.value = false;





        const upload_article = (article_id, title, brief, text, articleTemp) => {
            isLoading.value = true;


            if (articleTemp.title == title && articleTemp.brief == brief && articleTemp.content == text
                 ) {

                isLoading.value = false;

                alert("好像还没有更改文章")
                return

            }



            if (title == null || title.length <= 2) {
                alert("标题不能少于两个字符")
                return
            }
            if (text == null || text.length <= 20) {
                alert("文章不能少于20个字符")
                return
            }

            $.ajaxSetup({
                traditional: true
            });


            setTimeout(() => {
                $.ajax({
                    url: store.state.BaseUrl + "/api/admin",
                    type: "put",
                    headers: {
                        "Authorization":  store.state.user.token
                    },
                    data: {
                        "id": article_id,
                        "title": title,
                        "brief": brief,
                        "content": text,

                    },
                    success(resp) {
                        if (resp.msg == "ok") {
                            title = "",
                                brief = "",
                                text = "",
                            isLoading.value = false;
                            alert("文章发布成功")
                            router.push({
                                path: "/article",
                                query: {
                                    article_id: article_id
                                }
                            })
                        } else {
                            alert("文章发布失败")

                            isLoading.value = false;
                        }
                    },
                    error() {
                        alert("文章发布失败")
                        isLoading.value = false;
                    }
                })
            }, 10)

        }
        return {
            upload_article,
            isLoading,
        }
    }


};
</script>
<style scoped>
.set-article-padding {
    /* background-color: aliceblue; */
    /* height: 00px; */
    padding-top: 3vh;
}

.write-article-padding {}
</style>