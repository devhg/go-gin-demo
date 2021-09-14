package config

type Redis struct {
	// common config use dns name load balance
	Service serviceConfig `toml:"service"`

	// single redis service
	Config struct {
		Host        string `toml:"Host"`
		Password    string `toml:"Password"`
		DB          string `toml:"Db"`
		MaxIdle     int    `toml:"MaxIdle"`
		MaxActive   int    `toml:"MaxActive"`
		IdleTimeout int    `toml:"IdleTimeout"`
		MaxRetries  int    `toml:"MaxRetries"`
		PoolSize    int    `toml:"PoolSize"`
	}
}

var RedisConfig = &Redis{}
