package api

import (
	"net/http"

	u "github.com/dechristopher/sms.party/src/util"
)

// IndexHandler serves homepage
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	u.Templates.ExecuteTemplate(w, "index.html", nil)
}
