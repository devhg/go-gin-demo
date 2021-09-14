package config

import (
	"io/fs"
	"log"
	"path"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Servicer interface {
}

type serviceConfig struct {
	Name     string `toml:"Name"`
	Retries  int    `toml:"Retries"`
	SmartDNS string `toml:"SmartDNS"`
	Timeout  int    `toml:"Timeout"`
}

var ConfigLocation = "./config"

var ConfigMap map[string]Servicer

// /////////////////////////////////////////////////////////

type ServerConfig struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MailConfig struct {
	Host     string
	Port     int
	UserName string
	PassWord string
	To       []string
}

type AppConfig struct {
	// config for http server
	Server ServerConfig

	// config for application service
	App struct {
		PageSize    int
		JwtSecret   string
		AesSecret   string
		TokenHeader string

		RuntimeRootPath string

		ImagePrefixURL string
		ImageSavePath  string
		ImageMaxSize   int
		ImageAllowExts []string

		QrCodeSavePath string
		PrefixURL      string

		ExportSavePath string
		TimeFormat     string
	}

	// config for log
	Log struct {
		LogSavePath string
		LogSaveName string
		LogFileExt  string
		TimeFormat  string
	}

	Mail *MailConfig
}

var AppSetting = &AppConfig{}

// /////////////////////////////////////////////////////////

func MustLoadConfig(confLocation string) error {
	if confLocation != "" {
		ConfigLocation = confLocation
	}

	err := filepath.Walk(ConfigLocation, func(filepath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || path.Ext(filepath) != ".toml" {
			// return errors.New("not matched config file format")
			return nil
		}

		// filepath := "device/sdk/app.toml"
		filename := path.Base(filepath)           // app.toml
		ext := path.Ext(filepath)                 // .toml
		dir := path.Dir(filepath)                 // device/sdk
		name := filename[:len(filename)-len(ext)] // app

		if ConfigMap == nil {
			ConfigMap = make(map[string]Servicer)
		}

		switch name {
		case "app":
			loadConfig(name, ext[1:], dir, AppSetting)
		case "redis":
			loadConfig(name, ext[1:], dir, RedisConfig)
			ConfigMap["redis"] = RedisConfig
		case "mysql":
			loadConfig(name, ext[1:], dir, MysqlConfig)
			ConfigMap["mysql"] = MysqlConfig
		default:
			service := &serviceConfig{}
			loadConfig(name, ext[1:], dir, service)
			ConfigMap[name] = service
		}
		return nil
	})

	return err
}

func loadConfig(name, typ, dir string, v interface{}) {
	viper.SetConfigName(name)
	viper.SetConfigType(typ)
	viper.AddConfigPath(dir)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(v); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println(e)
		if err := viper.Unmarshal(v); err != nil {
			panic(err)
		}
	})
}

func GetServicer(name string) Servicer {
	return ConfigMap[name]
}
