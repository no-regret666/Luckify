# Luckify

## 项目简介
Luckify 是一个基于[go-zero](https://github.com/zeromicro/go-zero)框架开发的高性能抽奖系统后台。
该系统旨在提供灵活、可靠的抽奖功能，满足多种业务场景需求。

## 功能特性

## 技术栈

- **Go-zero**：高性能的Go微服务框架，支持快速开发和高效扩展
- **MySQL**：用于存储用户、抽奖活动和奖品数据
- **Redis**：分布式锁，缓存用户信息、抽奖活动状态和奖品库存，提升系统性能。
- **Asynq**：异步任务，用于处理抽奖活动和奖品发放
- **Kafka**：消息队列，用于日志收集和异步处理
- **Jaeger**：链路追踪，用于监控和分析系统性能
- **Go-stash**：日志处理，用于从Kafka中获取数据并进行处理
- **Filebeat**：日志收集，用于从容器中收集数据
- **Elasticsearch**：日志存储，用于存储和索引日志数据
- **Docker**：容器化部署，方便环境配置和项目运行

## 目录结构
- **app**：应用层，负责处理业务逻辑
- **common**：公共模块，包含常量、错误码、工具函数等
- **deploy**：部署脚本
- **doc**：文档目录，包含项目设计文档和API文档

## 快速开始
### 环境要求

- Go 1.23及以上版本
- MySQL 8.0及以上版本
- Redis 7.0及以上版本

### 启动步骤
#### 1.使用`docker-compose`启动
如果需要手动启动服务，运行以下命令
```shell
# 启动环境依赖
docker-compose -f docker-compose-env.yml up -d

# 启动主服务
docker-compose -f docker-compose.yml up -d
```

#### 2.使用`make`简化启动和关闭
```shell
# 启动所有服务
make docker-up-env
make docker-up-app
# 停止所有服务
make docker-down-env
make docker-down-app
```