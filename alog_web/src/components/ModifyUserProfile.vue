<template>
    <BlankBG v-if="screenWidth >= 600">

        <div class="row" style="padding-top: 5vh;">
            <div class="col-12">
                <div>
                    <!--sm.mstoken: YmN3mSDtPZZG8g0gtqFom4LQtj1cuYYY -->
                    <!-- 头像 -->


                    <div class="mb-3 row">
                        <label for="formFile" class="col-sm-2 col-form-label">头像</label>
                        <div class="col-sm-8">
                            <input class="form-control" type="file"
                                accept="image/gif,image/jpeg,image/jpg,image/png,image/svg" placeholder="选择文件"
                                id="avatarFile">
                        </div>


                        <div class="col-sm-2">
                            <div v-if="isLoadingAvatar" class="spinner-border text-primary col-sm-2" role="status">
                                <span class="visually-hidden">Loading...</span>
                                Loading...
                            </div>
                            <button v-else class="btn btn-warning" @click="update_avatar"
                                style="display:inline;">修改</button>
                        </div>
                    </div>
                    <!-- 背景 -->
                    <div class="mb-3 row">
                        <label for="formFile" class="col-sm-2 col-form-label">全局背景</label>
                        <div class="col-sm-8">
                            <input class="form-control" type="file"
                                accept="image/gif,image/jpeg,image/jpg,image/png,image/svg" placeholder="选择文件" id="bgFile">
                        </div>
                        <div class="col-sm-2">
                            <div v-if="isLoadingBg" class="spinner-border text-primary col-sm-2" role="status">
                                <span class="visually-hidden">Loading...</span>
                                Loading...

                            </div>
                            <button v-else class="btn btn-warning" @click="update_bg" style="display:inline;">修改</button>
                        </div>
                    </div>
                    <!-- 用户名 -->
                    <div class="mb-3 row">
                        <label for="inputUsername" class="col-sm-2 col-form-label">用户名：</label>
                        <div class="col-sm-8">
                            <input type="text" v-model="username" class="form-control" :placeholder="user.username"
                                id="inputUsername">
                        </div>
                        <div class="col-sm-2">
                            <div v-if="isLoadingUsername" class="spinner-border text-primary col-sm-2" role="status">
                                <span class="visually-hidden">Loading...</span>
                            </div>
                            <button v-else @click="update_username" class="btn btn-warning"
                                style="display:inline;">修改</button>
                        </div>
                    </div>
                    <!-- 密码 -->
                    <div class="mb-3 row">
                        <label for="inputPassword1" class="col-sm-2 col-form-label">密码：</label>
                        <div class="col-sm-8">
                            <input v-model="password" type="password" class="form-control" placeholder="******"
                                id="inputPassword1">
                        </div>
                        <div class="col-sm-2">
                            <div v-if="isLoadingPassword" class="spinner-border text-primary col-sm-2" role="status">
                                <span class="visually-hidden">Loading...</span>
                            </div>
                            <button v-else @click="update_password" class="btn btn-warning"
                                style="display:inline;">修改</button>
                        </div>
                    </div>
                    <!-- 重复密码 -->
                    <div class="mb-3 row">
                        <label for="inputPassword2" class="col-sm-2 col-form-label">重复密码：</label>
                        <div class="col-sm-8">
                            <input v-model="repeadPassword" type="password" class="form-control" placeholder="******"
                                id="inputPassword2">
                        </div>
                        <div class="col-sm-2">
                            <div v-if="isLoadingPassword" class="spinner-border text-primary col-sm-2" role="status">
                                <span class="visually-hidden">Loading...</span>
                            </div>
                            <button v-else @click="update_password" class="btn btn-warning">修改</button>
                        </div>

                    </div>
                    <!-- 个性签名 -->
                    <div class="mb-3 row">
                        <label for="exampleFormControlTextarea1" class="col-sm-2 col-form-label">个性签名：</label>
                        <div class="col-sm-8">
                            <textarea v-model="proverb" class="form-control" :placeholder="user.proverb"
                                id="exampleFormControlTextarea1" rows="3"></textarea>
                        </div>
                        <div class="col-sm-2">
                            <div v-if="isLoadingProverb" class="spinner-border text-primary col-sm-2" role="status">
                                <span class="visually-hidden">Loading...</span>
                            </div>
                            <button v-else @click="update_proverb" class="btn btn-warning">修改</button>
                        </div>

                    </div>
                </div>

            </div>



        </div>

    </BlankBG>
    <!-- shou+++++++++++++++++++++++++++++++++++手机端+++++++++++++++++++++++++++ -->
    <div v-else style="padding: 1vw;">

        <div>
            <!--sm.mstoken: YmN3mSDtPZZG8g0gtqFom4LQtj1cuYYY -->
            <!-- 头像 -->
            <div class="row">
                <label for="formFile" class="col-sm-2 col-form-label">头像</label>
                <div class="col-9">
                    <input class="form-control" type="file" accept="image/gif,image/jpeg,image/jpg,image/png,image/svg"
                        placeholder="选择文件" id="avatarFile">
                </div>


                <div class="col-3">
                    <div v-if="isLoadingAvatar" class="spinner-border text-primary col-sm-2" role="status">
                        <span class="visually-hidden">Loading...</span>
                        Loading...
                    </div>
                    <button v-else class="btn btn-warning" @click="update_avatar" style="display:inline;">修改</button>
                </div>
            </div>
            <!-- 背景 -->
            <div class="mb-3 row">
                <label for="formFile" class="col-sm-2 col-form-label">全局背景</label>
                <div class="col-9">
                    <input class="form-control" type="file" accept="image/gif,image/jpeg,image/jpg,image/png,image/svg"
                        placeholder="选择文件" id="bgFile">
                </div>
                <div class="col-3">
                    <div v-if="isLoadingBg" class="spinner-border text-primary col-sm-2" role="status">
                        <span class="visually-hidden">Loading...</span>
                        Loading...

                    </div>
                    <button v-else class="btn btn-warning" @click="update_bg" style="display:inline;">修改</button>
                </div>
            </div>
            <!-- 用户名 -->
            <div class="mb-3 row">
                <label for="inputUsername" class="col-sm-2 col-form-label">用户名：</label>
                <div class="col-9">
                    <input type="text" v-model="username" class="form-control" :placeholder="user.username"
                        id="inputUsername">
                </div>
                <div class="col-3">
                    <div v-if="isLoadingUsername" class="spinner-border text-primary col-sm-2" role="status">
                        <span class="visually-hidden">Loading...</span>
                    </div>
                    <button v-else @click="update_username" class="btn btn-warning" style="display:inline;">修改</button>
                </div>
            </div>
            <!-- 密码 -->
            <div class="mb-3 row">
                <label for="inputPassword1" class="col-sm-2 col-form-label">密码：</label>
                <div class="col-9">
                    <input v-model="password" type="password" class="form-control" placeholder="******" id="inputPassword1">
                </div>
                <div class="col-3">
                    <div v-if="isLoadingPassword" class="spinner-border text-primary col-sm-2" role="status">
                        <span class="visually-hidden">Loading...</span>
                    </div>
                    <button v-else @click="update_password" class="btn btn-warning" style="display:inline;">修改</button>
                </div>
            </div>
            <!-- 重复密码 -->
            <div class="mb-3 row">
                <label for="inputPassword2" class="col-sm-2 col-form-label">重复密码：</label>
                <div class="col-9">
                    <input v-model="repeadPassword" type="password" class="form-control" placeholder="******"
                        id="inputPassword2">
                </div>
                <div class="col-3">
                    <div v-if="isLoadingPassword" class="spinner-border text-primary col-sm-2" role="status">
                        <span class="visually-hidden">Loading...</span>
                    </div>
                    <button v-else @click="update_password" class="btn btn-warning">修改</button>
                </div>

            </div>
            <!-- 个性签名 -->
            <div class="mb-3 row">
                <label for="exampleFormControlTextarea1" class="col-sm-2 col-form-label">个性签名：</label>
                <div class="col-9">
                    <textarea v-model="proverb" class="form-control" :placeholder="user.proverb"
                        id="exampleFormControlTextarea1" rows="3"></textarea>
                </div>
                <div class="col-3">
                    <div v-if="isLoadingProverb" class="spinner-border text-primary col-sm-2" role="status">
                        <span class="visually-hidden">Loading...</span>
                    </div>
                    <button v-else @click="update_proverb" class="btn btn-warning">修改</button>
                </div>

            </div>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
        </div>



    </div>
