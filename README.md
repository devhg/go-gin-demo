
# go web 项目脚手架

```ecmascript 6
├── README.md
├── conf        #放置配置文件
│   └── app.ini
├── go.mod
├── go.sum
├── handler     #请求处理，类似controller层
│   ├── contest
│   ├── train
│   └── user
├── main.go     #程序入口
├── middleware  #中间件
│   ├── Cors    #跨域支持中间件
│   └── jwt     #jwt授权中间件
├── model       #数据层，对接数据库操作
│   ├── feedback.go
│   ├── model.go
│   ├── notice.go
│   ├── user.go
│   └── userinfo.go
├── pkg         #其他包
│   ├── app
│   ├── constvar#常量
│   ├── e       #错误代码，错误信息
│   ├── logging #日志打印
│   ├── setting #项目设置
│   ├── upload  #长传文件处理
│   └── util    #工具包
├── router      #存储路由信息
│   ├── auth.go
│   ├── router.go
│   └── upload.go
├── runtime     #存储日志
│   └── logs
└── service     #service层，业务逻辑处理
    ├── admin_service
    ├── contest_service
    ├── notice_service
    ├── problem_service
    └── user_service


```