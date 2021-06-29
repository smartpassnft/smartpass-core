package helper

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Pubkey    string    `json:"pubkey" sql:"pubkey"`
	TokenHash string    `json:"tokenhash" sql:"tokenhash"`
	CreatedAt time.Time `json:"createdat" sql:"createdat"`
	UpdatedAt time.Time `json:"updatedat" sql:"updatedat"`
}

type Ticket struct {
	UUID   string
	Pubkey string
	Status int
}

type NQuery struct {
	Pubkey string
	UUID   string
}

func GetString(config string, path string, variable string) string {
	viper.SetConfigName(config)
	viper.AddConfigPath(path)
	// viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	return viper.GetString(variable)
}

func GetWebServer() string {
	web := GetString("server", "./config", "web.web")
	return web
}

func GetRPCServer() string {
	rpc := GetString("server", "./config", "web.rpc")
	return rpc
}

func GetUserDB() string {
	db := GetString("server", "./config", "db.host")
	return db
}

func OpenDB() gorm.DB {
	// Helper variables
	user := GetString("server", "./config", "db.user")
	host := GetString("server", "./config", "db.host")
	password := GetString("server", "./config", "db.password")
	db := GetString("server", "./config", "db.dbname")
	port := GetString("server", "./config", "db.port")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + db + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config)
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}
