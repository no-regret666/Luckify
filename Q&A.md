一、项目整体设计

Q：请简单介绍一下你的项目？

A：Luckify是一个支持用户每日签到、抽奖、评论互动的综合性系统，主要面向微信小程序端用户，
功能模块包括用户中心、签到系统、抽奖系统、评论系统、上传服务和通知模块。
- 使用Go-Zero框架构建，采用前后端分离架构；
- 任务异步处理采用Asynq实现任务队列；
- 后台调度器使用 Cron 定时调度任务； 
- 数据持久化使用 MySQL，缓存采用 Redis； 
- 使用 RPC 和 API 分别作为服务间通信和前端入口。

二、核心技术与架构设计
1. Go-Zero

Q：Go-Zero为什么适合用来做这个项目

A：Go-Zero提供完善的代码生成工具、微服务支持能力，适合构建中大型服务。它封装了API/RPC层
通信、服务治理、配置管理、熔断限流，适合这种模块多、接口清晰的业务场景。

2. Asynq使用

Q：Asynq是什么？你在项目中怎么使用它？

A：Asynq 是一个 Go 编写的基于 Redis 的异步任务队列。
我在 app/mqueue 下定义了两个部分： 
- job：处理任务逻辑（比如开奖、发送通知等）；
- scheduler：负责使用 cron 表达式定时生成任务投递，比如每日定时触发开奖任务或签到提醒任务。

任务定义通过 jobtype 包配置（包含任务类型、payload），最终通过 asynq.Enqueue() 投递到 Redis 中。

3. Cron调度器

Q：怎么实现任务调度？

A：通过 github.com/robfig/cron/v3，在 app/mqueue/cmd/scheduler/internal/logic/*.go 中注册调度任务，在 register.go 里进行统一注册调度逻辑。
例如：每天中午 12 点，定时执行未开奖抽奖的开奖逻辑：
```
c.AddFunc("0 0 12 * * ?", func() {
    // 扫描未开奖抽奖
    // 创建任务投递到 Asynq 队列
})
```

三、数据库 & 模型设计

Q：如何设计数据库表结构？

A：
1. 自上而下方法：
- 从业务需求出发，先设计出数据模型，再根据数据模型设计数据库表结构，最后进行实现。
- 这种方法适合于对数据一致性要求高、数据结构比较稳定的场景，例如业务管理系统、电商系统。
- 自上而下的设计方法可以提高数据的一致性和可维护性，但可能会造成过度设计和性能问题。
- 提高了代码的可读性和可维护性，因为在设计表结构时，我们可以考虑表与表之间的关系，使用外键等方式实现表之间的关联。

2. 自下而上方法：
- 从数据库实现出发，先设计出数据库表结构，再根据表结构设计数据模型，最后根据数据模型实现业务需求。
- 这种方法适合于对数据结构变化较快、数据访问频繁、数据量较大的场景，例如物联网、大数据等。
- 可以提高性能和可扩展性，但可能会降低数据的一致性和可维护性。

四、服务间通信 & 微服务拆分

Q：为什么采用微服务拆分？模块间怎么通信？

A：
为了职责分离、便于扩展与部署，我们采用了 Go-Zero 的 RPC 模式进行模块拆分。
各模块通过 .proto 文件定义服务接口，使用 goctl 生成代码，服务间通过 gRPC 进行通信，例如：

- lottery 模块通过调用 checkin RPC 查询用户积分； 
- comment 模块调用 usercenter RPC 获取用户信息等。

五、错误处理

Q：怎么处理错误？

A：

**rpc：**
```
package codes // import "google.golang.org/grpc/codes"

import (
	"fmt"
	"strconv"
)

// A Code is an unsigned 32-bit error code as defined in the gRPC spec.
type Code uint32
.......
```
grpc的err对应的错误码是一个uint32，我们自定义错误用uint32并在rpc的全局拦截器返回时候转成grpc的err就可以被客户端正常接收。

自定义全局错误码在app/common/xerr中；

```
if err != nil {
	return errors.Wrapf(xerr.ErrDBError, "insertResult.LastInsertId err:%v,user:%+v", err, user)
}
``` 
第一个参数返回给前端，第二个参数记录日志。

通过在rpc的启动文件main方法中，添加grpc的全局拦截器实现。

```
package main

......

func main() {
  
	....... 业务逻辑

	//rpc log,grpc的全局拦截器
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	.......
}
```

LoggerInterceptor在common/interceptor/rpcserver中实现。

**api：**

common/response中实现了错误处理。

- 参数错误：ParamErrorResult 函数
- 业务返回的错误：HttpResult 函数
  - 成功直接返回
  - 如果错误先判断：
    - 1.自定义错误（api中定义的错误），直接返回给前端 
    - 2.grpc错误（rpc业务抛出来的），通过uint32转成我们自己的错误码，根据错误码再去我们自定义的错误信息中找到定义的错误信息返回给前端
    - 3.都找不到就返回默认错误“服务器开小差了”



六、部署与测试

Q：怎么部署与测试？

A：