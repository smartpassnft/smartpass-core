package storage

import (
  "log"

  "github.com/rapidloop/skv"
)

// Helpers
type Wallet struct{
  Address string
  nft   []string
}

// Returns Bool
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
  return wallet.Address
}

// Store NFT and wallet relation
func StoreNFT(uuid string, address string, store *skv.KVStore) {
  err := store.Put(uuid, address)
  if (err != nil) {
    log.Fatal(err)
  }
}
