FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/QXQZX/go-gin-demo
COPY . $GOPATH/src/github.com/QXQZX/go-gin-demo
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-demo"]