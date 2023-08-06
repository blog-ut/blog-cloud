// Package conf
// Time : 2023/8/6 12:49
// Author : PTJ
// File : conf
// Software: GoLand
package conf

import (
	"github.com/spf13/viper"
	"os"
)

func init() {
	InitConf()
}
func InitConf() {
	dir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir + "/conf")
	if err := viper.ReadInConfig(); err != nil {
		return
	}
}
