package Users

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	encryptedPass := EncriptDetails(password)
	newCar := fmt.Sprintf("INSERT INTO user VALUES (%d, '%s', '%s', '%s','%s','%s')", id, username, encryptedPass, name,
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

func EncriptDetails(pass string) string{
	symbols := "ABCDEFGH:IJKLMNOP;QRSTUVWX/YZabcdef?ghijklmn!opqrstuv(wxyz0123)456789 .-=+[],*{}@"
	encryption:=encryptPassword(symbols, pass, "encrypted.txt")
	return encryption
}

func encryptPassword(symbols string, input string, outputFile string) string{
	words, _ := os.Create(outputFile)
	w := bufio.NewWriter(words)
	for i := 0; i < len(input)-1; i = i + 2 {
		p1 := strings.Index(symbols, string(input[i]))
		p2 := strings.Index(symbols, string(input[i+1]))
		if p1 == -1 || p2 == -1 {
			_, _ = fmt.Fprintf(w, "%c%c", input[i], input[i+1])
		} else {
			l1 := p1 / 9
			c1 := p1 % 9
			l2 := p2 / 9
			c2 := p2 % 9
			p1 = l2*9 + c1
			p2 = l1*9 + c2
			_, _ = fmt.Fprintf(w, "%c%c", symbols[p1], symbols[p2])
		}
	}
	_ = w.Flush()
	encryptionByte, _ := ioutil.ReadFile(outputFile)
	encryptionString := string(encryptionByte[:])
	fmt.Println(encryptionString)
	return encryptionString
}