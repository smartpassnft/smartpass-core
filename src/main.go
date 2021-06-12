package main

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
	// "github.com/smartpass/v1/queue"
)

func main() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/user", UserHandler)
	r.HandleFunc("/market", MarketHandler)
	r.HandleFunc("/avalanche", QRCodeHandler)
	r.HandleFunc("/ticket", TicketHandler)
	r.HandleFunc("/rpc", RPCHandler)

	log.Fatal(srv.ListenAndServe())
}

/*
  User Functionality
*/
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
}

// Should maintain an asynchronous connection to send notifications to users
func TicketHandler(w http.ResponseWriter, r *http.Request) {
	// Handle this better
	// backend, err := ethclient.Dial("http://127.0.0.1:9650/ext/ipcs")
}

/*
  Ticket Functionality
*/
func QRCodeUri(method string) string {
	UUID := uuid.New().String()
	uri := "https://smartpass.link/avalanche/avalanche:" + UUID
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
