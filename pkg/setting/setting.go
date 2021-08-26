package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	PageSize    int
	JwtSecret   string
	AesSecret   string
	TokenHeader string

	RuntimeRootPath string

	ImagePrefixURL string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	QrCodeSavePath string
	PrefixURL      string

	ExportSavePath string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	Username string
	Password string
	Host     string
	DBName   string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var Cfg *ini.File

func Setup() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}

	loadApp()
	loadServer()
	loadDataBase()
}

// 加载服务器ip端口相关配置
func loadServer() {
	err := Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("cfg.MapTo ServerSetting err: %v", err)
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
}

// 加载数据库相关配置
func loadDataBase() {
	err := Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("cfg.MapTo DatabaseSetting err: %v", err)
	}

	fmt.Println(DatabaseSetting)
}

// 加载其他配置
func loadApp() {
	err := Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("cfg.MapTo AppSetting err: %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
}
