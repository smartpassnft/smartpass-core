package queue

import (
	"log"

	assets "github.com/smartpassnft/goavx/avm/assets"
	utils "github.com/smartpassnft/goavx/avm/utils"
	/*
	  "github.com/smartpassnft/goavx/avm"
	  "google.golang.org/grpc"
	*/)

// Possible API interfaces to use
type Connections interface {
	AVMRequest()
}

// TODO: Add functionality for other chains here
func RequestBuilder() utils.URI {
  address := helper.GetRPCAddress()
	var uri = utils.URI{Address: address, Port: "10000"}
	return uri
}

// TODO: Implement Queue function here
// Handling Incoming Connections
func AVMRequest() {

}

// Build NFT Requests
func NFTMint(template utils.Payload, request string) {
	var uri = RequestBuilder()
	switch request {
	case "create":
		assets.CreateNFTAsset(template, uri)
	case "mint":
		assets.MintNFTAsset(template, uri)
	case "send":
		assets.SendNFT(template, uri)
	default:
		log.Fatal("unknown request")
	}
}

// TODO: Initialize wallet
// Init Wallet Address
func InitWallet() {

}
