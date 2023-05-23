package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Course struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Price string `json: "price"`
	// เปลี่ยนชื่อ key
}

var CourseList []Course

func init() {
	CourseJson := `[
		{
			"id" : 1,
			"name": "Java",
			"price": "1000"
		},
		{
			"id" : 2,
			"name": "C#",
			"price": "2000"
		},
		{
			"id" : 3,
			"name": "Python",
			"price": "3000"
		}
	]`
	err := json.Unmarshal([]byte(CourseJson), &CourseList) // นำ json ไปใส่ใน coureslist
	if err != nil {
		log.Fatal(err)
	}
}
func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(bodybyte, newCourse)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
