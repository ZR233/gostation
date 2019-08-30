/*
@Time : 2019-06-26 14:38
@Author : zr
@File : config
@Software: GoLand
*/
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

func Init(configPath string) {

	configName := "config"
	configPathAndName := path.Join(configPath, configName) + ".yaml"

	// 若配置文件不存在，则创建
	if _, err := os.Stat(configPathAndName); err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(configPathAndName)
			if err != nil {
				log.Panic(err)
			}
			_ = file.Close()
		} else {
			log.Panic(err)
		}
	}
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.AddConfigPath(configPath)

	viper.SetDefault("debug", true)
	viper.SetDefault("server", map[string]string{
		"listen": "0.0.0.0",
		"host":   "localhost",
		"port":   "19080",
	})

	viper.SetDefault("log", map[string]string{"path": "log"})

	viper.SetDefault("redis", map[string]string{"host": "192.168.0.3:6379", "password": "asdf*123", "db": "0", "prefix": "wifidig"})
	viper.SetDefault("db", map[string]map[string]string{
		"db1": {
			"host":          "192.168.0.3",
			"port":          "3306",
			"user":          "sa",
			"password":      "asdf*123",
			"database_name": "WifiDig",
			"charset":       "utf8mb4",
		},
		"db_test": {
			"host":          "192.168.0.3",
			"port":          "3306",
			"user":          "sa",
			"password":      "asdf*123",
			"database_name": "WifiDig_Test",
			"charset":       "utf8mb4",
		},
	})

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.WriteConfig()
	if err != nil { // Handle errors reading the config file
		log.Panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		log.Panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
