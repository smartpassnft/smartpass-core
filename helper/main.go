package helper

import (
  "fmt"
  "github.com/spf13/viper"
)

type Notification struct {
 Pubkey string
 UUID string
 Status int
}

type NQuery struct {
  Pubkey string
  UUID string
}

func GetString(config string, path string, variable string) string {
  viper.SetConfigName(config)
  viper.AddConfigPath(path)
  err := viper.ReadInConfig()
  if err != nil {
    panic(fmt.Errorf("Fatal error config file: %s\n", err))
  }
  return viper.GetString(variable)
}

func GetWebServer() string {
  web := GetString("server", "../config", "web.web")
  return web
}

func GetRPCServer() string {
  rpc := GetString("server", "../config", "web.rpc")
  return rpc
}

func GetUserStorage() string {
  user := GetString("server", "../config", "storage.user")
  return "../" + user
}

func GetNFTStorage() string {
  nft := GetString("server", "../config", "storage.nft")
  // TODO: Clean up
  return "../" + nft
}
