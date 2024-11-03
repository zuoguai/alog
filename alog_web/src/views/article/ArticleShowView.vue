<template>
        <div id="mouse-line" style="opacity: 0.99; z-index: -1; position: fixed;">
            <vue-particles 
      color="#dedede"
      :particleOpacity="0.7"
      :particlesNumber="60"
      shapeType="circle"
      :particleSize="4"
      linesColor="#dedede"
      :linesWidth="1"
      :lineLinked="true"
      :lineOpacity="0.5"
      :linesDistance="150"
      :moveSpeed="3"
      :hoverEffect="true"
      hoverMode="grab"
      :clickEffect="true"
      clickMode="push"
    >
    
    </vue-particles>
        </div>
    <div class="myContainer "  id="wayBackTop" v-if="screenWidth >= 1000" style="background-color:  rgba(0, 0, 0, 0.5); min-height: 60vh;">

        <div class="container" >
            <div class="row">

                <div class="col-9 ">
                    <div class="article-show-padding">

                        <h2 style="padding-left: 2vw; padding-top: 3vh;  text-overflow:ellipsis; overflow: hidden; " ><b>{{ title }}</b></h2>

                        <br>
                        <hr class="hr-solid-content" :data-content="'作者：' + author + ' 阅读：' + article.view + ' 上次修改：' + timestampToTime(article.update_time*1000)">

                                <v-md-preview :text="text"  ref="preview" />

                        <hr class="hr-solid-content" :data-content="timestampToTime(article.create_time*1000)">

                        <div style="text-align: center;">
                           
                                作者：<kbd>{{ $store.state.user.username }} </kbd>&nbsp;
                                阅读：<mark>{{ article.view }}</mark>


                        </div>

                    </div>
                </div>
                <div class="col-3">
                    <div class="article-side-padding" style="opacity: 0.99; ">
                       
                        <p style="text-align: center; color:white">目录导航</p>

                        <div class="right-toc-padding" style="position: sticky; top: 10px; ">

                            <div v-for="anchor in titles" :key="anchor.id"
                                :style="{ padding: `1px 0 1px ${anchor.indent * 20}px` }"
                                @click="handleAnchorClick(anchor)">
                                <a href="#" style="font-size: medium; ">
                                    <kbd>
                                        {{ anchor.title }}
                                    </kbd>
                                </a>
                            </div>
                            <a style="text-decoration: underline; font-size: large;" href="#wayBackTop"><code>
                                    <kbd>回到顶部</kbd></code></a>
                        </div>
                        <!-- <hr class="hr-wavy"> -->

                        

                        


                        
                        
                    </div>

                </div>

            </div>
        </div>
    </div>


    <!-- 手机端============================================================== -->

    <div class="myContainer " id="wayBackTop" v-else style="min-height: 1027px;">
        <a style="text-decoration: underline; font-size: large; 
    z-index: 100; float: right; padding-right: 5vw; position: sticky; top: 93vh; scroll-behavior: smooth"
            href="#wayBackTop">
            <!-- <mark style="background-color: white; border-radius: 10%;">顶部</mark> -->
            <img src="/ico/up.png" style="height: 30px;">
        </a>
        <div style="padding-left: 5vw;  max-height: 200px;  ">
            标题：<kbd>{{ title }}</kbd>
        </div>
        <hr class="hr-wavy">

        <div style="padding-left: 5vw;  max-height: 200px;  overflow-y: auto;">
            <!-- <div style="background-color: rgba(255, 255, 255, 0.5);"> -->

            目录：
            <!-- <div class="card"> -->
            <!-- <div class="card-body"> -->
            <!-- <p style="text-align: center;">目录导航</p> -->

            <div v-for="anchor in titles" :key="anchor.id" :style="{ padding: `1px 0 1px ${anchor.indent * 20}px` }"
                @click="handleAnchorClick(anchor)">
                <a href="#" style="font-size: medium; padding: 10vw;"><kbd>{{ anchor.title }}</kbd></a>
            </div>
            <a style="text-decoration: underline; font-size: large; padding: 10vw;"
                href="#wayBackTop"><code>回到顶部</code></a>
            <!-- </div> -->
            <!-- </div> -->
            <!-- </div> -->
        </div>
        <hr class="hr-wavy">
        <div style="padding-left: 5vw; padding-right: 10vw;">
           
        </div>
        <hr class="hr-wavy">

        <div class="" style="padding: 1vw; opacity: 0.99;">
            <!-- <div class="card">
                <div class="card-body"> -->
            <div style="background-color: white; overflow-x:scroll;">


                <div style="width: 100%; word-break: break-all; ">
                    <h2 style="padding-left: 2vw; padding-top: 3vh;"><b>{{ title }}</b></h2>

                    <hr class="hr-solid-content" :data-content="'作者：' + author + ' 阅读：' + article.view">

                    <div style="font-size: smaller;">
                        <v-md-preview style="font-size: smaller;" :text="text" ref="preview" />
                    </div>
                    <div style="text-align: center;">
                        
                            作者：<kbd>{{ author }} </kbd>&nbsp;
                            阅读：<mark>{{ article.view }}</mark>


                    </div>

                    <hr class="hr-solid-content" :data-content="'发布：' + timestampToTime(article.create_time * 1000)">
                    <hr class="hr-solid-content" :data-content="'上次修改：' + timestampToTime(article.update_time * 1000)">

                </div>
            </div>
            <!-- </div>
            </div> -->
        </div>

        <hr class="hr-wavy">

        <div class="" style="padding-left: 5vw; padding-right: 5vw;">
 

            
            <hr class="hr-wavy">

            <!-- <div class="card">
                <div class="card-body"> -->
           
            <!-- </div>
            </div> -->


        </div>

    </div>
