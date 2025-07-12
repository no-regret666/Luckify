# `asynq`

使用`asynq`的几个原因：
- 直接基于redis,一般项目都有redis,而asynq本身就是基于redis所以可以少维护一个中间件
- 支持消息队列、延时队列、定时任务调度，因为希望项目支持定时任务而asynq直接就支持
- 有webui界面，每个任务都可以暂停、归档、通过ui界面查看成功失败、监控

go-zero集成asynq： 

定义任务：

`job`模块：定义、注册和处理任务，消息队列中的普通异步任务，也叫即时任务或延迟任务。

投递任务：

1. 业务事件触发的任务，由`api/rpc`主动投递

如notice模块的rpc层中：
```
_, err = asynqClient.Enqueue(
asynq.NewTask(jobtype.MsgWxMiniProgramNotifyUser, payloadBytes),
)

```
- 适用于：当某件事立即需要做，但又不想阻塞业务线程时

2. 周期执行的任务，由`scheduler`定时投递

就在`mqueue/scheduler`层中：
```
scheduler.Register("0 0 * * *", asynq.NewTask(jobtype.ScheduleLotteryDraw, nil))
```
- 每天 0 点调度一次 ScheduleLotteryDraw 任务，任务类型仍然是 jobtype 中定义的。

尽管投递来源不同，但本质上：
- 所有任务最终都会进入 Redis 队列
- 都由 Asynq worker + mux.Handle(...) 注册的 handler 来处理

