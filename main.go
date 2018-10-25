package main

import (
	"log"
	"net/http"
	"os"
)

var server string

func init() {
	// Configure Wavefront proxy address
	if os.Getenv("WF_PROXY") != "" {
		server = os.Getenv("WF_PROXY")
	} else {
		log.Println("No Wavefront Proxy Address Specified")
	}
	// Configure webserver listening address
	if os.Getenv("SERVER") != "" {
		server = os.Getenv("SERVER")
	} else {
		log.Panic("No Server Address Specified")
	}
}

//localhost:8080/?name=bob
func greeter(w http.ResponseWriter, r *http.Request) {
	log.Println("/ visited by:" + r.RemoteAddr)
	name := r.URL.Query().Get("name")
	log.Println("Name:", name)
	message := "Hello " + name
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", greeter)
	if err := http.ListenAndServe(server, nil); err != nil {
		panic(err)
	}
}
