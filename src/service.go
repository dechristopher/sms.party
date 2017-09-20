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
	// mux.Post(u.Conf.Prefix+"/batch", http.HandlerFunc(BatchHandler))
	// mux.Post(u.Conf.Prefix+"/cast", http.HandlerFunc(CastHandler))
	mux.Post(u.Conf.Prefix+"/send", http.HandlerFunc(a.SendHandler))

	// Helper Endpoints
	mux.Get(u.Conf.Prefix+"/key", http.HandlerFunc(a.SendHandler))
	mux.Post(u.Conf.Prefix+"/key", http.HandlerFunc(a.SendHandler))

	// Webserver Endpoints
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
