# GoAdmin 介绍

GoAdmin 是一个帮你快速搭建数据可视化管理应用平台的框架。 

- [github](https://github.com/GoAdminGroup/go-admin)
- [论坛](http://discuss.go-admin.com)
- [文档](https://book.go-admin.cn)

## 目录介绍

```
.
├── Dockerfile          Dockerfile
├── Makefile            Makefile
├── adm.ini             adm配置文件
├── admin.db            sqlite数据库
├── build               二进制构建目标文件夹
├── config.json         配置文件
├── go.mod              go.mod
├── go.sum              go.sum
├── html                前端html文件
├── logs                日志
├── main.go             main.go
├── main_test.go        CI测试
├── pages               页面控制器
├── tables              数据模型
└── uploads             上传文件夹
```

## 生成CRUD数据模型

### 在线工具

管理员身份运行后，访问：http://127.0.0.1:80/admin/info/generate/new

### 使用命令行工具

```
adm generate -l cn -c adm.ini
```

