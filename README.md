# FastApp

快速开发Gin 服务

## Build

        ./build.sh
        或者
        DOCKER_USER=aker DOCKER_PASS=123 ./build.sh #DOCKER_USER DOCKER_PASS 为对应docker 登陆账号


## 项目结构

```
|- frontend // vue开发前端
|- src // 项目逻辑代码
    |- config // 配置文件
    |- db // 数据
        |- schema // Gorm数据库定义
|- .gitignore
|- .travis.yml // travis 发布到gitub 之后自动打包编译
|- build.sh // 打包发布工具
|- Dockerfile // build.sh 配置打包
```

## 开发说明

本项目依赖项目: [ginFastApp](https://github.com/shaohung001/ginFastApp.git)

当前支持写入配置项启动:

```go
app := ginFastApp.New(&Config{...})
err := app.Start()
```

以上即启动服务器。还可先ConnectDB, 实现接口方法IConfig.getDB()相应配置即可:

```go
app.ConnectDB(func(db *gorm.DB, err error) {
  ...
  app.Start()
})
```

如果要加入中间件，当前支持`app.addBeforeMiddleware` 和 `app.addAfterMiddleware` 对应加入到中间件中。