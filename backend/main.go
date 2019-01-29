package main

import (
	"encoding/json"
	"fmt"
	"gorm/src/backend/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type tempComment struct {
	Name    string    `json:"name"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func main() {
	models.InitDB("root:@tcp(127.0.0.1:3306)/carcomments?parseTime=true")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/cars", allCars).Methods("GET")
	myRouter.HandleFunc("/brand", allBrands).Methods("GET")
	myRouter.HandleFunc("/cars/brand/{name}", allCarsByBrand).Methods("GET")
	myRouter.HandleFunc("/cars/contains/{name}", contains).Methods("POST")
	myRouter.HandleFunc("/comments", allComments).Methods("GET")
	myRouter.HandleFunc("/comments/add", addComment).Methods("POST")
	log.Fatal(http.ListenAndServe(":3100", myRouter))

}

func allBrands(w http.ResponseWriter, r *http.Request) {
	brands := models.AllBrands()
	fmt.Println(brands)

	json.NewEncoder(w).Encode(brands)
}

func allCars(w http.ResponseWriter, r *http.Request) {
	cars := models.GetAllCars()

	json.NewEncoder(w).Encode(cars)

}

func allCarsByBrand(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	models.AllCarsByManifacturer(name)
}

func contains(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	cars := models.Contains(name)

	json.NewEncoder(w).Encode(cars)
}

func addComment(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
	var t tempComment
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Message)

	models.CreateComment(t.Name, t.Message)

	comment := models.LastAddedComment()

	json.NewEncoder(rw).Encode(comment)

}

func allComments(w http.ResponseWriter, r *http.Request) {

	comments := models.AllComments()
	json.NewEncoder(w).Encode(comments)
}
