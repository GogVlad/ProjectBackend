package Users

import (
	"database/sql"
	"fmt"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "projectcv"
	DB_USER = "root"
	DB_PASS = ""
)

type User struct {
	id         int    `json:"id"`
	username   string `json:"username"`
	password   string `json:"password"`
	name       string `json:"name"`
	birth_date string `json:"birth_date"`
	details    string    `json:"details"`
}

func dbConnect() *sql.DB {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func AddUser(id int, username string, password string, name string, birth_date string,
	details string) {

	db := dbConnect()
	defer db.Close()
	newCar := fmt.Sprintf("INSERT INTO user VALUES (%d, '%s', '%s', '%s','%s','%s')", id, username, password, name,
		birth_date, details)
	insert, err := db.Query(newCar)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func (user *User) GetAllUsers() []User{
	db:=dbConnect()
	defer db.Close()
	var users []User
	rows, err1 := db.Query("SELECT * FROM user")
	defer rows.Close()
	for rows.Next() {
		err1 = rows.Scan(&user.id, &user.username, &user.password, &user.name, &user.birth_date, &user.details)
		if err1 != nil {
			if err1 == sql.ErrNoRows {
				fmt.Println("Zero rows found")
			} else {
				panic(err1)
			}
		}
		users=append(users, *user)
	}
	return users
}

func (user *User) GetId() int{
	return user.id
}
func (user *User) GetUsername() string{
	return user.username
}
func (user *User) GetPassword() string{
	return user.password
}
func (user *User) GetName() string{
	return user.name
}
func (user *User) GetBirthDate() string{
	return user.birth_date
}
func (user *User) GetDetails() string{
	return user.details
}