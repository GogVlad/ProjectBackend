package SuperCars

import (
	"database/sql"
	"fmt"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "super_cars"
	DB_USER = "root"
	DB_PASS = ""
)

type iCar interface {
	AddCar(id int, brand string, model string, class string, productionCountry string,
		horsepower int, nrOfDoors int, lastKnownPrice string)
	GetCarByName(brand string, model string) Car
	GetAllCars() []Car
	updateCar(id int, brand string, model string, class string, productionCountry string,
		horsepower int, nrOfDoors int, lastKnownPrice string)
	deleteCar(id int)
}

type Car struct {
	id                int `json:"id"`
	brand             string `json:"brand"`
	model             string `json:"model"`
	class             string `json:"class"`
	productionCountry string `json:"production_country"`
	horsepower        int `json:"horsepower"`
	nrOfDoors         int `json:"nr_of_doors"`
	lastKnownPrice    string `json:"last_known_price"`
}

func dbConnect () *sql.DB{
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func AddCar(id int, brand string, model string, class string, productionCountry string,
	horsepower int, nrOfDoors int, lastKnownPrice string){

	db:=dbConnect()
	defer db.Close()
	newCar := fmt.Sprintf("INSERT INTO cars VALUES (%d, '%s', '%s', '%s','%s', %d, %d, '%s')", id, brand, model, class,
		productionCountry, horsepower, nrOfDoors, lastKnownPrice)

	insert, err := db.Query(newCar)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func (searchedCar *Car)GetCarByName(brand string, model string) Car {
	db:=dbConnect()
	defer db.Close()
	sqlStatement := `SELECT * FROM cars WHERE brand=? AND model=? `
	row := db.QueryRow(sqlStatement, brand, model)
	err1 := row.Scan(&searchedCar.id, &searchedCar.brand, &searchedCar.model, &searchedCar.class, &searchedCar.productionCountry, &searchedCar.horsepower, &searchedCar.nrOfDoors,
		&searchedCar.lastKnownPrice)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err1)
		}
	}
	return *searchedCar
}

func (car *Car) GetAllCars() []Car{
	db:=dbConnect()
	defer db.Close()
	var cars []Car
	rows, err1 := db.Query("SELECT * FROM cars")
	defer rows.Close()
	for rows.Next() {
		err1 = rows.Scan(&car.id, &car.brand, &car.model, &car.class, &car.productionCountry, &car.horsepower, &car.nrOfDoors,
			&car.lastKnownPrice)
		if err1 != nil {
			if err1 == sql.ErrNoRows {
				fmt.Println("Zero rows found")
			} else {
				panic(err1)
			}
		}
		cars=append(cars, *car)
	}
	return cars
}

func  DeleteCar(id int){
	db:=dbConnect()
	defer db.Close()
	del, err1 := db.Prepare("DELETE FROM Cars WHERE id=?")
	if err1 != nil {
		panic(err1.Error())
	}
	del.Exec(id)
	defer db.Close()
}

func UpdateCar(id int, brand string, model string, class string, productionCountry string,
	horsepower int, nrOfDoors int, lastKnownPrice string){

	db:=dbConnect()
	defer db.Close()
	ins, err1 := db.Prepare("UPDATE cars SET brand=?, model=?, class=?, production_country=?, horsepower=?, number_of_doors=?, last_known_price=? WHERE id=?")
	if err1 != nil {
		panic(err1.Error())
	}
	ins.Exec(brand, model, class, productionCountry, horsepower, nrOfDoors, lastKnownPrice, id)
}

func (car *Car) GetId() int{
	return car.id
}
func (car *Car) GetBrand() string{
	return car.brand
}
func (car *Car) GetModel() string{
	return car.model
}
func (car *Car) GetClass() string{
	return car.class
}
func (car *Car) GetProductionCountry() string{
	return car.productionCountry
}
func (car *Car) GetHP() int{
	return car.horsepower
}
func (car *Car) GetNrOfDoors() int{
	return car.nrOfDoors
}
func (car *Car) GetPrice() string{
	return car.lastKnownPrice
}

