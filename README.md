# restful-api-demo

#### 使用说明
```
.
├── apps
│   ├── all
│   │   └── impl.go                // 所有handler和impl的init函数执行的地方
│   ├── models
│   │   └── {{appName}}.go         // pojo，与数据库的字段一一对应
│   └── dao
│   |    └── mysql.go              // 初始化和mysql的连接
│   |    └── ×××××.go              // 执行sql的文件
│   └── {{hanler_dir}}             // app的业务代码
│   |    └── http                  // handler
│   |        └── {{appName}}.go    // handler业务代码
│   |        └── http.go           // 把handler中的路由暴露出去
│   |    └── impl
│   |        └── {{appName}}.go    // 对接口的实现
|   |         └── ×××××.go         // 定义业务接口和参数
│   ├── ioc.go                     // impl和handler的注册中心，并且从其中取
│   └── app.go                     // 定义app的名称，以便在ioc中注册
├── cmd                            // 用户自定义路由写入位置
|   └── root.go                    // 命令行的根
|   └── start.go                   // start命令，项目的初始化在这里执行
├── etc                            // 配置文件
├── config
|   └── config.go                  // 配置文件对应的结构体
|   └── load.go                    //从配置文件中加载配置
├── protocol
|   └── http.go                    // http服务启动的配置
├── Makefile.go                    // hz 生成的路由注册调用
├── version
|   └── version.go                 // 版本控制
├── go.mod
├── main.go                        // 启动入口
└── router_gen.go                  // hz 生成的路由注册调用
```


