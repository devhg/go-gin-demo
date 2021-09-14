package config

type MysqlMeta struct {
	Host        string `toml:"Host"`
	UserName    string `toml:"UserName"`
	Password    string `toml:"Password"`
	DBName      string `toml:"DbName"`
	MaxIdleConn int    `toml:"MaxIdleConn"`
	MaxOpenConn int    `toml:"MaxOpenConn"`
	MaxLifeTime int    `toml:"MaxLifeTime"`
}

type MySQL struct {
	// common config use dns name load balance
	Service serviceConfig `toml:"service"`

	// mysql service
	Read  *MysqlMeta `toml:"read"`
	Write *MysqlMeta `toml:"write"`
}

var MysqlConfig = &MySQL{}

// func loadMysqlConfig() {
// 	viper.SetConfigName("mysql")
// 	viper.SetConfigType("toml")
// 	viper.AddConfigPath(path.Join(ConfigLocation, "servicer/"))

// 	if err := viper.ReadInConfig(); err != nil {
// 		panic(err)
// 	}

// 	if err := viper.Unmarshal(MysqlConfig); err != nil {
// 		panic(err)
// 	}

// 	viper.WatchConfig()
// 	viper.OnConfigChange(func(e fsnotify.Event) {
// 		if err := viper.Unmarshal(RedisConfig); err != nil {
// 			panic(err)
// 		}
// 	})
// }
