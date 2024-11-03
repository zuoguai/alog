import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';
import Prism from 'prismjs'

import VMdPreview from '@kangc/v-md-editor/lib/preview';
import '@kangc/v-md-editor/lib/style/preview.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';

import createLineNumbertPlugin from '@kangc/v-md-editor/lib/plugins/line-number/index';
import createKatexPlugin from '@kangc/v-md-editor/lib/plugins/katex/cdn';

import createEmojiPlugin from '@kangc/v-md-editor/lib/plugins/emoji/index';
import '@kangc/v-md-editor/lib/plugins/emoji/emoji.css';

import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index';
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css';

import createTodoListPlugin from '@kangc/v-md-editor/lib/plugins/todo-list/index';
import '@kangc/v-md-editor/lib/plugins/todo-list/todo-list.css';
// highlightjs
import hljs from 'highlight.js';
// import 'katex/dist/katex.min.css'
// import 'katex/dist/katex.min.js'

//鼠标连线
import VueParticles from 'vue-particles'



VMdPreview.use(githubTheme, {
  Hljs: hljs,
  // extend(md) {

  // }
})
VMdPreview.use(createKatexPlugin())
VMdPreview.use(createCopyCodePlugin())
VMdPreview.use(createTodoListPlugin())
VMdPreview.use(createEmojiPlugin())

VueMarkdownEditor.use(vuepressTheme, {
  Prism,
  // extend(md) {

  // }
}).use(createKatexPlugin()).use(createEmojiPlugin()).use(createCopyCodePlugin()).use(createTodoListPlugin());

// Vue.prototype.BaseUrl = "https://alowlife.com";

createApp(App).use(store).use(router).use(VueMarkdownEditor).use(VMdPreview).use(createLineNumbertPlugin).use(VueParticles)
  .mount('#app')
