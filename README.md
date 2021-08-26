
# go web 项目脚手架

关于目录结构建议参考
https://github.com/golang-standards/project-layout

```ecmascript 6
.
├── Dockerfile
├── README.md
├── conf
│   └── app.ini
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go-gin-demo
├── go.mod
├── go.sum
├── main.go
├── middleware
│   ├── cors
│   │   └── crossorigin.go
│   ├── crypto
│   │   └── crypto.go
│   └── jwt
│       └── jwt.go
├── model
│   ├── dao
│   │   ├── feedback.go
│   │   ├── model.go
│   │   ├── notice.go
│   │   ├── user.go
│   │   └── userinfo.go
│   └── service
│       ├── admin
│       ├── contest
│       ├── notice
│       ├── problem
│       └── user
├── pkg
│   ├── app
│   │   ├── localtime.go
│   │   ├── request.go
│   │   ├── response.go
│   │   └── validation.go
│   ├── constvar
│   │   └── constvar.go
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── export
│   │   └── excel.go
│   ├── file
│   │   └── file.go
│   ├── gredis
│   │   └── redis.go
│   ├── logging
│   │   ├── file.go
│   │   └── log.go
│   ├── qrcode
│   │   └── qrcode.go
│   ├── setting
│   │   └── setting.go
│   ├── upload
│   │   └── image.go
│   └── util
│       ├── md5.go
│       └── pagination.go
├── router
│   ├── handler
│   │   ├── common
│   │   ├── contest
│   │   ├── train
│   │   └── user
│   └── router.go
├── runtime
│   └── logs
│       └── log20210826.log
└── views
    ├── exam
    │   ├── favicon.ico
    │   ├── index.html
    │   └── static
    ├── index.html
    ├── layouts
    │   ├── footer.html
    │   └── master.html
    └── page.html
```

Docker
```shell script
$ docker build -t gin-demo-docker .

$ docker run -d -p 8081:8081 gin-demo-docker

```

使用小镜像
```shell script

$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-demo .

$ docker build -t gin-demo-docker-scratch .

$ docker run -d -p 8081:8081 gin-demo-docker-scratch

```