package Config

import "github.com/spf13/viper"

var config *viper.Viper

func GetInt(key string) int {
	return config.GetInt(key)
}

func GetString(key string) string {
	return config.GetString(key)
}

func Get() *viper.Viper {
	return config
}

func init() {
	v := viper.New()
	v.AddConfigPath("./Config/")
	v.SetConfigName("Config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err) // 读取配置文件失败
	}
	config = v

}
