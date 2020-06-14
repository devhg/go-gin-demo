package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg      *ini.File
	HttpPort int

	RunMode string

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize    int
	JwtSecret   string
	TokenHeader string

	//数据库相关
	DbType, DbName, Username, Password, Host string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadDataBase()
	LoadApp()
}

//加载基础配置
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//加载服务器ip端口相关配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

//加载数据库相关配置
func LoadDataBase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	DbType = sec.Key("TYPE").String()
	DbName = sec.Key("NAME").String()
	Username = sec.Key("USER").String()
	Password = sec.Key("PASSWORD").String()
	Host = sec.Key("HOST").String()

}

//加载其他配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	TokenHeader = sec.Key("TOKEN_HEADER").MustString("token")
}
