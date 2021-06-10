package main

import (
  "os"
  "log"
  "time"
  "net/http"
  // gova "github.com/trvon/govalanche"
  "github.com/gorilla/sessions"
  "github.com/gorilla/mux"
  "github.com/skip2/go-qrcode"
  "math/rand"
  "strings"
)


// This needs a rewritten :3
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
type Middleware func(http.HandlerFunc) http.HandlerFunc

func main() {
  r := mux.NewRouter()
  r.Use(mux.CORSMethodMiddleware(r))
  srv := &http.Server{
    Handler: r,
    Addr: "0.0.0.0:80",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  r.HandleFunc("/user", UserHandler)
  r.HandleFunc("/login", LoginHandler)
  r.HandleFunc("/contract", ContractHandler)
  r.HandleFunc("/avalanche", SmartContractHandler)

  log.Fatal(srv.ListenAndServe())
}

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "session-name")
  // This is a c-chain address
  // publicAddr := gova.GetAddress(nil, nil, session.Values["publicAddr"])
  // session.Values["publicAddr"] = string(publicAddr)
  // log.Print(publicAddr)
  session.Values["token"] = randomString()
  session.Save(r, w)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
  // vars := mux.Vars(r)

}

func SmartContractHandler(w http.ResponseWriter, r *http.Request) {
  // Handle this better
  // backend, err := ethclient.Dial("http://127.0.0.1:9650/ext/ipcs")
}

func ContractHandler(w http.ResponseWriter, r *http.Request) {
  

}

// TODO: Check to make sure this supports
/* 
  References
  https://github.com/ethereum/EIPs/pull/831
  https://eips.ethereum.org/EIPS/eip-831
*/
func ContractUri(method string) string {
  smartContractAddress := ""
  //TODO: Deduce value or execute smart contract function
  // https://github.com/ethereum/EIPs/issues/67
  execution := "method=usetoken"
  uri := "https://avouch.dev/avalanche/avalanche:" + smartContractAddress + "?" + execution
  return uri
}

func GenQR() string {
  // TODO: Implement method for ContractUri function
  method := ""
  var png []byte

  contract := ContractUri(method)
  png, err := qrcode.Encode(contract, qrcode.Medium, 256)
  if ( err != nil ){
    log.Fatal(err)
  }

  file := "/tmp/qr" + randomString() + ".png"
  err = qrcode.WriteFile("https://avouch.dev", qrcode.Medium, 256, file)
  if ( err != nil ){
    log.Fatal(err)
  }

  // TODO: Remove
  log.Print(png)
  // TODO: Generate NFT with generated file
  return file
}
