package routes

//Imports
import (
	"net/http"

	"github.com/boskitech/goHttp/controller"
)

func Routes() {
	//Get all students route
	http.HandleFunc("/students", controller.GetAll)

	//Get one student route
	http.HandleFunc("/getstudent/", controller.GetOne)

	//Post student route
	http.HandleFunc("/post/", controller.PostOne)

	//Delete student route
	http.HandleFunc("/student/", controller.DeleteOne)
}
