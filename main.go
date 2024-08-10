package main

// import (
// 	"fmt"

// 	"rsc.io/quote"
// )

// func main() {
// 	fmt.Println(
// 		"Hello, World!",
// 	)
// 	fmt.Println(quote.Go())
// }

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	details "github.com/vikassfteng/go-microservices/details"
)

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {

// 	http.HandleFunc("/", rootHandler)

// 	Fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", Fs))
// 	log.Println("web server started")
// 	http.ListenAndServe(":80", nil)
// }

func heatlhHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking application health")
	reponse := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(reponse)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Applicaiton is rnnning")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIp()
	fmt.Println(hostname, IP)
	reponse := map[string]string{
		"Hostname": hostname,
		"IP":       IP.String(),
	}
	json.NewEncoder(w).Encode(reponse)

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/health", heatlhHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}
