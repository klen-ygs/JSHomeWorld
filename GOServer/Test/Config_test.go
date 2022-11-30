package Test

import (
	"github.com/spf13/viper"
	"testing"
)

func TestFileConfig(t *testing.T) {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	testing := map[string]interface{}{
		"port":      "8000",
		"localhost": "127.0.0.1",
		"number":    "66666",
		"string":    "testString",
	}
	for k, v := range testing {
		if v != config.GetString(k) {
			t.Error("want:", v, "having:", config.GetString(k))
		}
	}
}
