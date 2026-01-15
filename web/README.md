# 短视频网站前端

根据后端 proto 接口定义生成的短视频网站，完整适配所有API接口。

## 功能

- 用户：注册、登录、资料编辑
- 视频：流式播放、发布、分类浏览
- 互动：点赞、收藏、关注、评论、举报
- 个人：作品、喜欢、收藏列表

## 运行

```bash
cd web
go run server.go
# 访问 http://localhost:3000
```

## 配置

修改 `static/js/api.js` 中的 `BASE_URL` 指向网关地址。

## 文件

- `index.html` - 主页面
- `static/css/style.css` - 样式
- `static/js/api.js` - 接口封装
- `static/js/app.js` - 应用逻辑
