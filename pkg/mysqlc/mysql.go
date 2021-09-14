package mysqlc

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/devhg/go-gin-demo/pkg/config"
)

var _ Repo = (*DBRepo)(nil)

type Repo interface {
	GetDBR() *gorm.DB
	GetDBW() *gorm.DB
	DBRClose() error
	DBWClose() error
}

type DBRepo struct {
	DBR *gorm.DB
	DBW *gorm.DB
}

func New(serviceName string) (Repo, error) {
	servicer := config.GetServicer(serviceName)
	cfg := servicer.(*config.MySQL)

	repo := &DBRepo{}

	if cfg.Read != nil {
		DBR, err := dbConnect(cfg.Read)
		if err != nil {
			return nil, err
		}
		repo.DBR = DBR
	}

	if cfg.Write != nil {
		DBW, err := dbConnect(cfg.Write)
		if err != nil {
			return nil, err
		}
		repo.DBW = DBW
	}

	return repo, nil
}

func (d *DBRepo) GetDBR() *gorm.DB {
	return d.DBR
}

func (d *DBRepo) GetDBW() *gorm.DB {
	return d.DBW
}

func (d *DBRepo) DBRClose() error {
	sqlDB, err := d.DBR.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *DBRepo) DBWClose() error {
	sqlDB, err := d.DBW.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func dbConnect(conf *config.MysqlMeta) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.DBName,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", conf.DBName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(conf.MaxOpenConn)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(conf.MaxIdleConn)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(conf.MaxLifeTime))

	// 使用插件
	db.Use(&TracePlugin{})

	return db, nil
}
