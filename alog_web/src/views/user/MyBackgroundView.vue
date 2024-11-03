<template>
    <div class="container" style="margin-top: 5vh;">

        <div class="show-padding">

                        <br>
                        
                        <v-md-preview :text="text" ref="preview" />

                        

                        <div style="text-align: center;">
                            

                        </div>
                    </div>
    </div>
</template>


<script>

import $ from 'jquery'

// import router from '@/router';
// import { useStore } from "vuex";
import store from '@/store';
// import { ref } from 'vue';
export default {
    name: "showMePage",
    components: {
        // BlankBG,
    },
    data() {
        return {
            text: "![zuoguai240313iLU.png](https://img.alowlife.com/i/3/2024/03/13/zuoguai240313iLU-3.webp)",
        };
    },

    watch: {
        $route: {
            immediate: true,
            handler() {
                
                    
                    this.show_article(94)
                    
                
            }
        },
    },


    mounted() {
        this.screenWidth = window.innerWidth;
        // console.log("屏幕宽度：" + this.screenWidth);


        //版权信息
        document.oncopy = function addLink(e) {
            e.preventDefault();
            //获取复制的文本内容
            const selection = window.getSelection();



            //获取当前网页地址
            const localLink = document.location.href;


            var appendLink =  `\n\r
作者：` + store.state.user.username + `
链接：${localLink}
来源：`+store.state.BaseUrl+`
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 `
            

            let copyText = selection;
            if (copyText.toString().length >= 20) {
                copyText = selection + appendLink;
            }
            if (e.clipboardData) {
                e.clipboardData.setData('text', copyText)
            } else {
                //在ie中 clipboardData是window的属性
                window.clipboardData.setData('text', copyText);
            }
        }
    },

    methods:{
        show_article(articleId) {
            const that = this
            $.ajax({
                url: store.state.BaseUrl + "/api/article/common/" + articleId,
                type: "GET",
                headers: {
                    "Authorization":  store.state.user.token,
                },
                success(resp) {

                    if (resp.code == 1) {

                        that.text = resp.content.content
                        // that.article = resp.content
                        // that.authorId = resp.content.authorId
                        // that.tittle = resp.content.tittle
                        // console.log(resp)

                        // setTimeout(that.test, 1000);
                    } else {
                        console.log("获取文章失败")
                    }


                },
                error(resp) {
                    console.log(resp)
                }

            })
        },
    }

}


</script>


<style scoped>

.show-padding {
    border: 1px solid green;
    background-color: white;
    /* border-radius: 0.1%; */
    min-height: 50vh;
    opacity: 0.99;
}


</style>