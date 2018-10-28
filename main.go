package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
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

// Basic checking on string input
func sValidation(u string) bool {
	sCheck, _ := regexp.MatchString("^[a-zA-Z]+$", u)
	if len(u) < 30 && sCheck {
		return true
	}
	return false
}

//localhost:8080/?name=bob
func greeter(w http.ResponseWriter, r *http.Request) {
	var name string
	if sValidation(r.URL.Query().Get("name")) {
		name = r.URL.Query().Get("name")
	} else {
		name = "invalid"
	}
	log.Println("/ visited by:" + r.RemoteAddr)
	log.Println("Name:", name)
	message := "Hello " + name
	w.Write([]byte(message))
}

func main() {
	log.Printf("Starting on %s\n", server)
	http.HandleFunc("/", greeter)
	if err := http.ListenAndServe(server, nil); err != nil {
		panic(err)
	}
}
