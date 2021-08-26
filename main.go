package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/devhg/go-gin-demo/docs"
	"github.com/devhg/go-gin-demo/model/dao"
	"github.com/devhg/go-gin-demo/pkg/logging"
	"github.com/devhg/go-gin-demo/pkg/setting"
	"github.com/devhg/go-gin-demo/router"
)

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
	setting.Setup()
	dao.Setup()
	logging.Setup()

	logging.Info("Ready to start.")

	router := router.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	logging.Info("Started in ", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server error %v\n", err)
	}
}
