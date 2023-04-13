# golang contract project admin 👍

## init project

### 1 安装 adm

``` shell
go install github.com/GoAdminGroup/adm
```

### 2 修改 config.json

### 3 配置gorm

``` go
 // 本代码中配置多数据源 sqlite数据库为Admin框架所需,mysql为业务数据库
 orm, err = gorm.Open(sqlite.Open(s.GetConfig("default").File), &gorm.Config{
  Logger: logger.Default.LogMode(logger.Info),
 })
 orm.Use(dbresolver.Register(
  dbresolver.Config{
   Replicas: []gorm.Dialector{mysql.Open(m.GetConfig("main").GetDSN())},
  }, "blind"),
 )

本地数据库配置
     "local": {
      "max_idle_con": 50,
      "max_open_con": 150,
      "driver": "mysql",
      "host": "",
      "port": "3306",
      "user": "",
      "pwd": "",
      "name": ""
    },
```

### 4 运行

初始登录账户 admin 密码 123456

```shell
go run main.go
go run main.go -f etc/config.json -p port
```

### 5 开发说明  里面都有示例

1. cron 计划任务
2. eventW3 合约事件监听
3. models 数据库操作
4. handler gin api 扩展
5. pages 页面扩展
6. tables 数据管理
7. docker 容器部署 目前只有进入容器启动服务版本
  