</template>

<script>
import BlankBG from "@/components/BlankBG.vue"
import $ from "jquery"
// import { ref } from "vue";
// import router from "@/router";
import store from "@/store";
// import axios from "axios";

export default {
    name: "ModifyUserProfile",
    components: {
        BlankBG,
    },
    props: {
        user: Object
    },
    data() {
        return {
            username: "",
            password: "",
            repeadPassword: "",
            proverb: "",

            avatarFile: "",

            isLoadingAvatar: false,
            isLoadingBg: false,
            isLoadingUsername: false,
            isLoadingPassword: false,
            isLoadingProverb: false,
            // isLoadingbg: false,
            drawingBedToken: "",
            screenWidth: 1024,

        }
    },
    mounted() {
        this.get_drawing_bed_token()
        this.screenWidth = window.innerWidth


        // this.username = this.user.username
        // this.password = this.user.password
        // this.proverb = this.user.proverb


    },
    methods: {

        update_username() {
            this.isLoadingUsername = true;

            const that = this
            if (that.username == null || that.username.length < 2 || that.username.length > 10) {
                alert("用户名填写有误，需2-10字符")
                that.isLoadingUsername = false;
                return;
            }

            $.ajax({
                url: store.state.BaseUrl + "/api/user/update/",
                type: "put",
                headers: {
                    Authorization: "Bearer " + store.state.user.token,
                },
                data: {
                    username: that.username,
                    userId: that.user.userId
                },
                success(resp) {
                    that.isLoadingUsername = false;


                    if (resp.code == 1) {

                        store.dispatch("getInfo", {
                            success() {
                                // console.log("拉取用户信息成功"),
                                console.log(resp)
                                that.username = ""
                                alert("修改用户名成功")
                            },
                            error() {
                                alert("修改用户名失败,可能已存在用户" + that.username)
                            }
                        })

                    } else {
                        alert("修改用户名失败,可能已存在用户" + that.username)

                    }
                },
                error() {
                    that.isLoadingUsername = false;


                    alert("修改用户名失败,可能已存在用户" + that.username)

                }
            })

        },

        update_password() {
            this.isLoadingPassword = true;

            const that = this

            if (that.password != that.repeadPassword) {
                alert("两次密码不相符")
                that.isLoadingPassword = false;

                return;
            }

            if (that.password == null || that.password.length < 6 || that.password.length > 20) {
                that.isLoadingPassword = false;

                alert("密码填写有误，需6-20字符")
                return;
            }

            $.ajax({

                url: store.state.BaseUrl + "/api/user/update/",
                type: "put",
                headers: {
                    Authorization: "Bearer " + store.state.user.token,
                },
                data: {
                    password: that.password,
                    userId: that.user.userId
                },
                success(resp) {
                    that.isLoadingPassword = false;


                    if (resp.code == 1) {

                        store.dispatch("getInfo", {
                            success() {
                                // console.log("拉取用户信息成功"),
                                console.log(resp)
                                that.password = ""
                                that.repeadPassword = ""
                                alert("修改密码成功")
                            },
                            error() {
                                alert("修改密码失败")
                            }
                        })

                    } else {
                        console.log(resp)
                    }
                }, error() {
                    that.isLoadingPassword = false;


                    alert("修改密码失败")

                }
            })
        },
        update_proverb() {
            this.isLoadingProverb = true;

            const that = this
            if (that.proverb == null || that.proverb.length < 2 || that.username.length > 100) {
                that.isLoadingProverb = false;

                alert("个性签名填写有误，需2-100字符")
                return;
            }

            $.ajax({

                url: store.state.BaseUrl + "/api/user/update/",
                type: "put",
                headers: {
                    Authorization: "Bearer " + store.state.user.token,
                },
                data: {
                    proverb: that.proverb,
                    userId: that.user.userId
                },
                success(resp) {
                    that.isLoadingProverb = false;


                    if (resp.code == 1) {

                        store.dispatch("getInfo", {
                            success() {
                                // console.log("拉取用户信息成功"),
                                console.log(resp)
                                that.proverb = ""
                                alert("修改个性签名成功")
                            },
                            error() {
                                alert("修改个性签名失败")
                            }
                        })

                    } else {
                        alert("修改个性签名失败")

                    }
                },
                error() {
                    that.isLoadingProverb = false;


                    alert("修改个性签名失败")

                }
            })
        },
        update_avatar() {
            this.isLoadingAvatar = true;
            const file = document.getElementById("avatarFile");
            if (file == null) {
                this.isLoadingAvatar = false;
                return;
            }
            const fileOb = file.files[0];
            if (fileOb == null) {
                this.isLoadingAvatar = false;
                return
            }
            if (fileOb.length > 2 * 1024 * 1024) {
                this.isLoadingAvatar = false;
                alert("文件过大")
                return
            }
            const params = new FormData();
            params.append('image', fileOb)
            params.append("token", this.drawingBedToken)
            // console.log("token" + this.drawingBedToken)
            this.upload_image(params, true)
            // $.ajax({

            //     url: store.state.BaseUrl + "/api/user/avatar/",
            //     type: "put",
            //     headers: {
            //         Authorization: "Bearer " + store.state.user.token,
            //     },
            //     data: params,
            //     processData: false,
            //     contentType: false,
            //     success(resp) {
            //         that.isLoadingAvatar = false;

            //         if (resp.code == 1) {

            //             store.dispatch("getInfo", {
            //                 success() {
            //                     console.log(resp)
            //                     alert("修改头像成功")
            //                 },
            //                 error() {
            //                     alert("修改头像成功，请刷新页面")
            //                 }
            //             })

            //         } else {
            //             alert("修改头像失败，请稍后再试")
            //             console.log(resp.content)
            //         }
            //     },
            //     error() {
            //         that.isLoadingAvatar = false;

            //         alert("修改头像失败，请检查网络连接")

            //     }
            // })
        }
        ,
        update_bg() {
            this.isLoadingBg = true;
            const file = document.getElementById("bgFile");

            if (file == null) {
                this.isLoadingBg = false;

                return;
            }
            const fileOb = file.files[0];
            // console.log(fileOb)

            if (fileOb == null) {
                this.isLoadingBg = false;

                return
            }
            if (fileOb.length > 5 * 1024 * 1024) {
                this.isLoadingBg = false;

                alert("文件过大")
                return
            }

            const params = new FormData();
            params.append('image', fileOb)
            params.append("token", this.drawingBedToken)
            this.upload_image(params, false)
        },

        get_drawing_bed_token() {
            const that = this
            $.ajax({

                url: store.state.BaseUrl + "/api/user/image/token/",
                type: "post",
                headers: {
                    Authorization: "Bearer " + store.state.user.token,
                },
                data: {
                },
                success(resp) {

                    if (resp.code == 1) {
                        that.drawingBedToken = resp.content.token
                    }
                }

            })
        },

        upload_bgUrl(bgUrl) {
            const that = this
            $.ajax({

                url: store.state.BaseUrl + "/api/user/bg/updateUrl/",
                type: "put",
                headers: {
                    Authorization: "Bearer " + store.state.user.token,
                    // 'Content-Type': 'multipart/form-data'
                },
                data: {
                    bg: bgUrl,
                    userId: that.user.userId
                },
                success(resp) {
                    that.isLoadingBg = false;

                    if (resp.code == 1) {

                        store.dispatch("getInfo", {
                            success() {
                                // console.log("拉取用户信息成功"),
                                // console.log(resp)
                                alert("修改全局背景成功")
                            },
                            error() {
                                alert("修改背景成功，请刷新页面")
                            }
                        })

                    } else {
                        alert(resp.content)
                    }
                },
                error() {
                    that.isLoadingBg = false;

                    alert("修改背景失败")

                }
            })
        },


        upload_image(params, isAvatar) {
            const that = this
            $.ajax({
                url: "https://img.alowlife.com/api/index.php",
                type: "post",
                async: true,
                headers: {
                },
                data: params,
                mimeType: 'multipart/form-data',
                crossDomain: true,
                processData: false,
                contentType: false,
                success(resp) {
                    // console.log(resp)
                    resp = $.parseJSON(resp)
                    if (resp.code == 200) {
                        if (isAvatar) {
                            that.upload_avatarUrl(resp.url)

                        } else {
                            that.upload_bgUrl(resp.url)
                        }
                    } else {
                        alert("修改头像失败，请稍后再试")
                        // console.log(resp)
                    }
                },
                error() {
                    that.isLoadingAvatar = false;
                    that.isLoadingBg = false;
                    alert("修改头像失败，请检查网络连接")

                }

            })
        },
        upload_avatarUrl(avatarUrl) {
            const that = this
            $.ajax({

                url: store.state.BaseUrl + "/api/user/avatar/updateUrl/",
                type: "put",
                headers: {
                    Authorization: "Bearer " + store.state.user.token,
                    // 'Content-Type': 'multipart/form-data'
                },
                data: {
                    avatar: avatarUrl,
                    userId: that.user.userId
                },
                success(resp) {
                    that.isLoadingAvatar = false;

                    if (resp.code == 1) {

                        store.dispatch("getInfo", {
                            success() {
                                // console.log("拉取用户信息成功"),
                                // console.log(resp)
                                alert("修改头像成功")
                            },
                            error() {
                                alert("修改头像成功，请刷新页面")
                            }
                        })

                    } else {
                        alert(resp.content)
                    }
                },
                error() {
                    that.isLoadingAvatar = false;

                    alert("修改头像失败")

                }
            })
        },




    }


}

</script>


<style scoped></style>