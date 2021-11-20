package config

import (
	"TEST/bibi_test/internal/config/client"
	"TEST/bibi_test/internal/config/db"
	"TEST/bibi_test/internal/config/logger"
	"TEST/bibi_test/internal/config/server"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type config struct {
	Client   client.ClientList
	Logger   logger.Logger
	Server   server.ServerList
	Database db.DatabaseList
}

var cfg config

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	var err error

	viper.AddConfigPath(basepath + "/logger")
	viper.SetConfigType("yaml")
	viper.SetConfigName("logger.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load log config: %v", err))
	}

	viper.AddConfigPath(basepath + "/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("server.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load client config: %v", err))
	}
	viper.AddConfigPath(basepath + "/client")
	viper.SetConfigType("yaml")
	viper.SetConfigName("client.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load client config: %v", err))
	}

	viper.AddConfigPath(basepath + "/db")
	viper.SetConfigType("yaml")
	viper.SetConfigName("database.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load log config: %v", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	viper.Unmarshal(&cfg)

	fmt.Println("============================")
	fmt.Println(Stringify(cfg))
	fmt.Println("============================")
}

func GetConfig() *config {
	return &cfg
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}

func Stringify(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
