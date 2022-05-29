package controller

//Imports
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//Student struct
type Student struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
}

//Students slice
var students = []Student{
	{"Umar", "Aliyu", 52, "M"},
	{"Kamal", "Ibrahim", 40, "M"},
	{"Hauwa", "Jibril", 30, "F"},
	{"Bilya", "Rabiu", 60, "M"},
}

//Get all students
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	result, _ := json.Marshal(students)
	fmt.Fprint(w, string(result))
}

//Get one student functionality
func GetOne(w http.ResponseWriter, r *http.Request) {
	//Sending header to browser
	w.Header().Add("Content-Type", "Application/json")
	path := r.URL.Path
	paths := strings.Split(path, "/getstudent/")
	id, _ := strconv.Atoi(paths[1])
	result, _ := json.Marshal(students[id])
	fmt.Fprint(w, string(result))
}

//Post student functionality
func PostOne(w http.ResponseWriter, r *http.Request) {
	//Checking if method = post
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		var result Student = Student{}
		//Changing json request to bytes
		json.Unmarshal(body, &result)
		//Adding post request to slice
		students = append(students, result)
		//Changing slice to json before printing
		fmt.Fprint(w, string(body))
	} else {
		fmt.Fprint(w, "Invalid Request")
	}
}

//Remove slice index
func RemoveIndex(s []Student, index int) []Student {
	return append(s[:index], s[index+1:]...)
}

//Delete student functionality
func DeleteOne(w http.ResponseWriter, r *http.Request) {
	//checking if method = delete
	if r.Method == "DELETE" {
		w.Header().Add("Content-Type", "Application/json")
		path := r.URL.Path
		paths := strings.Split(path, "/student/")
		id, _ := strconv.Atoi(paths[1])
		students = RemoveIndex(students, id)
		result, _ := json.Marshal(students)
		fmt.Fprint(w, string(result))
	} else {
		fmt.Fprint(w, "Invalid Request")
	}
}
