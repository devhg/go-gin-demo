
# go web 项目脚手架

```ecmascript 6
├── Dockerfile
├── README.md
├── conf
│   └── app.ini
├── go.mod
├── go.sum
├── main.go
├── middleware
│   ├── cors
│   ├── crypto
│   └── jwt
├── model
│   ├── feedback.go
│   ├── model.go
│   ├── notice.go
│   ├── user.go
│   └── userinfo.go
├── pkg
│   ├── app
│   ├── constvar
│   ├── e
│   ├── export
│   ├── file
│   ├── gredis
│   ├── logging
│   ├── qrcode
│   ├── setting
│   ├── upload
│   └── util
├── router
│   ├── handler
│   └── router.go
├── runtime
│   ├── logs
│   ├── qrcode
│   └── upload
├── service
│   ├── admin
│   ├── contest
│   ├── notice
│   ├── problem
│   └── user
└── views
    ├── exam
    ├── index.html
    ├── layouts
    └── page.html

```

Docker
```shell script
$ docker build -t gin-demo-docker .

$ docker run -p 8081:8081 gin-demo-docker

```