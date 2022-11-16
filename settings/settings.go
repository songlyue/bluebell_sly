package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name              string `mapstructure:"name"`
	Mode              string `mapstructure:"mode"`
	Port              string `mapstructure:"port"`
	*LogConfig        `mapstructure:"log"`
	*PostgresqlConfig `mapstructure:"postgresql"`
	*RedisConfig      `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type PostgresqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	PoolSize     int    `mapstructure:"poolSize"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

var Conf = new(AppConfig)

func Init() error {
	viper.SetConfigFile("./conf/config.yaml")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了")
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("viper读取配置文件失败"))
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Println("反序列失败，err：", err)
	}

	return err
}
