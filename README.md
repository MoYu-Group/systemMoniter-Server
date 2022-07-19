# systemMoniter-Server
Golang 写的系统资源监控服务器端

使用 Golang 开发，MySQL 存储，使用 cppla 的 [ServerStatus](https://github.com/cppla/ServerStatus)，作为简要前端

使用框架依赖库 Gin，Zap，Gorm，gopsutil

本人代码水平有限，仅供学习参考使用，请勿商用。

客户端 [systemMoniter-Node](https://github.com/MoYu-Group/systemMoniter-Node)

## 已知问题

- 网络流量计算算法问题
- 前端侦测节点离线算法优化
- 前端调用 API 接口地址目前是写死的
- LINUX 未测试

## TODO

- LINUX 系统测试
- 前端显示优化改进
- 完善的鉴权角色系统
- 前端管理系统
- 代码标准化改进
- 性能测试优化

## 鸣谢

* cppla 的 ServerStatus: https://github.com/cppla/ServerStatus
* ServerStatus: https://github.com/BotoX/ServerStatus
* mojeda: https://github.com/mojeda 
* mojeda's ServerStatus: https://github.com/mojeda/ServerStatus
* BlueVM's project: http://www.lowendtalk.com/discussion/comment/169690#Comment_169690
