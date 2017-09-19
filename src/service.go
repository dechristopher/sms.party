package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"

	s "github.com/dechristopher/sms.party/src/strings"

	u "git.tetra.vodka/csms/csms-backend/src/util"
)

func main() {
	// Read configuration from config.json
	u.Conf = u.ReadConfig()

	fmt.Println("~ Init sms.party backend v" + u.Conf.Version)

	// Do port checks
	if u.Conf.Port == "" {
		//Bind to default port 80
		fmt.Println(s.InfoDefaultPort)
		u.Conf.Port = "80"
	}

	fmt.Println("[sms.p] > Starting...")

	// Instantiate new bone mux
	mux := bone.New()

	// API Endpoints
	mux.Post(u.Conf.Prefix+"/batch", http.HandlerFunc(BatchHandler))
	mux.Post(u.Conf.Prefix+"/cast", http.HandlerFunc(CastHandler))
	mux.Post(u.Conf.Prefix+"/send", http.HandlerFunc(SendHandler))

	// Helper Endpoints
	mux.Get(u.Conf.Prefix+"/key", http.HandlerFunc(CreditHandler))
	mux.Post(u.Conf.Prefix+"/credit/promo", http.HandlerFunc(PromoHandler))

	//Webserver Endpoints
	mux.Get("/", http.HandlerFunc(IndexHandler))

	// New Negroni Instance
	n := negroni.New(
		// Middlewares
		negroni.HandlerFunc(IPLogMiddleware),
		negroni.HandlerFunc(AuthMiddleware),

		// Logger
		negroni.NewLogger(),

		// Static dirs. FUCK static dirs.
		negroni.NewStatic(http.Dir("static")),
		//negroni.NewStatic(http.Dir("images")),
	)

	n.UseHandler(mux)
	n.Run(":" + u.Conf.Port)

	fmt.Println("[sms.p] > Awaiting queries...")
}

/*
	Sends a single message to the target recipient

	200 OK if spooled properly
	Otherwise something went wrong
*/
// SendHandler handles basic single SMS sending
func SendHandler(w http.ResponseWriter, r *http.Request) {
	key := d.APIKey(r.Header.Get("apikey"))

	//Send message
	message := u.Send{Number: sms.Target, Message: sms.Message, MessageID: msgID, Timestamp: int32(time.Now().Unix())}
	payload, _ := json.Marshal(message)

	u.Cli.SendMessage("csms/msg/"+dev.UUID, string(payload))

	u.Okay(w)
}
