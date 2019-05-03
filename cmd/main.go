package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rcrowley/go-metrics"
	"github.com/wavefronthq/go-metrics-wavefront/reporting"
	"github.com/wavefronthq/wavefront-sdk-go/application"
	wavefront "github.com/wavefronthq/wavefront-sdk-go/senders"
	"log"
	"net/http"
	"os"
	"regexp"
)

var wfProxy string

type Server struct {
	greetC metrics.Counter
	homeC metrics.Counter
	db *sql.DB
}

type User struct {
	Uid int
	Name string
}

var server Server

func init() {
	// Configure Wavefront proxy address
	if os.Getenv("WF_PROXY") != "" {
		wfProxy = os.Getenv("WF_PROXY")
		log.Println("WFProxy:", wfProxy)
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

func (s *Server)healthz(w http.ResponseWriter, r *http.Request) {
	s.greetC.Inc(1)
	log.Println("/healthz visited by:" + r.RemoteAddr)
	message := "Health: OK"
	w.Write([]byte(message))
}

// Temporary route to try clarity and react
func (s *Server)home(w http.ResponseWriter, r *http.Request) {
	s.homeC.Inc(1)
	file := "public/home/index.html"
	log.Println("Home File:", file)
	http.ServeFile(w, r, file)
}

// Route to serve javascript
func (s *Server)homejs(w http.ResponseWriter, r *http.Request) {
	file := "public/home/js/" + r.URL.Path[len("/js/"):]
	log.Println("JS File:", file)
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, file)
}
// Route to serve images
func (s *Server)images(w http.ResponseWriter, r *http.Request) {
	file := "public/home/images/" + r.URL.Path[len("/images/"):]
	log.Println("Image File:", file)
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, file)
}

func (s *Server)users(w http.ResponseWriter, r *http.Request) {
	file := "public/users/index.html"
	log.Println("Users File:", file)
	http.ServeFile(w, r, file)
}

// Route to serve javascript
func (s *Server)usersjs(w http.ResponseWriter, r *http.Request) {
	file := "public/users/js/" + r.URL.Path[len("/users/js/"):]
	log.Println("JS File:", file)
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, file)
}

func (s *Server)testSQL(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query(`SELECT * FROM webusers`)
	if err != nil {
		fmt.Println("PSQL Error:", err)
	}
	var users []User
	for rows.Next() {
		var uid int
		var username string
		err = rows.Scan(&uid, &username)
		user := User{uid, username}
		users = append(users, user)
	}

	response, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // TODO may need for CORS
	w.Write(response)
}

func main() {
	//wavefront setup
	proxyCfg := &wavefront.ProxyConfiguration {
		Host : wfProxy,

		// At least one port should be set below.
		MetricsPort : 8082,      // set this (typically 2878) to send metrics
		DistributionPort: 8082,  // set this (typically 2878) to send distributions
		TracingPort : 30000,     // set this to send tracing spans

		FlushIntervalSeconds: 10, // flush the buffer periodically, defaults to 5 seconds.
	}
	sender, errSender := wavefront.NewProxySender(proxyCfg)
	if errSender != nil {
		panic("error:" + errSender.Error())
	}
	_ = reporting.NewReporter(
		sender,
		application.New("app", "srv"),
		reporting.Source("j-go-metrics-test"),
		reporting.Prefix("jsome.jprefix"),
		reporting.LogErrors(true),
	)

	// Create monitoring metrics
	server := Server{metrics.NewCounter(), metrics.NewCounter(), nil}

	errHealthz := metrics.Register("healthz.requests", server.greetC)
	if errHealthz != nil {
		fmt.Println("Healthz register metric error: ", errHealthz)
	}

	errHGreetC :=  metrics.Register("home.requests", server.homeC)
	if errHGreetC != nil {
		fmt.Println("Greetc register metric error: ", errHGreetC)
	}

	// Verify metrics are in registry
	metrics.DefaultRegistry.Each(func(key string, metric interface{}) {
		log.Println("metric:", key)
	})


	// Connect to PSQL Database
	/*
	connStr := "postgres://postgres:mysecretpassword@postgres:5432/webapp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	server.db = db
	*/

	// Webserver setup
	log.Printf("Starting server.")
	//http.HandleFunc("/", server.greeter) // TODO old remove
	http.HandleFunc("/", server.home)
	http.HandleFunc("/js/", server.homejs)
	http.HandleFunc("/images/", server.images)
	http.HandleFunc("/healthz", server.healthz)
	http.HandleFunc("/api/v1/users/", server.testSQL)
	http.HandleFunc("/users", server.users)
	http.HandleFunc("/users/js/", server.usersjs)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
