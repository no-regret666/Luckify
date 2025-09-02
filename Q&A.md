Q：请简单介绍一下你的项目？

A：Luckify是一个支持用户每日签到、抽奖、评论互动的综合性系统，主要面向微信小程序端用户，
功能模块包括用户中心、签到系统、抽奖系统、评论系统、上传服务和通知模块。
- 使用Go-Zero框架构建，采用前后端分离架构；
- 任务异步处理采用Asynq实现任务队列；
- 后台调度器使用 Cron 定时调度任务； 
- 数据持久化使用 MySQL，缓存采用 Redis； 
- 使用 RPC 和 API 分别作为服务间通信和前端入口。

Q：Go-Zero为什么适合用来做这个项目？和Gin、Echo、kratos等相比有什么优势

A：Go-Zero提供完善的代码生成工具、微服务支持能力，适合构建中大型服务。它封装了API/RPC层
通信、服务治理、配置管理、熔断限流，适合这种模块多、接口清晰的业务场景。相比Gin/Echo,它自带服务治理和代码生成共工具(goctl)，开发效率高，规范统一。和kratos相比，go-zero上手更快，文档和社区更活跃，适合团队快速落地。

Q：微服务拆分原则是什么？各服务之间如何解耦与协作？

A：我按业务边界拆分服务，如用户中心、抽奖、签到、评论、通知等，每个服务独立部署，独立数据库表，服务间通过gRPC通信，接口通过proto文件定义，解耦性强。部分通用能力（如错误码、工具库）抽到common模块，避免重复开发。

Q：服务间通信为什么选择gRPC？有没有遇到什么坑?

A：gRPC性能高、IDL强类型、支持多语言，适合服务间高效通信。遇到的坑主要有proto版本兼容、字段变更要注意向后兼容，以及gRPC的超时和错误处理需要统一封装。

Q：Asynq是什么？你在项目中怎么使用它？

A：Asynq 是一个 Go 编写的基于 Redis 的异步任务队列。
我在 app/mqueue 下定义了两个部分： 
- job：处理任务逻辑（比如开奖、发送通知等）；
- scheduler：负责使用 cron 表达式定时生成任务投递，比如每日定时触发开奖任务或签到提醒任务。

任务定义通过 jobtype 包配置（包含任务类型、payload），最终通过 asynq.Enqueue() 投递到 Redis 中。

Q：怎么实现任务调度？

A：通过 github.com/robfig/cron/v3，在 app/mqueue/cmd/scheduler/internal/logic/*.go 中注册调度任务，在 register.go 里进行统一注册调度逻辑。
例如：每天中午 12 点，定时执行未开奖抽奖的开奖逻辑：
```
c.AddFunc("0 0 12 * * ?", func() {
    // 扫描未开奖抽奖
    // 创建任务投递到 Asynq 队列
})
```

Q：Redis在系统中主要承担哪些角色？如何保证其高可用和数据一致性？

A：Redis用于缓存热点数据、分布式锁、Asynq队列。高可用用主从+哨兵，数据一致性通过合理设置过期、双写一致性和定期修正（如缓存失效时回源DB）来保证。分布式锁用redsync库实现，防止超卖。

Q：Luckify如何支撑高并发抽奖？有哪些具体的性能优化措施？

A：
- 关键路径用Redis缓存和分布式锁，减少DB压力
- 抽奖库存预减用原子操作+分布式锁，防止超卖
- 异步任务解耦耗时操作
- 监控链路追踪，及时发现瓶颈

Q：即抽即中模式下如何防止超卖？分布式锁的实现细节是什么？

A：

Q：数据库和缓存的读写策略是怎样的？如何避免缓存击穿和雪崩？

A：
- 读多写少的数据优先查缓存，缓存失效时回源DB并回填
- 热点key设置合理过期时间，防止同时失效
- 对于高并发key,采用互斥锁/预热等方式防止击穿
- 雪崩时可用降级策略或限流

Q：抽奖算法是怎么设计的？如何保证公平性和高性能？

A：采用Fisher-Yates洗牌算法，保证每个参与者中奖概率均等。即抽即中模式下，先判断库存，再根据概率区间判断是否中奖，整个过程在内存和事务中完成，性能高且公平。

Q：如何监控和定位系统的性能瓶颈？用过哪些链路追踪和监控工具？

A：用Prometheus和Grafana做指标监控，Jaeger做链路追踪，Filebeat+Elasticsearch收集日志。通过监控面板和trace分析定位慢查询、热点接口、异常等。

Q：抽奖和库存扣减是如何保证强一致性的？用到了哪些事务和补偿机制？

A：库存扣减和中奖记录写入在同一个数据库事务中完成，保证原子性。分布式场景下用分布式锁防止并发冲突。异步任务失败后会自动重试，保证最终一致性。

Q：如果Redis和MySQL数据不一致，如何修复？

A：以MySQL为准，定期比对数据，发现不一致时回源DB修正缓存。关键操作采用双写或延迟双删策略，减少不一致概率。

Q：分布式事务在系统中是怎么处理的？有没有用到消息最终一致性？

Q：Luckify如何保证高可用？服务挂了会发生什么？

Q：如何做服务的健康检查和自动恢复？

Q：限流、熔断、降级策略是怎么做的？

Q：代码生成工具（goctl）在实际开发中带来了哪些便利和问题？

Q：Luckify如果要支持更多抽奖玩法（如拼团、秒杀），你会怎么扩展架构？

Q：积分体系和任务系统如何设计，如何防止刷积分？

Q：多端适配如何做？

Q：遇到的最大技术挑战是什么？怎么解决的？

Q：这个项目让你在哪些方面成长最大？如果重做，有哪些地方会优化？

Q：如何做单元测试和集成测试？

Q：安全方面做了哪些防护？

Q：如何防止刷奖、作弊等行为？

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
  
	....... 业务逻辑func f(s string) bool{
    i,j := 0,len(s) - 1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}
func partition(s string) [][]string {
    ans := [][]string{}
    path := []string{}
    n := len(s)
    var dfs func(i,start int) // i：下一个要判断添不添加的逗号  start：子串开始的位置
    dfs = func(i,start int) {
        if i == n {
            ans = append(ans,slices.Clone(path))
            return
        }

        // 不分割
        if i < n {
            dfs(i + 1,start)
        }

        //分割
        str := s[start:i + 1]
        if f(str) {
            path = append(path,str)
            dfs(i + 1,i + 1)
            path = path[:len(path) - 1]
        }
    }
    dfs(0,0)
    return ans
}

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

