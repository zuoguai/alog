<template>

  <div  @mousemove="handleMouseMove"  class="is_login-div">
    <div>
      <img style="object-fit: cover" src="https://img.alowlife.com/i/2024/11/02/ne3fgg.jpg" class="img-window-background">
    </div>

    <div
      style="z-index: -9;  opacity: 1; min-height: 100vh; background-color: rgb(0,0,0,0.2);  font-family: 宋体; font-weight:800;'">
      <nav>
        <NavBar :mouseX="mouseXX" :mouseY="mouseYY" v-if="isRouterAlive" />
        <!-- <NavBar class="sticky-top"></NavBar> -->
      </nav>
      <router-view  v-if="isRouterAlive" />
      <br>
    </div>

  </div>
  <div v-if="screenWidth >= 800" class="footer">

    <span>
      <!-- <kbd style="font-size: smaller; color:red "> -->
        © 2024
      <!-- </kbd> -->
    </span>
    &nbsp;
    <span>
        <!-- <kbd style="font-size: smaller;  color:skyblue"> -->
          作怪
        <!-- </kbd> -->
    </span>
    &nbsp;
    <span>
      <kbd style="font-size: smaller; color:greenyellow">
        <a href="mailto:zuoguai2023@foxmail.com" style="color:greenyellow; text-decoration-line: none">zuoguai2023@foxmail.com</a>
      </kbd>
    </span>
    &nbsp;
    <span>
      <!-- <kbd style="font-size: smaller;color:pink "> -->
        运行：{{ run_day }}
      <!-- </kbd> -->
    </span>
    <!-- <span>作者：</span>
    <span>作者：</span> -->

  </div>
</template>

<script>

import NavBar from "@/components/NavBar.vue"
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap/dist/js/bootstrap.min"

import { ref, nextTick, provide } from "vue"


export default {
  name: "App",
  components: {
    NavBar
  },
  data() {
        return {
            mouseXX:'100',
            mouseYY:'100'
        }
    },
    methods: {

        alert_msg(msg) {
            alert(msg)
        },

        handleMouseMove(event) {
      this.mouseXX = (parseInt(event.clientX * Math.random() % 101)).toString()
      this.mouseYY =(parseInt(event.clientY* Math.random() % 101)).toString()
    }},
  setup() {
    // 局部组件刷新
    const screenWidth = ref("")
    screenWidth.value = 1024;
    screenWidth.value = window.innerWidth
    const isRouterAlive = ref(true);
    const reload = () => {
      isRouterAlive.value = false;
      nextTick(() => {
        isRouterAlive.value = true;
      });
    };
    provide("reload", reload);

    let now_time = ref("")
    let birth_day = 1730605114
    let run_day = ref("")
    setInterval(() => {
      now_time.value = Math.round(new Date().getTime() / 1000)
      run_day.value = cal_length((now_time.value - birth_day) * 1000)
    }, 10)



    const cal_length = (len) => {

      let res = ""
      let seconds = (len / 1000)

      let minutes = Math.floor(seconds / 60)
      let hours = Math.floor(minutes / 60)
      let days = Math.floor(hours / 24)

      minutes = minutes - hours * 60
      seconds = seconds - hours * 60 * 60 - minutes * 60
      hours = hours - days * 24

      if (days >= 1) {
        res += days + " 天 "
      }

      if (hours >= 1) {
        res += hours + " 小时 "
      }
      if (minutes >= 1) {
        res += " " + minutes + " 分"
      }

      res += " " + seconds.toFixed(0) + " 秒"

      if (res == "") {
        res = " 0 秒"
      }
      return res

    }



    return {
      isRouterAlive,
      screenWidth,
      run_day,
    };
  },

}


</script>

<style scoped>
.footer {
  text-align: center;
  /* font-family: 宋体; */
  /* font-weight: 10; */
  font-size: 15px;
  background-color: rgba(0, 0, 0, 0.5);
  color: rgb(255, 255, 255);

}

.footer:hover {
  /* background-color: rgba(255, 255, 255, 0.1); */
}

nav a.router-link-exact-active {
  color: #42b983;
}

.is_login-div {
  padding-bottom: 0vh;
  /* background-color: #f7f8ec; */
  /* position: relative; */
}

.img-window-background {
  z-index: -10;
  position: fixed;
  width: 100%;
  height: auto;
  /* background-repeat: repeat-y; */
  overflow: hidden;
}

.img-phone-background {
  z-index: -10;
  position: fixed;
  width: auto;
  height: 100%;
  background: repeat;
  overflow: hidden;
}
</style>
