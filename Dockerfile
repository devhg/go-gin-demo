FROM scratch

WORKDIR $GOPATH/src/github.com/QXQZX/go-gin-demo
COPY . $GOPATH/src/github.com/QXQZX/go-gin-demo

EXPOSE 8000
ENTRYPOINT ["./go-gin-demo"]