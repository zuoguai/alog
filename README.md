### 一、Golang + Vue 的极简前后端分析博客程序

后端仅提供：文章编写/修改/删除，获取文章列表

前端仅包含：首页，文章页，归档，登录。（编写文章使用Markdown语法）

[我用来写Java笔记的博客](https://java.alowlife.com)

### 二、使用

#### 前端：

```
npm install
```

```
npm run serve
```

```
npm run build
```

#### 后端

```
go mod tidy
```

```
go build
```

配备好数据库

 

### 三、后记

开发这个程序是想把一些笔记记录在网页博客，敦促自己学习

不直接使用市面上封装好的博客程序，是因为他们虽然功能丰富，但我自己没有这些需求

自己在使用Java开发的时候，发现只实现简单的文章管理，多用户，标签，评论等功能，在云服务器上居然能吃几百MB到G的内存。Java各种依赖实在臃肿。我的几百MB小弱机服务器根本扛不住/(ㄒoㄒ)/~~

于是改用Golang，花了一天多，只把文章部分写完

感觉已经够自己用了

就不再写标签/评论/用户/这些模块啦





TIPS：与Java相比，Golang运行起来仅仅吃了几MB内存。

我的弱鸡服务器：终于可以了  我可以了！！！😍😍😍










