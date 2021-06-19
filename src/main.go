package main

import (
	"os"
	"log"
	"fmt"
	"time"
	"context"
	"strings"
	"os/signal"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
)

func main() {
  var wait time.Duration
  fmt.Print(uuid.New().String())
  r := mux.NewRouter()
  // r.Use(mux.CORSMethodMiddleware(r))
  srv := &http.Server{
  	Handler:      r,
  	Addr:         "0.0.0.0:8000",
  	WriteTimeout: 15 * time.Second,
  	ReadTimeout:  15 * time.Second,
  }
  
  r.HandleFunc("/user", UserHandler)
  r.HandleFunc("/market", MarketHandler)
  r.HandleFunc("/handler", QRCodeHandler)
  r.HandleFunc("/nft/mint/{params}", NFTMintHandler)
  r.HandleFunc("/nft/query/{params}", NFTQueryHandler)
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
  
  // Block until we receive our signal.
  <-c
  
  // Create a deadline to wait for.
  ctx, cancel := context.WithTimeout(context.Background(), wait)
  defer cancel()
  // Doesn't block if no connections, but will otherwise wait
  // until the timeout deadline.
  srv.Shutdown(ctx)
  // Optionally, you could run srv.Shutdown in a goroutine and block on
  // <-ctx.Done() if your application should wait for other services
  // to finalize based on context cancellation.
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
}
func NFTMintHandler(w http.ResponseWriter, r *http.Request) {
  // vars := mux.Vars(r)	
}
func NFTQueryHandler(w http.ResponseWriter, r *http.Request) {
  // vars := mux.Vars(r)	
}
func NFTSellHandler(w http.ResponseWriter, r *http.Request) {
  // vars := mux.Vars(r)	
}

func QRCodeUri(method string) string {
	UUID := uuid.New().String()
	// uri := "https://smartpass.link/nft/id/" + UUID
	uri := "https://127.0.0.1:8000/nft/id/" + UUID
	return uri
}

// Handles browser view of QR code
func QRCodeHandler(w http.ResponseWriter, r *http.Request) {

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
