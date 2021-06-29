package storage

// Docs https://gorm.io/docs/index.html
import (
	"errors"
	"log"

	helper "github.com/smartpassnft/smartpass-core/helper"
	"gorm.io/gorm"
)

/*
  Structure (helper.User)
  	Pubkey           string
	  TokenHash        string
	  CreatedAt        time.Time
  	UpdatedAt        time.Time

  Structure (helper.Ticket)
  	UUID   string
	  Pubkey string
	  Status int

  Values for notification
    0 - none
    1 - active query
    10 - yes
    20 - deny
*/

// Set notification value
func CreateTicket(pubkey string, uuid string, value int) {
	db := helper.OpenDB()
	ticket := helper.Ticket{UUID: uuid, Pubkey: pubkey, Status: value}
	result := db.Create(&ticket)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatal("error saving NFT")
	}
}

// Update notification value
func UpdateNotification(pubkey string, uuid string, value int) {
	db := helper.OpenDB()
	query := "UUID = " + uuid
	db.Model(&helper.Ticket{}).Where(query, true).Update("Value", value)
}

// Query notification status
func QueryNotification(pubkey string, uuid string) int {
	db := helper.OpenDB()
	ticket := db.First(&helper.Ticket{}, uuid)
	if errors.Is(ticket.Error, gorm.ErrRecordNotFound) {
		log.Fatal("not found")
	}
	ticket.Get("Status")

	// val, _ := ticket.Get("Status")

	// TODO Fix return
	return 1
	// return val.RowsAffected
}

// Query User
func QueryUser(pubkey string) bool {
	db := helper.OpenDB()
	user := db.First(&helper.User{}, pubkey)
	if errors.Is(user.Error, gorm.ErrRecordNotFound) {
		log.Fatal("not found")
		return false
	}
	return true
}

/*
  NFT Wallet Functionality
*/
func Exists(uuid string) bool {
	db := helper.OpenDB()
	ticket := db.First(&helper.Ticket{}, uuid)
	return !errors.Is(ticket.Error, gorm.ErrRecordNotFound)
}

func GetNFTOwner(uuid string) string {
	db := helper.OpenDB()
	ticket := db.First(&helper.Ticket{}, uuid)
	if errors.Is(ticket.Error, gorm.ErrRecordNotFound) {
		log.Fatal("not found")
	}
	return "TODO fix"
}
