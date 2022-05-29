package main

//Imports
import (
	"net/http"

	"github.com/boskitech/goHttp/routes"
)

func main() {
	//App routes
	routes.Routes()
	//Start Server
	http.ListenAndServe(":4000", nil)
}
