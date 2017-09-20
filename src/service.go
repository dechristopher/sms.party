package main

import (
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"

	a "github.com/dechristopher/sms.party/src/api"
	s "github.com/dechristopher/sms.party/src/strings"
	u "github.com/dechristopher/sms.party/src/util"
)

func main() {
	// Read configuration from config.json
	u.Conf = u.ReadConfig()

	// Kill service if Redis connection fails
	if !u.RedisConnected {
		return
	}

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
	// POST /sms/batch sends a group of messages all sent to the API at once to conserve bandwidth
	// mux.Post(u.Conf.Prefix+"/sms/batch", http.HandlerFunc(BatchHandler))

	// POST /sms/cast sends a single SMS message to multiple numbers
	// mux.Post(u.Conf.Prefix+"/sms/cast", http.HandlerFunc(CastHandler))

	// POST /sms/send sends a single SMS message to a single number
	mux.Post(u.Conf.Prefix+"/sms/send", http.HandlerFunc(a.SendHandler))

	// GET /sms returns information about total app sms statistics
	mux.Get(u.Conf.Prefix+"/sms", http.HandlerFunc(a.SendHandler))

	// Helper Endpoints
	// GET /key gets information and stats about a given API key
	mux.Get(u.Conf.Prefix+"/key", http.HandlerFunc(a.SendHandler))

	// POST /key generates a new API key with given information
	mux.Post(u.Conf.Prefix+"/key", http.HandlerFunc(a.SendHandler))

	// Webserver Endpoints
	// GET / is the homepage (also the only page)
	mux.Get("/", http.HandlerFunc(a.IndexHandler))

	// New Negroni Instance
	n := negroni.New(
		// Middlewares
		negroni.HandlerFunc(a.IPLogMiddleware),
		negroni.HandlerFunc(a.AuthMiddleware),

		// Logger
		negroni.NewLogger(),

		// Serve static HTML on GET /
		negroni.NewStatic(http.Dir("web")),
	)

	n.UseHandler(mux)

	fmt.Println(s.LogPrefix + "> Awaiting queries...")
	n.Run(":" + u.Conf.Port)
}
