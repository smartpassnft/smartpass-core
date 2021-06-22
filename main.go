package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rapidloop/skv"
	"github.com/skip2/go-qrcode"
	assets "github.com/smartpassnft/goavx/avm/assets"
	utils "github.com/smartpassnft/goavx/avm/utils"
	storage "github.com/smartpassnft/smartpass-core/storage"
)

// Helper Variables
var store, err = skv.Open("log/Store.db")

func main() {
	if err != nil {
		log.Fatal(err)
	}

	var wait time.Duration

	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/user", UserHandler)
	r.HandleFunc("/market", MarketHandler)
	r.HandleFunc("/nft/mint/{params}", NFTMintHandler)
	r.HandleFunc("/nft/query/{UUID}", NFTQueryHandler)
	r.HandleFunc("/nft/sell/{params}", NFTSellHandler)
	r.HandleFunc("/nft/id/{UUID}", NFTIDHandler)
	r.HandleFunc("/rpc", RPCHandler)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("shutting down")
}

/*
  User Functionality
*/
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
}

/*
  Ticket Functionality
*/
func NFTIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["UUID"]
	wallet := ""
	// Get wallet address tied to NFT
	if storage.Exists(uuid, store) {
		wallet = storage.GetWallet(uuid, store)
		// Send Notification
	}
	// TODO: Remove when can retrieve wallet
	log.Print(wallet)
}

func NFTMintHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Iron out parameters for mint function
	// vars := mux.Vars(r)
	uri := utils.URI{Address: "", Port: ""}
	// payload := goavx.avm.assets.CreateNFTPayload()
	var payload utils.Payload
	assets.CreateNFTAsset(payload, uri)
}

func NFTQueryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["UUID"]
	if storage.Exists(uuid, store) {
		// Add some functionality here
		log.Print("exists")
	}
}

func NFTSellHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Also change ownership in storage
	// vars := mux.Vars(r)
}

func QRCodeUri(method string) string {
	UUID := uuid.New().String()
	// uri := "https://smartpass.link/nft/id/" + UUID
	uri := "https://127.0.0.1:8000/nft/id/" + UUID
	return uri
}

/*
  Random string for storing image
  TODO: Change to include minted family
*/
func randomString() string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder
	charSet := "abcdedfghipqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length := 30
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

// Handles generation of QR code
func GenQR() string {
	// TODO: Implement method for QRCodeUri function
	method := ""
	var png []byte

	contract := QRCodeUri(method)
	png, err := qrcode.Encode(contract, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Implement dynamic storage with a bucket or custom ipfs server
	file := "/tmp/qr" + randomString() + ".png"
	err = qrcode.WriteFile("https://smartpass.link", qrcode.Medium, 256, file)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Remove
	log.Print(png)
	// TODO: Generate NFT with generated file
	return file
}

/*
  Market Functionality
*/
func MarketHandler(w http.ResponseWriter, r *http.Request) {

}

/*
  RPC Functionality
*/
func RPCHandler(w http.ResponseWriter, r *http.Request) {

}
