package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/wavefronthq/go-metrics-wavefront"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"
)

var wfProxy string

func init() {
	// Configure Wavefront proxy address
	if os.Getenv("WF_PROXY") != "" {
		wfProxy = os.Getenv("WF_PROXY")
	} else {
		log.Println("No Wavefront Proxy Address Specified")
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

type Server struct {
	greetC metrics.Counter
}

//localhost:8080/?name=bob
func (s *Server)greeter(w http.ResponseWriter, r *http.Request) {
	s.greetC.Inc(1)
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
	server := Server{metrics.NewCounter()}
	hostTags := map[string]string{
		"source": "j-go-metrics-test",
	}
	wavefront.RegisterMetric("requests", server.greetC, hostTags)
	wfAddr, err := net.ResolveTCPAddr("tcp", wfProxy)
	if err != nil {
		fmt.Println("wf proxy resolve address error:", err)
	}
	go wavefront.WavefrontProxy(metrics.DefaultRegistry, 1*time.Minute, hostTags, "some.prefix", wfAddr)
	log.Printf("Starting server.")
	http.HandleFunc("/", server.greeter)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
