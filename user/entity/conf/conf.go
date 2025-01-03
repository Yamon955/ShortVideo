package conf

import (
	"github.com/Yamon955/ShortVideo/user/entity/errcode"
	"github.com/spf13/viper"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/log"
)

type AppConfig struct {
	UserDefaultConf UserDefaultConfig `mapstructure:"user_default_conf"`
	MySQLConf       MySQLConfig       `mapstructure:"mysql"`
}

type UserDefaultConfig struct {
	Avator string `mapstructure:"avator"`
	Sign   string `mapstructure:"sign"`
}

type MySQLConfig struct {
	Address      string `mapstructure:"address"`
	User         string `mapstructure:"user"`
	Pwd          string `mapstructure:"pwd"`
	DBName       string `mapstructure:"db_name"`
	Port         string `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

var AppConf *AppConfig

func Init() error {
	viper.SetConfigFile("./conf/app_config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("read config failed, err:%v", err)
		return errs.New(errcode.ErrReadConfig, "viper read config err")
	}
	if err := viper.Unmarshal(&AppConf); err != nil {
		log.Errorf("read config failed, err:%v", err)
		return errs.New(errcode.ErrUnmarshalConfig, "viper unmarshal config err")
	}
	log.Infof("app config:%+v", *AppConf)
	return nil
}
