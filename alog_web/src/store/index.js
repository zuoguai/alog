import { createStore } from 'vuex'
import ModuleUser from '@/store/user'



export default createStore({
  state: {
    //上线
    BaseUrl: "https://java.alowlife.com",


    //测试
    // BaseUrl: "http://localhost:16666",
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user: ModuleUser,
  }
})
