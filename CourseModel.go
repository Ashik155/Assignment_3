package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Length string `json:"length"`
}

var allcourses []Course

func GetAllCourses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allcourses)

}
func GetCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	from_browser := mux.Vars(r)
	for _, values_in_allcourses := range allcourses {
		if values_in_allcourses.Id == from_browser["id"] {
			json.NewEncoder(w).Encode(values_in_allcourses)
			return
		}

	}

}
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new_course Course
	json.NewDecoder(r.Body).Decode(&new_course)
	new_course.Id = strconv.Itoa(len(allcourses) + 1)
	allcourses = append(allcourses, new_course)
	json.NewEncoder(w).Encode(allcourses)

}
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	from_browser := mux.Vars(r)
	for i, values_in_allcourses := range allcourses {
		if values_in_allcourses.Id == from_browser["id"] {
			allcourses = append(allcourses[:i], allcourses[i+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(allcourses)

}
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	from_browser := mux.Vars(r)
	for i, values_in_allcourses := range allcourses {
		if values_in_allcourses.Id == from_browser["id"] {
			allcourses = append(allcourses[:i], allcourses[i+1:]...)
			var new_updated_course Course
			json.NewDecoder(r.Body).Decode(&new_updated_course)
			new_updated_course.Id = from_browser["id"]
			allcourses = append(allcourses, new_updated_course)
			json.NewEncoder(w).Encode(allcourses)
			return
		}

	}
	json.NewEncoder(w).Encode(allcourses)
}

func main() {
	router := mux.NewRouter()
	allcourses = append(allcourses, Course{Id: "1", Name: "Learn Go quick", Length: "2 hours"})
	router.HandleFunc("/courses", GetAllCourses).Methods("GET")
	router.HandleFunc("/courses/{id}", GetCourse).Methods("GET")
	router.HandleFunc("/course", CreateCourse).Methods("POST")
	router.HandleFunc("/courses/{id}", DeleteCourse).Methods("DELETE")
	router.HandleFunc("/courses/{id}", UpdateCourse).Methods("POST")

	log.Fatal(http.ListenAndServe(":2000", router))
}
