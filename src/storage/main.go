package storage

import (
  "skv"
  "fmt"
  "log"

  "github.com/rapidloop/skv"
)

// Returns Bool
func Exists(uuid string, store &skv.KVStore) bool {
  err := store.Get(uuid, store)
  if ( err != nil ) {
    return true
  }
  return false
}

// Get Wallet from Map 
func GetWallet(uuid string, store &skv.KVStore) string {
  wallet := store.Get(uuid, store)
  return wallet
}

// Store NFT and wallet relation
func StoreNFT(uuid address string, store &skv.KVStore) {
  err := store.Put(uuid, address)
  if (err) {
    log.Fatal(err)
  }
}
