import { createRouter, createWebHistory } from 'vue-router'
import HomePageView from '../views/page/HomePageView.vue'
import ArcihvePageView from '../views/page/ArchivePageView.vue'

import LoginView from '../views/user/LoginView.vue'
import MyBackgroundView from '../views/user/MyBackgroundView.vue'

import ArticleView from '../views/article/ArticleShowView.vue'
import ArticleDraft from '../views/article/ArticleDraftView.vue'
import ArticleWrite from '../views/article/ArticleWriteView.vue'
import ArticleManage from '../views/article/ArticleManageView.vue'
import ArticleModify from '../views/article/ArticleModifyView.vue'




const routes = [


  //文章部分
  {
    path: '/article',
    name: 'article',
    component: ArticleView,
    meta: {
      requestAuth: false,
      requestAdmin: false,

    }

  },
  {
    path: '/article-draft',
    name: 'article_draft',
    component: ArticleDraft,
    meta: {
      requestAuth: true,
      requestAdmin: false,

    }

  },
  {
    path: '/article-write',
    name: 'article_write',
    component: ArticleWrite,
    meta: {
      requestAuth: true,
      requestAdmin: false,

    }

  },
  {
    path: '/article-manage',
    name: 'article_manage',
    component: ArticleManage,
    meta: {
      requestAuth: true,
      requestAdmin: false,

    }

  },
  {
    path: "/article-modify",
    name: 'article_modify',
    component: ArticleModify,
    meta: {
      requestAuth: true,
      requestAdmin: false,

    }
  },


  // //个人
  // {
  //   path: '/modifyProfile',
  //   name: 'modify_profile',
  //   component: ModifyProfileView,
  //   meta: {
  //     requestAuth: true,
  //     requestAdmin: false,

  //   }

  // },

  
  {
    path: '/login',
    name: 'login',
    component: LoginView,
    meta: {
      requestAuth: false,
      requestAdmin: false,
    }

  },

  {
    path: '/myBg',
    name: "myBg",
    component: MyBackgroundView,
    meta: {
      requestAuth: false,
      requestAdmin: false,

    }
  }
  ,
  {
    path: '/404',
    name: 'error',
    component: () => import("../views/page/ErrorPageView.vue"),
    meta: {
      requestAuth: false,
      requestAdmin: false,

    }

  },


  {
    path: '/',
    name: 'home',
    component: HomePageView,
    meta: {
      requestAuth: false,
      requestAdmin: false,

    }

  },
  {
    path: '/archive',
    name: 'archive',
    component: ArcihvePageView,
    meta: {
      requestAuth: false,
      requestAdmin: false,

    }

  },

  // {
  //   path: '/about',
  //   name: 'about',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  // }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router



// router.beforeEach((to, from, next) => {
//   if (to.name == "register" || to.name == "login") {
//     next();
//   } else if (!store.state.user.is_login) {
//     next({ name: "login" })
//   } else {
//     next();
//   }
// })

import store from '@/store/index'
router.beforeEach((to, from, next) => {
  if (to.meta.requestAuth) {
    if (!store.state.user.is_login) {
      next({ name: "login" })
    } else if (!to.meta.requestAdmin || (to.meta.requestAdmin && store.state.user.is_admin)) {
      next();
    } else {
      next({ name: "home" });
    }
  }
  else {
    next();
  }
})

