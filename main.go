package main

import (
	"fmt"
	"github.com/QXQZX/go-exam/pkg/setting"
	"github.com/QXQZX/go-exam/router"
	"net/http"
)

// @title go-exam
// @version 1.0
// @description 用go+gin搭建web网站后端接口
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8081
// @BasePath
func main() {
	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
