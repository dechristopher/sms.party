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

	/*--- START API ENDPOINTS ---*/
	// POST /sms/batch sends a group of messages all sent to the API at once to conserve bandwidth
	mux.Post(u.Conf.Prefix+"/sms/batch", http.HandlerFunc(a.UnimplementedHandler))
	// POST /sms/cast sends a single SMS message to multiple numbers
	mux.Post(u.Conf.Prefix+"/sms/cast", http.HandlerFunc(a.UnimplementedHandler))
	// POST /sms/send sends a single SMS message to a single number
	mux.Post(u.Conf.Prefix+"/sms/send", http.HandlerFunc(a.SendHandler))
	// GET /sms returns information about total app sms statistics
	mux.Get(u.Conf.Prefix+"/sms", http.HandlerFunc(a.UnimplementedHandler))
	/*--- END API ENDPOINTS ---*/

	/*--- START HELPER ENDPOINTS ---*/
	// GET /key gets information and stats about a given API key
	mux.Get(u.Conf.Prefix+"/key", http.HandlerFunc(a.UnimplementedHandler))
	// POST /key generates a new API key with given information
	mux.Post(u.Conf.Prefix+"/key", http.HandlerFunc(a.UnimplementedHandler))
	// GET /nomsg gets information about a number's status on the do not message list
	mux.Get(u.Conf.Prefix+"/nomsg", http.HandlerFunc(a.UnimplementedHandler))
	// POST /nomsg adds a number to the do not message list
	mux.Post(u.Conf.Prefix+"/nomsg", http.HandlerFunc(a.UnimplementedHandler))
	// GET /number gets information and stats about a given number
	mux.Get(u.Conf.Prefix+"/number", http.HandlerFunc(a.UnimplementedHandler))
	//GET /host returns the container hostname
	mux.Post(u.Conf.Prefix+"/host", http.HandlerFunc(a.HostHandler))
	/*--- END HELPER ENDPOINTS ---*/

	/*--- START WEB ENDPOINTS ---*/
	// GET / is the homepage (also the only page)
	mux.Get("/", http.HandlerFunc(a.IndexHandler))
	/*--- END WEB ENDPOINTS ---*/

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
