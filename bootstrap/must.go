package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/devhg/go-gin-demo/pkg/config"
	"github.com/devhg/go-gin-demo/pkg/gredis"
	"github.com/devhg/go-gin-demo/pkg/logger"
	"github.com/devhg/go-gin-demo/pkg/mysqlc"
	"github.com/devhg/go-gin-demo/pkg/shutdown"
	"github.com/devhg/go-gin-demo/resource"
	"github.com/devhg/go-gin-demo/router"
)

var accessLogger *zap.Logger
var cronLogger *zap.Logger

func InitHTTPServer() *http.Server {
	// 初始化logger
	initLogger()
	resource.Logger = accessLogger
	resource.CronLogger = cronLogger

	// 初始化DB
	initDBServer(accessLogger)

	// 初始化Web Server
	return initHTTPServer()
}

func initLogger() {
	// 初始化 access logger
	var err error
	accessLogger, err = logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", "configs.ProjectName", "env.Active().Value()")),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(config.AppSetting.Log.LogSaveName),
	)
	if err != nil {
		panic(err)
	}

	// 初始化 cron logger
	cronLogger, err = logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", "configs.ProjectName", "env.Active().Value()")),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(config.AppSetting.Log.LogSaveName+"cron"),
	)
	if err != nil {
		panic(err)
	}
}

func initDBServer(logger *zap.Logger) {
	// 初始化 DB With ServiceName, 例如 mysql.toml
	mysqlConn, err := mysqlc.New("mysql")
	if err != nil {
		logger.Fatal("new db err", zap.Error(err))
	}
	resource.MysqlRepo = mysqlConn

	// 初始化 Cache With ServiceName, 例如 redis.toml
	// cacheRepo, err := gredis.New("redis")
	err = gredis.New("redis")
	if err != nil {
		logger.Fatal("new cache err", zap.Error(err))
	}
	// r.cache = cacheRepo

	// // 初始化 gRPC client
	// gRPCRepo, err := grpc.New()
	// if err != nil {
	// 	logger.Fatal("new grpc err", zap.Error(err))
	// }
	// r.grpConn = gRPCRepo

	// // 初始化 CRON Server
	// cronServer, err := cron_server.New(cronLogger, dbRepo, cacheRepo)
	// if err != nil {
	// 	logger.Fatal("new cron err", zap.Error(err))
	// }
	// cronServer.Start()
	// r.cronServer = cronServer
}

func initHTTPServer() *http.Server {
	router := router.NewHTTPRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.AppSetting.Server.HTTPPort),
		Handler:        router,
		ReadTimeout:    config.AppSetting.Server.ReadTimeout,
		WriteTimeout:   config.AppSetting.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accessLogger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	return server
}

// GracefulClose
func GracefulClose(server *http.Server) {
	defer func() {
		_ = resource.Logger.Sync()
		_ = resource.CronLogger.Sync()
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if resource.MysqlRepo != nil {
				if err := resource.MysqlRepo.DBWClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}

				if err := resource.MysqlRepo.DBRClose(); err != nil {
					accessLogger.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// // 关闭 cache
		// func() {
		// 	if s.Cache != nil {
		// 		if err := s.Cache.Close(); err != nil {
		// 			accessLogger.Error("cache close err", zap.Error(err))
		// 		}
		// 	}
		// },

		// // 关闭 gRPC client
		// func() {
		// 	if s.GrpClient != nil {
		// 		if err := s.GrpClient.Conn().Close(); err != nil {
		// 			accessLogger.Error("gRPC client close err", zap.Error(err))
		// 		}
		// 	}
		// },

		// // 关闭 cron Server
		// func() {
		// 	if s.CronServer != nil {
		// 		s.CronServer.Stop()
		// 	}
		// },
	)
}
