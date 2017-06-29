package payroll

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"rabbit/route"
)

// Index - serve the index page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Payroll!\n")
}

func init() {
	route.Router.GET("/payroll", Index)
}
