package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{param}", NameHendler).Methods(http.MethodGet)
	router.HandleFunc("/bad", BadHendler).Methods(http.MethodGet)
	router.HandleFunc("/data", DataHendler).Methods(http.MethodPost)
	router.HandleFunc("/headers", HeaderHendler).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func NameHendler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["param"]
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hello, %s", param)
}

func BadHendler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func DataHendler(w http.ResponseWriter, r *http.Request) {
	param, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}
	r.Body.Close()
	w.WriteHeader(200)
	fmt.Fprintf(w, "I got message:\n %s", string(param))
}

func HeaderHendler(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	valA, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(500)
	}
	valB, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(500)
	}
	response := strconv.Itoa(valA + valB)
	w.Header().Set("a+b", response)
	w.WriteHeader(200)
}
