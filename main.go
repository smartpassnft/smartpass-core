package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	handler "github.com/smartpassnft/smartpass-core/handlers"
)

func main() {
	var wait time.Duration

	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Registration Handlers
	postRequest := r.Methods(http.MethodPost).Subrouter()
	postRequest.HandleFunc("/user/login", handler.UserLoginHandler)
	// r.HandleFunc("/user/status", UserStatusHandler)
	postRequest.HandleFunc("/user/logout", handler.UserLogoutHandler)
	postRequest.Use()

	// Market Functionality
	// TODO: Update to handle market functionality
	r.HandleFunc("/market", handler.MarketHandler)
	r.HandleFunc("/nft/sell/{params}", handler.NFTSellHandler)

	// Query Handler
	getRequest := r.Methods(http.MethodGet).Subrouter()
	getRequest.HandleFunc("/user", handler.UserHandler)
	getRequest.HandleFunc("/nft/mint/{params}", handler.NFTMintHandler)
	getRequest.HandleFunc("/nft/query/{UUID}", handler.NFTQueryHandler)
	getRequest.HandleFunc("/nft/id/{UUID}", handler.NFTIDHandler)

	// r.HandleFunc("/rpc", RPCHandler)

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
