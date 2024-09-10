# zlog
Very small logging library

Source code from：[zlsgo](https://github.com/sohaha/zlsgo)


### 日志工具

```go
package main

import (
    "github.com/lei006/zlog"
)

func main(){
    logs := []string{"这是一个测试","这是一个错误"}
    zlog.Debug(logs[0])
    zlog.Error(logs[1])
    zlog.Dump(logs)
    // zlog...
}
```




### 基本用法
```go
  package main

  import (
      "github.com/lei006/zlog"
  )

  func main(){
          zlog.Debug("我是一个测试")
          zlog.Warnf("%s\n","我是一个警告")
          zlog.Error("我是一个错误日志")
          hi := "我是特别的有用的，不信你可以用我打印一个变量试试"
          zlog.Dump(hi)
          zlog.Stack("打印下堆栈")
  }
```


### 日志前缀
```go

  // 初始化一个新实例
  var logger = zlog.New("前缀")

  // 动态设置
  zlog.SetPrefix("前缀")

```


### 日志等级
```go

  // 日志等级信息参考下面：
  const (
      LogFatal = iota
      LogPanic
      LogError
      LogWarn
      LogSuccess
      LogInfo
      LogDebug
      LogDump
  )

  // 设置输出的日志等级，比如不输出 info debug 等级的日志
  zlog.SetLogLevel(zlog.LogSuccess)

  // 获取当前输出的日志等级
  zlog.GetLogLevel()

```


### 颜色日志
```go
  // 关闭颜色输出
  zlog.DisableConsoleColor()

  // 重新开启颜色输出
  zlog.ForceConsoleColor()
```



### 保存日志
```go
  // 把日志保存到指定文件
  zlog.SetFile("文件路径")

  // 日志按日期归档
  zlog.SetFile("文件路径", true)
  // 日志按日期归档，那么日志最长的保存日期是半个月，15 天之前的日志会自动删除，
  // 如果想自定义移除时长为 30 天
  zlog.LogMaxDurationDate = 30

  // 把日志保存到指定文件同时也在终端输出
  zlog.SetSaveFile("文件路径")

  // 我实际使用的
  zlog.SetSaveFile("logs.log", true)

```
