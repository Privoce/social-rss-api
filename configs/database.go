package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// DatabaseCfg  database config struct
type DatabaseCfg struct {
	URL, Username, Password string
}

// ToString return url string
func (dbcfg *DatabaseCfg) ToString() string {
	return fmt.Sprintf("%s:%s@tcp%s?charset=utf8&parseTime=True&loc=Local", dbcfg.Username, dbcfg.Password, dbcfg.URL)
}

// LoadDatabaseCfg 加载新的数据库配置
func LoadDatabaseCfg() *DatabaseCfg {
	var DbCfg DatabaseCfg
	viper.SetConfigName("db") // 设置配置文件名 (不带后缀)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")    // 第一个搜索路径
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
	viper.Unmarshal(&DbCfg) // 将配置信息绑定到结构体上
	return &DbCfg
}
