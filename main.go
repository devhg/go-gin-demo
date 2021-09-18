package main

import (
	"flag"
	"github.com/devhg/go-gin-demo/bootstrap"

	_ "github.com/devhg/go-gin-demo/docs"
	"github.com/devhg/go-gin-demo/pkg/config"
	"github.com/devhg/go-gin-demo/pkg/logging"
)

var conf = flag.String("conf_path", "./config", "input config path")

func init() {
	flag.Parse()
}

// @title go-gin-demo
// @version 1.0
// @description 用go+gin搭建web网站后端接口
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8081
// @BasePath
func main() {
	err := config.MustLoadConfig(*conf)
	if err != nil {
		panic(err)
	}

	logging.Setup()
	logging.Info("Ready to start.")

	server := bootstrap.InitHTTPServer()

	// r := router.NewHTTPRouter()
	//
	// server := &http.Server{
	// 	Addr:           fmt.Sprintf(":%s", config.AppSetting.Server.HTTPPort),
	// 	Handler:        r,
	// 	ReadTimeout:    config.AppSetting.Server.ReadTimeout * time.Second,
	// 	WriteTimeout:   config.AppSetting.Server.WriteTimeout * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	//
	// log.Fatal(server.ListenAndServe())

	bootstrap.GracefulClose(server)
}
