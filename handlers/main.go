package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
	assets "github.com/smartpassnft/goavx/avm/assets"
	utils "github.com/smartpassnft/goavx/avm/utils"
	"github.com/smartpassnft/smartpass-core/helper"
	helper "github.com/smartpassnft/smartpass-core/helper"
	"github.com/smartpassnft/smartpass-core/storage"
	"golang.org/x/crypto/bcrypt"
)

/*
  User Functionality
*/
func UserStatusHandler(w http.ResponseWriter, r *http.Request) {
	var u helper.NQuery

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Fatal(err)
		http.Error(w, "Error decoding reponse object", http.StatusBadRequest)
		return
	}

	/* Data needed for request
	   uuid : uuid
	   pubkey : pubkey
	*/
	status := storage.QueryNotification(u.UUID, u.Pubkey)
	var n = helper.Ticket{Pubkey: u.Pubkey, UUID: u.UUID, Status: status}

	response, err := json.Marshal(&n)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var user helper.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err error
		err = SetError(err, "error in ready body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	if helper.QueryUser(user.Pubkey) {
		// TODO: Change token generating function
		user.Tickets = 
	}

}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {

}

/*
  Ticket Functionality
*/
func NFTIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["UUID"]
	wallet := ""
	// Get wallet address tied to NFT
	if storage.Exists(uuid) {
		wallet = storage.GetNFTOwner(uuid)
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
	if storage.Exists(uuid) {
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
	// uri := "https://" + site + "/nft/id/" + UUID
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

	handler := QRCodeUri(method)
	png, err := qrcode.Encode(handler, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("/tmp/qr")
	if os.IsNotExist(err) {
		err = os.Mkdir("/tmp/qr", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// TODO: Implement dynamic storage with a bucket or custom ipfs server
	file := "/tmp/qr/" + randomString() + ".png"
	err = qrcode.WriteFile(handler, qrcode.Medium, 256, file)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Use PNG data to backup PNG file to IPFS
	// https://github.com/ipfs/go-ipfs/blob/master/docs/examples/go-ipfs-as-a-library/main.go#L240-L243
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
func RPCHandler(w http.ResponseWriter, r *http.Request) {

}
*/
