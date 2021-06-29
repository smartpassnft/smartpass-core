package storage

// Docs https://gorm.io/docs/index.html
import (
	"log"

	helper "github.com/smartpassnft/smartpass-core/helper"
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
	if result.Error {
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
	if ticket.Error {
		log.Fatal("not found")
	}
	val, _ := ticket.Get("Status")
	return val
}

// Query User
func QueryUser(pubkey string) bool {
	db := helper.OpenDB()
	user := db.First(&helper.User{}, pubkey)
	if user.Error {
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
	if ticket.Error {
		return false
	}
	return true
}