</template>

<script>

import $ from 'jquery'

import router from '@/router';
import store from '@/store';

export default {
    name: "ArticePage",
    components: {
        // BlankBG,
    },
    data() {
        return {
            text: "![zuoguai240314zkR.png](https://img.alowlife.com/i/3/2024/03/14/zuoguai240314zkR-3.webp)",
            titles: "",
            article: "",
            author: "作怪",
            title: "糟糕，页面未找到",

            authorId: "",
            articleId: "",


            screenWidth: 1024,

        };
    },

    watch: {
        $route: {
            immediate: true,
            handler() {
                if (this.$route.query.article_id != this.articleId) {//需要监听的参数
                    this.article_id = this.$route.query.article_id
                    this.show_article(this.article_id)
                    this.loadAuthorInfo(this.authorId)
                }
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


            var appendLink = `\n\r
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
    methods: {

        timestampToTime(timestamp) {
            timestamp = timestamp ? timestamp : null;
            let date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
            let Y = date.getFullYear() + '-';
            let M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
            let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' ';
            let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
            let m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
            let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
            return Y + M + D + h + m + s;
        },

        loadAuthorInfo() {
            const that = this
            that.author = "作怪"

        },

       


       

        goto_article(article_id) {

            if (this.article_id == article_id) {
                // console.log("当前文章展示")
                return
            }

            router.push({
                path: "/article",
                query: {
                    article_id: article_id
                }
            })
        },

        handleAnchorClick(anchor) {
            const { preview } = this.$refs;
            const { lineIndex } = anchor;

            const heading = preview.$el.querySelector(`[data-v-md-line="${lineIndex}"]`);

            if (heading) {
                preview.scrollToTarget({
                    target: heading,
                    scrollContainer: window,
                    top: 60,
                });
            }
        },

    
        show_article(article_id) {
            const that = this

                $.ajax({
                    url: store.state.BaseUrl + "/api/article/" + article_id,
                    type: "GET",
                    headers: {
                    "Authorization":  store.state.user.token,
                },
                    success(resp) {

                        if (resp.msg == "ok") {   
                            // console.log(resp.data.list[0])
                            resp = resp.data.list[0]
                            that.text = resp.content
                            that.article = resp
                            that.authorId = resp.user_id
                            that.title = resp.title



                            setTimeout(that.test, 1000);
                        } else {
                            // console.log("获取文章失败")
                            that.text = "![zuoguai240314zkR.png](https://img.alowlife.com/i/3/2024/03/14/zuoguai240314zkR-3.webp)"
                            that.title = "页面未找到"
                            // router.push({ name: "error" })
                        }


                    },
                    

                })
           

        },


        test() {

            const anchors = this.$refs.preview.$el.querySelectorAll('h1,h2,h3,h4,h5,h6');
            const titles = Array.from(anchors).filter((title) => !!title.innerText.trim());

            if (!titles || !titles.length) {
                this.titles = [];
                return;
            }

            const hTags = Array.from(new Set(titles.map((title) => title.tagName))).sort();

            this.titles = titles.map((el) => ({
                title: el.innerText,
                lineIndex: el.getAttribute('data-v-md-line'),
                indent: hTags.indexOf(el.tagName),
            }));
        }
    },

    setup() {





    }

}

</script>


<style scoped>

#mouse-line{
    width: 100%;
    height: 100%

}
kbd:hover {
    /* text-decoration: underline; */
    border: 2px solid rgb(255, 255, 255);
}

.myContainer {
    /* padding-top: 4vh; */
    /* padding-left: 4vw; */
    opacity: 1;

}

.left-toc-top-padding {
    /* background-color: white; */
    height: 15vh;
    width: 17vw;
    /* overflow-y: scroll; */
    padding-left: 1vw;
}


.left-toc-padding {
    /* border-radius: 2%; */

    height: 30vh;
    overflow-y: auto;
    width: 100%;
}

.article-show-padding {
    /* background-color: white; */
    /* background-color:  rgba(207, 207, 207, 0.2); */
    color: white;
    /* border-radius: 0.1%; */
    min-height: 50vh;
    opacity: 1;
}

.article-side-padding {
    /* background-color: rgba(0, 0, 0, 0.2); */
    position: sticky;
    top: 5vh;
}

.right-tag-padding {
    max-height: 20vh;
    overflow-y: auto;
    overflow: hidden;
    /* padding-left: 1vw; */
    padding-right: 1vw;
    /*隐藏滚动条但不取消滚动功能 */
    -ms-overflow-style: none;
    /* IE and Edge */
    scrollbar-width: none;
    /*火狐 */
    /* background-color: rgba(255, 255, 255, 0.5); */
    color: white;

    /* border: 1px solid green; */

}

.right-toc-padding {
    padding-right: 1vw;
    padding-left: 1vw;
    max-height: 80vh;

    overflow-y: auto;

    /* background-color: rgba(255, 255, 255, 0.5); */
    color: white;
    position: relative;

    /* border: 1px solid green; */
}


.right-article-padding {

    padding-right: 1vw;
    max-height: 25vh;
    overflow-y: auto;

    /* background-color: rgba(255, 255, 255, 0.5); */

    color: white;
    /* border: 1px solid green; */
}

.right-comment-padding {
    /* border-radius: 2%; */
    /* padding-right: 1vw; */
    padding-top: 100px;
    padding-left: 1vw;
    padding-bottom: 1vh;
    height: 25vh;
    overflow-y: auto;
    /* background-color: rgba(255, 255, 255, 0.5); */

    color: white;
    border: 1px solid green;
}


a {
    color: white;
    /*字体颜色*/
    font-size: small;
    font-weight: 1000;
    text-decoration: none;
    /*去除a标签下划线*/
    /* outline: none; */
    /*去掉a标签默认轮廓*/
}

/* p {
    font-weight: 50;
    color: black;
} */

.hr-solid-content {
    /* color: #a2a9b6; */
    color: black;
    border: 0;
    font-size: 12px;
    padding: 1em 0;
    position: relative;
}

.hr-solid-content::before {
    content: attr(data-content);
    position: absolute;
    padding: 0 1ch;
    line-height: 1px;
    border: solid #d0d0d5;
    border-width: 0 99vw;
    width: fit-content;
    /* for IE浏览器 */
    white-space: nowrap;
    left: 50%;
    transform: translateX(-50%);
}

.hr-solid-content::after {
    content: attr(data-content);
    position: absolute;
    padding: 4px 1ch;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    color: transparent;
    border: 1px solid #d0d0d5;
}

.hr-wavy {
    border: 0;
    padding: 3px;
    background: repeating-linear-gradient(135deg, black 0px, #0358f7 1px, transparent 1px, transparent 6px);
}


/* 禁止复制文章 */
/* div {
    user-select: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    -webkit-user-drag: none;
} */


.right-toc-padding::-webkit-scrollbar {
    /*滚动条整体样式*/
    width: 8px;
    /*高宽分别对应横竖滚动条的尺寸*/
    height: 1px;
}


.right-toc-padding::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    /* border-radius: 10px; */
    /* background: #9e9e9e; */
    background: grey;
    /* box-shadow: inset 0 0 5px rgba(0, 122, 204); */
    /* box-shadow: inset 0 0 5px black; */

}

/* 滚动条底层颜色! */
.right-toc-padding::-webkit-scrollbar-track {
    /* border-radius: 10px; */
    background: #ededed;

    /*滚动条里面轨道*/
    /* box-shadow: inset 0 0 5px rgba(0, 0, 0, .1); */
    box-shadow: inset 0 0 5px white;
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