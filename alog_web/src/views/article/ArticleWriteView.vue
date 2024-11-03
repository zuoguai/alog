<template>
    <BlankG v-if="screenWidth >= 1000">
        <div class="row" style="padding-top: 2vh;">
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
                            @click="upload_article(title, brief, text)">发布文章</button>

                    </div>
                    <div class="mb-3">

                    </div>
                </div>
            </div>
            <div class="col-9">

                <div class="write-article-padding">
                    <label class="form-label">文章编写：</label>
                    <v-md-editor v-model="text" height="70vh" mode="editable" placeholder="在这里编写您的文章"
                        left-toolbar="undo redo clear | emoji | h bold italic strikethrough quote | ul ol table hr | link image code "></v-md-editor>

                </div>


                <!-- <button class="btn btn-sm btn-warning" @click="upload_article(title, brief, text, labelList)">存为草稿</button> -->
            </div>

        </div>
    </BlankG>

    <div v-else style="padding: 1vw;">
        <div class="mb-3">
            <label for="titleControlInput" class="form-label">标题：</label>
            <input type="text" class="form-control" id="titleControlInput" v-model="title" placeholder="标题不能小于2字符">
        </div>
        <div class="write-article-padding">
            <label class="form-label">文章编写：</label>
            <v-md-editor v-model="text" height="60vh" mode="edit" placeholder="在这里编写您的文章"
                left-toolbar="undo redo clear | emoji | h bold italic strikethrough quote | ul ol table hr | link image code "></v-md-editor>

        </div>

        <div class="set-article-padding">

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
                    @click="upload_article(title, brief, text)">发布文章</button>

            </div>
            <div class="mb-3">

            </div>
            <br>
            <br>
        </div>




    



    </div>
</template>

<script>
import BlankG from '@/components/BlankBG.vue';
import { useStore } from 'vuex';
import $ from "jquery"
import router from '@/router';

import { ref } from 'vue';

export default {
    components: {
        BlankG,
    },
    data() {
        return {
            title: "",
            text: '',
            brief: "",

            screenWidth: 1024,



        };
    },

    mounted() {
        this.screenWidth = window.innerWidth

    },
    methods: {





    },

    setup() {
        const store = useStore();
        // let newLabel = ref("");
        const isLoading = ref("")
        isLoading.value = false;


        const upload_article = (title, brief, text) => {
            isLoading.value = true;


            if (title == null || title.length <= 2) {
                isLoading.value = false;
                alert("标题不能少于两个字符")
                return
            }
            if (text == null || text.length <= 20) {
                isLoading.value = false;
                alert("文章不能少于20个字符")
                return
            }

            $.ajax({
                url: store.state.BaseUrl + "/api/admin",
                type: "POST",
                headers: {
                    "Authorization":   store.state.user.token,
                },
                data: {
                    "title": title,
                    "brief": brief,
                    "content": text,
                },
                success(resp) {
                    isLoading.value = false;
                    if (resp.msg == "ok") {
                            title = "",
                            brief = "",
                            text = "",
                            alert("文章发布成功")
                        router.push({
                            path: "/article-write",
                        })

                    } else {
                        alert("文章发布失败")
                    }
                },
                error() {
                    isLoading.value = false;
                    alert("文章发布失败")

                }
            })
        }
        return {
            isLoading,
            upload_article,
            // newLabel
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

.write-article-padding {
    color: black;
}
</style>