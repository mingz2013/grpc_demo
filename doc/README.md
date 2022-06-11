## 我们的服务需要什么功能？
- 数据库访问redis/mongo/mysql，orm
- 配置管理，同步
- gateway长连接，tcp/ws
- gateway http短链接
- rpc同步调用
- mq异步发送
- 服务集群管理，扩容/缩容，滚动发布。有状态/无状态
- 业务服务定义与业务逻辑实现
- 可以单进程本地机器部署调试
- 多版本服务支持。
- 限流，削峰，
- 高可用
- 消息协议。
- 分布式事务
- 



## 开源方案
- 注册中心 etcd，zookeeper，consul。。。   http://www.360doc.com/content/22/0315/07/412471_1021570651.shtml
- 配置中心 etcd，consul，zookeeper kv服务。配置大小限制。。  http://dockone.io/article/8767
- rpc  grpc
- 协议 protobuf
- 负载均衡 一致性hash，轮询。。
- 服务集群 docker swarm，k8s
- 消息队列 kafka，rocketmq，。。
- 同步消息rpc，异步消息mq，
- 有状态服务扩容缩容，一致性hash。无状态扩容缩容，轮询



## 思考 
- redis替换注册中心是否可以？
- redis替换配置中心是否可以？
- redis pubsub mq是否可以？
- redis替换存储是否可以？



## 后续
简单演示，后续有机会详解每个部分
