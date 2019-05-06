# timewheel
timewheel on golang

### 使用
```golang
timewheel:=NewTimeWheel()
timewheel.Start()
timewheel.SetTimeout(3000, func() {})
id:=timewheel.SetInterval(3000, func() {})
timewheel.Stop(id)
```

### 主要逻辑
- 使用最小堆，取注册的离当前最近的定时任务
- 若任务过期，立即执行，否则等待至规定的时间
- 在等待过程中有其他任务注册时，将重新计算最近的定时任务

### 其他
- 加入极简协程池 goroutine pool，减少阻塞的情况
- 使用 map 记录被关闭的任务，当任务为最近的定时任务时，将被直接抛出


### 后续计划
- 加入 WaitGroup,与主线程同步
- 加入多个时间轮
- 协程池将根据任务量，执行情况动态管理