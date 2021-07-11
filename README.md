# 简介

互联网访问工具分享管理后台

# 配置

请参考 [./config](./config) 下的配置文件

# 使用

```sh
docker run -d --restart always -v $PWD/v2ray:/app/data -p 3005:3005 shynome/cobweb:1.3.0
```

管理界面: http://127.0.0.1:3005/admin/info/v2ray_users
默认账号: admin
默认密码: admin

# 进阶使用

```sh
docker run -d --restart always \
  -v $PWD/v2ray:/app/data \
  -p 3005:3005 \
  -e LISTEN=":3005" \ # 监听的地址
  -e DB="/app/data/cobweb.db" \ # 数据库地址
  -e V2RAY_URL="/ray" \ # v2ray ws path
  -e ADMIN="admin" \ # 管理后台地址, 最前面不用加 `/`, go-admin 会加
  # 下面是分享链接的配置
  -e USE_DOMAIN="example.com" \ # ip 也可以
  -e USE_PORT="443" \
  -e USE_TLS="tls" \
  -e USE_REMARK_PREFIX="中转" \
  -e USE_PATH="/ray" \ # 默认是 V2RAY_URL , 但有反向代理的话可能不是这个地址
  shynome/cobweb:1.3.0
```

# 构建

```sh
make build
# 或者
go build . -o build/cobweb
```

# 运行

```sh
./build/cobweb
```
