# ziroom-reservation，公寓预订

# 项目环境变量
```
HTTP_LISTEN_ADDRESS=:3000
JWT_SECRET=somethingsupersecretthatNOBODYKNOWS
MONGO_DB_NAME=hotel-reservation
MONGO_DB_URL=mongodb://localhost:27017
MONGO_DB_URL_TEST=mongodb://localhost:27017
```

## 项目概要


- 用户-> 酒店预订房间
- 管理员-> 查询预订情况
- 身份验证和授权-> JWT 令牌
- 酒店-> CRUD API-> JSON
- 房间-> CRUD API-> JSON
- 脚本-> 数据库管理-> 随机初始话数据、迁移


## 资源
### Mongodb 驱动
文档
```
https://mongodb.com/docs/drivers/go/current/quick-start
```


### gin 
文档
```
https://gin-gonic.com/zh-cn/docs/quickstart/
```


## Docker

### 部署 Mongodb 容器

```bash

# 启动容器
docker run -d --name mongodb \
-p 27017:27017 \
-e MONGO_INITDB_ROOT_USERNAME=root \
-e MONGO_INITDB_ROOT_PASSWORD=123456 \
mongo

docker start mongodb
docker stop  mongodb

# 查看网关
docker inspect mongodb | grep IPAddress

# 宿主机操作
docker exec -it mongodb /bin/bash

# mongodb cli
mongosh -u root -p 123456 --authenticationDatabase admin


```
