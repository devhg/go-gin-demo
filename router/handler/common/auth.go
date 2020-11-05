package common

import (
	"github.com/QXQZX/go-gin-demo/middleware/jwt"
	"github.com/QXQZX/go-gin-demo/model"
	"github.com/QXQZX/go-gin-demo/pkg/app"
	"github.com/QXQZX/go-gin-demo/pkg/e"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	var auth auth
	httpStatus, eCode, _ := app.BindAndValid(c, &auth)

	data := ""
	code := e.INVALID_PARAMS
	if eCode == 200 {
		checkAuth, _ := model.CheckAuth(auth.Username, auth.Password)
		if checkAuth {
			token, err := jwt.GenerateToken(auth.Username, auth.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	}
	appG.Response(httpStatus, code, data)
}
