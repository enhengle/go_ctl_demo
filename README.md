### go_ctl的练习demo

### 功能讲解
```.env
①连接mysql
②连接redis
③go_ctl本身配置文件读取
④自定义配置文件读取
⑤全局异常捕捉
⑥日志打印分割
⑦请求接口示例（get/post）
⑧路由转发
```

### 目录讲解
```.env
│  .env 自定义配置文件
│  .env.develop 自定义配置文件
│  go.mod
│  go.sum
│  greet.api
│  greet.exe
│  greet.go 主方法
│  README.md
│
├─.idea
│      .gitignore
│      greet.iml
│      misc.xml
│      modules.xml
│      vcs.xml
│      workspace.xml
│
├─etc
│      greet-api.yaml
│
├─internal  web层
│  ├─config 配置层
│  │      config.go
│  │      config_olf.go
│  │
│  ├─handler    路由层+数据层
│  │  │  greethandler.go    
│  │  │  routes.go  
│  │  │
│  │  ├─cache
│  │  │      db.go
│  │  │      redis.go
│  │  │
│  │  ├─http
│  │  │      get.go
│  │  │      postJson.go
│  │  │
│  │  ├─redis
│  │  │      redis.go
│  │  │      redis2.go
│  │  │      redisPool.go
│  │  │
│  │  └─sql
│  │          mysql.go
│  │          mysql2.go
│  │
│  ├─logic
│  │      greetlogic.go
│  │
│  ├─svc
│  │      servicecontext.go
│  │
│  ├─types
│  │  │  types.go
│  │  │
│  │  └─business    数据库映射
│  │          table_name.go
│  │
│  └─xlogger    日志层
│          log.go
│          logCore.go
│
└─support   配置文件
    └─genv
            dotfile.go
            loader.go
            unmarshal.go
```