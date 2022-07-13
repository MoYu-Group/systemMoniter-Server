# systemMoniter-Server
Golang 写的系统资源监控服务器端

开发项目，只是写着玩的，方便自用，请不要商用！

客户端正在开发中。

## 简要接口说明

以下所有错误均可在 `common.ErrorCode` 查到

### 登录

POST `http://localhost:8085/api/user/login`

表单数据如下：

```json
{

  "user": "node",
  "password": "123456"
}
```

正常返回结构如下，其中 token 为唯一令牌

```json
{
    "error": 0,
    "error_msg": "",
    "data": {
        "name": "node",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5vZGUiLCJleHAiOjE2NTc2OTg0MDMsImlhdCI6MTY1NzY4NzYwMywibmJmIjoxNjU3Njg3NjAzfQ.fHyo2px-u37qADdKzf2PmFj-7OV2z9z7mkgI1av3sWs",
        "userId": "e5d86c3c-c801-495f-bcd4-2edea6316b91"
    }
}
```

### 创建用户（需鉴权，如不需要修改路由文件跳过）

POST ` http://localhost:8085/api/user/register`

Header 头部含有有效 token

````
Authorization：eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5vZGUiLCJleHAiOjE2NTc2ODk1NDcsImlhdCI6MTY1NzY3ODc0NywibmJmIjoxNjU3Njc4NzQ3fQ.P1dJnpJL9ZYcRYRPhRcHJBPj5oBv8gi6WAAXXT1yC3E`
````

正常表单数据如下：

```json
{
    "user": "admin",
    "password": "123456"
}
```

返回如下：

```
{
    "error": 0,
    "error_msg": "",
    "data": {
        "name": "admin",
        "userId": "2be5aab2-9c7d-4189-9249-7f22b00b6d70"
    }
}
```

### 创建节点（需鉴权，如不需要修改路由文件跳过）

POST ` http://localhost:8085/api/node/register`

Header 头部含有有效 token

````
Authorization：eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5vZGUiLCJleHAiOjE2NTc2ODk1NDcsImlhdCI6MTY1NzY3ODc0NywibmJmIjoxNjU3Njc4NzQ3fQ.P1dJnpJL9ZYcRYRPhRcHJBPj5oBv8gi6WAAXXT1yC3E`
````

正常表单数据如下，均需填写，uid 为关联的用户 id：

```json
{
    "name": "local",
    "uid": "e5d86c3c-c801-495f-bcd4-2edea6316b91",
    "type": "local",
    "host": "127.0.0.1",
    "location": "local",
    "custom":"this is local test node"
}
```

返回如下：

```json
{
    "error": 0,
    "error_msg": "",
    "data": {
        "host": "127.0.0.1",
        "id": "0499fd55-3a26-43f7-aab9-4d6b267c13a0",
        "name": "local"
    }
}
```

### 创建节点状态数据（需鉴权，如不需要修改路由文件跳过）

POST ` http://localhost:8085/api/node/saveStatus`

Header 头部含有有效 token

````
Authorization：eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5vZGUiLCJleHAiOjE2NTc2ODk1NDcsImlhdCI6MTY1NzY3ODc0NywibmJmIjoxNjU3Njc4NzQ3fQ.P1dJnpJL9ZYcRYRPhRcHJBPj5oBv8gi6WAAXXT1yC3E`
````

正常表单数据如下，其中 `name`和`host`用来找寻节点关联：

```json
{
    "node_id": "",
    "name": "local",
    "host": "127.0.0.1",
    "load_1": 0,
    "load_5": 0,
    "load_15": 0,
    "ip_status": true,
    "thread_count": 0,
    "process_count": 348,
    "network_tx": 141168646,
    "network_rx": 2690416228,
    "network_in": 40597,
    "network_out": 647051,
    "ping_10010": 0,
    "ping_10086": 0,
    "ping_189": 0,
    "time_10010": 20,
    "time_10086": 54,
    "time_189": 100,
    "tcp_count": 311,
    "udp_count": 133,
    "cpu_count": 22.5,
    "memory_total": 16156036,
    "memory_used": 12411340,
    "swap_total": 29263236,
    "swap_used": 19486392,
    "uptime": 439346,
    "hdd_total": 485550,
    "hdd_used": 304738
}
```

返回如下：

```json
{
    "error": 0,
    "error_msg": "",
    "data": {
        "host": "127.0.0.1",
        "id": "bcd2c8d8-e7c0-457e-9afd-efe10e6a535f",
        "name": "local"
    }
}
```

### 得到全部节点数据

- 为兼容 [Server-Status](https://github.com/cppla/ServerStatus) 前端，保留此部分。

GET `http://localhost:8085/api/node/allStatus`

返回数据如下

```json
{
    "servers": [
        {
            "node_id": "0499fd55-3a26-43f7-aab9-4d6b267c13a0",
            "type": "local",
            "location": "local",
            "disabled": false,
            "custom": "",
            "name": "local",
            "host": "127.0.0.1",
            "load_1": 0,
            "load_5": 0,
            "load_15": 0,
            "ip_status": true,
            "thread_count": 0,
            "process_count": 348,
            "network_tx": 141168646,
            "network_rx": 2690416228,
            "network_in": 40597,
            "network_out": 647051,
            "ping_10010": 0,
            "ping_10086": 0,
            "ping_189": 0,
            "time_10010": 20,
            "time_10086": 54,
            "time_189": 100,
            "tcp_count": 311,
            "udp_count": 133,
            "cpu_count": 22.5,
            "memory_total": 16156036,
            "memory_used": 12411340,
            "swap_total": 29263236,
            "swap_used": 19486392,
            "uptime": 439346,
            "hdd_total": 485550,
            "hdd_used": 304738,
            "online4": false,
            "online6": true
        }
    ],
    "updated": 1657702111
}
```

## TODO

- 客户端撰写
- 完善的鉴权角色系统
- 前端管理系统
- 代码标准化改进

## 鸣谢

* ServerStatus：https://github.com/BotoX/ServerStatus
* mojeda: https://github.com/mojeda 
* mojeda's ServerStatus: https://github.com/mojeda/ServerStatus
* BlueVM's project: http://www.lowendtalk.com/discussion/comment/169690#Comment_169690
