# SneakerWorker for Golang


### Dependencies 依赖
* RabbitMQ

### Usage 使用方法
Create two files below in config folder which is in your project folder:

在你的项目中的config目录下创建以下两个文件

```
amqp.yml
workers.yml
```
This is a [example](https://github.com/oldfritter/sneaker-go/blob/master/example/config).

### workers.yml配置说明
```
---
- name: TreatWorker  # worker的名称
  exchange: sneaker.example.default  # 消息经过的Exchange
  routing_key: sneaker.example.treat  # 消息经过的routing_key
  queue: sneaker.example.treat  # 消息进入的queue
  durable: true
  ack: true  # 是否ack
  threads: 1  # 并发处理数量
  steps:  # 重试队列的延时配置
    - 5000       # 5 Second
    - 30000      # 30 Second
    - 60000      # 1 Minute
```