package helper

import (
  "fmt"
  "github.com/spf13/viper"
)

func GetRPCServer() string {
  viper.SetConfigName("server")
  viper.AddConfigPath("../config/")
  err := viper.ReadInConfig()
  if err != nil {
    panic(fmt.Errorf("Fatal error config file: %s\n", err))
  }
  rpc := viper.GetString("settings.rpc")
  return rpc
}
