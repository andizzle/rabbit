package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"rabbit/route"

	_ "rabbit/chat"
	_ "rabbit/payroll"
)

// Index - serve the index page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome\n")
}

func main() {
	route.Router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", route.Router))
}
