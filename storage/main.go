package storage

import (
  "log"

  "github.com/rapidloop/skv"
)

// Helpers
type Wallet struct{
  Pubkey string
  nft   []string
}

type User struct{
  Pubkey string
  Value map[string]int
}

/*
  User Storage

  Values for notification
  0 - none
  1 - active query
  10 - yes
  20 - deny
*/
func SetNotification(pubkey string, uuid string, u *skv.KVStore, value int){
  // TODO: Change for map so that values can be updated
  val := map[string]int{
    uuid : value,
  }
  err := u.Put(pubkey, val)
  if err != nil {
    log.Fatal(err)
  }
}

// Update user value of use
func UpdateNotification(pubkey string, uuid string, u *skv.KVStore, value int) {
  var user User 
  u.Get(pubkey, &user)
  user.Value[uuid] = value
  err := u.Put(pubkey, user)
  if err != nil {
    log.Fatal(err)
  }
}

// Query notification status
func QueryNotification(pubkey string, uuid string, u *skv.KVStore) int {
  var user User
  // user storage
  u.Get(pubkey, &user)
  return user.Value[uuid]
}

/*
  NFT Wallet Functionality
*/
func Exists(uuid string, store *skv.KVStore) bool {
  err := store.Get(uuid, store)
  if ( err != nil ) {
    return true
  }
  return false
}

// Get Wallet from Map 
func GetWallet(uuid string, store *skv.KVStore) string {
  var wallet Wallet
  // TODO: Test where value is being stored from passed interface
  store.Get(uuid, &wallet)
  return wallet.Pubkey
}

// Store NFT and wallet relation
func StoreNFT(uuid string, pubkey string, store *skv.KVStore, u *skv.KVStore) {
  err := store.Put(uuid, pubkey)
  if (err != nil) {
    log.Fatal(err)
  }
  SetNotification(pubkey, uuid, u, 0)
}
