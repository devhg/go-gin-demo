package resource

import (
	"go.uber.org/zap"

	"github.com/devhg/go-gin-demo/pkg/mysqlc"
)

var Logger *zap.Logger

var CronLogger *zap.Logger

var MysqlRepo mysqlc.Repo
