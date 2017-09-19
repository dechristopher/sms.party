package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"

	d "github.com/dechristopher/sms.party/src/data"
	s "github.com/dechristopher/sms.party/src/strings"
	u "github.com/dechristopher/sms.party/src/util"
)

func main() {
	// Read configuration from config.json
	u.Conf = u.ReadConfig()

	fmt.Println(s.LogPrefix + s.InfoStartup + u.Conf.Version)

	// Do port checks
	if u.Conf.Port == "" {
		//Bind to default port 80
		fmt.Println(s.LogPrefix + s.InfoDefaultPort)
		u.Conf.Port = "80"
	}

	fmt.Println(s.LogPrefix + "> Starting...")

	// Instantiate new bone mux
	mux := bone.New()

	// API Endpoints
	// mux.Post(u.Conf.Prefix+"/batch", http.HandlerFunc(BatchHandler))
	// mux.Post(u.Conf.Prefix+"/cast", http.HandlerFunc(CastHandler))
	mux.Post(u.Conf.Prefix+"/send", http.HandlerFunc(SendHandler))

	// Helper Endpoints
	mux.Get(u.Conf.Prefix+"/key", http.HandlerFunc(SendHandler))
	mux.Post(u.Conf.Prefix+"/key", http.HandlerFunc(SendHandler))

	//Webserver Endpoints
	mux.Get("/", http.HandlerFunc(IndexHandler))

	// New Negroni Instance
	n := negroni.New(
		// Middlewares
		negroni.HandlerFunc(IPLogMiddleware),
		negroni.HandlerFunc(AuthMiddleware),

		// Logger
		negroni.NewLogger(),

		// Serve static HTML on GET /
		negroni.NewStatic(http.Dir("web")),
	)

	n.UseHandler(mux)

	fmt.Println(s.LogPrefix + "> Awaiting queries...")
	n.Run(":" + u.Conf.Port)
}

/*--- START MIDDLEWARE ---*/

// IPLogMiddleware simply prefixes request IP to negroni logger output for every request
func IPLogMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Printf(s.LogPrefix+"Request from %v -> ", r.RemoteAddr)
	next(w, r)
}

// AuthMiddleware verifies presence and validity of API key header
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	/*
		401 Unauthorized if invalid key or no key and JSON error returned
		{"error": "Invalid sms.party API key"}
	*/
	if r.RequestURI == "/" || strings.Contains(r.RequestURI, "/files") {
		fmt.Println("homepage - no auth")
		next(w, r)
		return
	}

	key := d.APIKey(r.Header.Get("apikey"))
	fmt.Printf("%v - checking auth - %v\n", r.RequestURI, key)

	// Check API key validity
	/*if _, existsErr := u.DBC.GetEmail(key); existsErr != nil {
		u.SendResponse(w, true, 401, `{"error": "`+s.ErrBadAPIKey+`"}`)
		return
	}*/

	next(w, r)
}

/*--- END MIDLEWARE*/

/*--- START HANDLERS ---*/

// IndexHandler serves homepage
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	u.Templates.ExecuteTemplate(w, "index.html", nil)
}

// SendHandler handles basic single SMS sending
func SendHandler(w http.ResponseWriter, r *http.Request) {
	/*
		200 OK if spooled properly
		Otherwise something went wrong
	*/
	key := d.APIKey(r.Header.Get("apikey"))
	fmt.Println(key)

	//Send message...

	u.Okay(w)
}

/*--- END HANDLERS ---*/
