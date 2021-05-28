package main

import (
	"backend/backend.go/CVs"
	"backend/backend.go/Users"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.HandleFunc("/addUser", AddUser)
	mux.HandleFunc("/getUser", GetUser)
	mux.HandleFunc("/getCV",GetCV)
	//mux.HandleFunc("/GetAllCars",GetAllCars)
	//mux.HandleFunc("/UpdateCar",UpdateCar)
	//mux.HandleFunc("/DeleteCar",DeleteCar)

	fmt.Println("Starting server...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func GetCV(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getCV" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id1, _ := strconv.Atoi(id)
	switch r.Method {
	case "GET":
		{
			var cv CVs.CV
			cv.GetCVByID(id1)

			fmt.Fprintf(w, "User info: \n")
			fmt.Fprintf(w, "Name: %s \n", cv.GetName())
			fmt.Fprintf(w, "LinkedIn link: %s \n", cv.GetLinkedinLink())
			fmt.Fprintf(w, "GitHub link: %s \n", cv.GetGitLink())
			fmt.Fprintf(w, "Studies: %s \n", cv.GetStudies())
			fmt.Fprintf(w, "Experience: %s \n", cv.GetExperience())
			fmt.Fprintf(w, "Personal competencies: %s \n", cv.GetPersonalCompetencies())
		}
	default:
		fmt.Fprintf(w, "Expected method GET")
	}

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/addUser" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Missing username parameter", http.StatusBadRequest)
		return
	}
	password := r.URL.Query().Get("password")
	if password == "" {
		http.Error(w, "Missing password parameter", http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}
	birth_date := r.URL.Query().Get("birth_date")
	if birth_date == "" {
		http.Error(w, "Missing birth_date parameter", http.StatusBadRequest)
		return
	}
	details := r.URL.Query().Get("details")
	if details == "" {
		http.Error(w, "Missing details parameter", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		{
			var user Users.User
			user.GetAllUsers()
			Users.AddUser(user.GetId()+1, username, password, name, birth_date, details)

			fmt.Fprintf(w, "New user added! \n")

		}
	default:
		fmt.Fprintf(w, "Expected method POST")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getUser" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Missing username parameter", http.StatusBadRequest)
		return
	}
	password := r.URL.Query().Get("password")
	if password == "" {
		http.Error(w, "Missing password parameter", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		{
			var user Users.User
			user.GetUser(username, password)

			fmt.Fprintf(w, "User info: \n")
			fmt.Fprintf(w, "Name: %s \n", user.GetName())
			fmt.Fprintf(w, "Birth Date: %s \n", user.GetBirthDate())
			fmt.Fprintf(w, "Details: %s \n", user.GetDetails())
		}
	default:
		fmt.Fprintf(w, "Expected method GET")
	}
}

/*
func enableCors(w *http.ResponseWriter) {(*w).Header().Set("Access-Control-Allow-Origin", "*")}

func GetAllCars(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	if r.URL.Path != "/GetAllCars" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":{
		var car SuperCars.Car
		cars := car.GetAllCars()

		fmt.Fprintf(w, "All the cars in the database are: \n")
		for _, existingCar := range cars {
			fmt.Fprintf(w, "The Brand is %s\n", existingCar.GetBrand())
			fmt.Fprintf(w, "The model is %s\n", existingCar.GetModel())
			fmt.Fprintf(w, "The Class is %s\n", existingCar.GetClass())
			fmt.Fprintf(w, "The Production Country is %s\n", existingCar.GetProductionCountry())
			fmt.Fprintf(w, "Horsepower = %d\n", existingCar.GetHP())
			fmt.Fprintf(w, "The number of doors is %d\n", existingCar.GetNrOfDoors())
			fmt.Fprintf(w, "The latest price estimation for the car is = %s\n", existingCar.GetPrice())
		}
	}
	default:
		fmt.Fprintf(w, "Expected method GET")
	}
}

func UpdateCar(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/UpdateCar" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	brand := r.URL.Query().Get("brand")
	if brand == "" {
		http.Error(w, "Missing brand parameter", http.StatusBadRequest)
		return
	}
	model := r.URL.Query().Get("model")
	if model == "" {
		http.Error(w, "Missing model parameter", http.StatusBadRequest)
		return
	}
	class := r.URL.Query().Get("class")
	if class == "" {
		http.Error(w, "Missing class parameter", http.StatusBadRequest)
		return
	}
	productionCountry := r.URL.Query().Get("productionCountry")
	if productionCountry == "" {
		http.Error(w, "Missing productionCountry parameter", http.StatusBadRequest)
		return
	}
	horsepower := r.URL.Query().Get("horsepower")
	if horsepower == "" {
		http.Error(w, "Missing horsepower parameter", http.StatusBadRequest)
		return
	}
	nrOfDoors := r.URL.Query().Get("nrOfDoors")
	if nrOfDoors == "" {
		http.Error(w, "Missing nrOfDoors parameter", http.StatusBadRequest)
		return
	}
	lastKnownPrice := r.URL.Query().Get("lastKnownPrice")
	if lastKnownPrice == "" {
		http.Error(w, "Missing lastKnownPrice parameter", http.StatusBadRequest)
		return
	}
	intID, _ := strconv.Atoi(id)
	hp, _ := strconv.Atoi(horsepower)
	nD, _ := strconv.Atoi(nrOfDoors)

	switch r.Method {
	case "PUT":
		{
			SuperCars.UpdateCar(intID, brand, model, class, productionCountry, hp, nD, lastKnownPrice)

			fmt.Fprintf(w, "The last updated car's details are: \n")
			fmt.Fprintf(w, "The Brand is %s\n", brand)
			fmt.Fprintf(w, "The model is %s\n", model)
			fmt.Fprintf(w, "The Class is %s\n", class)
			fmt.Fprintf(w, "The Production Country is %s\n", productionCountry)
			fmt.Fprintf(w, "Horsepower = %d\n", hp)
			fmt.Fprintf(w, "The number of doors is %d\n", nD)
			fmt.Fprintf(w, "The latest price estimation for the car is = %s\n", lastKnownPrice)
		}
	}
}

func DeleteCar(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/DeleteCar" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	intID, _ := strconv.Atoi(id)
	switch r.Method {
	case "DELETE":
		{
			SuperCars.DeleteCar(intID)

			fmt.Fprintf(w, "The car was deleted! \n")
		}
	}
}
*/
