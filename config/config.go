package config

type ServerConfig struct {
	Host        string      `mapstructure:"host" json:"host"`
	Port        int         `mapstructure:"port" json:"port"`
	MySQLConfig MySQLConfig `mapstructure:"mysql" json:"mysql"`
	RedisConfig RedisConfig `mapstructure:"redis" json:"redis"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Prefix   string `mapstructure:"prefix" json:"prefix"`
}

type RedisConfig struct {
	Host       string `mapstructure:"host"`
	Password   string `mapstructure:"password"`
	Port       int    `mapstructure:"port"`
	Db         int    `mapstructure:"db"`
	PoolSize   int    `mapstructure:"pool_size"`
	ExpireHour int    `mapstructure:"expire_hour"`
}
