# Luckify

## 项目简介
Luckify 是一个基于[go-zero](https://github.com/zeromicro/go-zero)框架开发的高性能抽奖系统后台。
该系统旨在提供灵活、可靠的抽奖功能，满足多种业务场景需求。

## 核心特性
- 完整的抽奖生态：支持即时抽奖、定时抽奖、多种奖品类型配置
- 用户体系：完整的用户注册、登录、资料管理
- 评论互动：评论系统、支持点赞、回复等社区功能
- 签到系统：每日签到、任务系统、积分体系
- 文件管理：统一的文件上传服务，支持 MinIO 对象存储
- 实时监控：链路追踪、日志收集、性能监控

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