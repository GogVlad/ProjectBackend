package CVs

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

type CV struct {
	id                   int    `json:"id"`
	name                 string `json:"name"`
	linkedinLink         string `json:"linkedinLink"`
	gitLink              string `json:"gitLink"`
	studies              string `json:"studies"`
	experience           string `json:"experience"`
	personalCompetencies string `json:"personal_competencies"`
	address              string `json:"address"`
}

func dbConnect() *sql.DB {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (user *CV) GetId() int {
	return user.id
}

func (user *CV) GetName() string {
	return user.name
}

func (user *CV) GetLinkedinLink() string {
	return user.linkedinLink
}

func (user *CV) GetGitLink() string {
	return user.gitLink
}

func (user *CV) GetStudies() string {
	return user.studies
}

func (user *CV) GetExperience() string {
	return user.experience
}

func (user *CV) GetPersonalCompetencies() string {
	return user.personalCompetencies
}

func (user *CV) GetAddress() string {
	return user.address
}

func (cv *CV) GetCVByID(id int) {
	db := dbConnect()
	defer db.Close()
	sqlStatement := `SELECT * FROM cv WHERE id=? `
	row := db.QueryRow(sqlStatement, id)
	err1 := row.Scan(&cv.id, &cv.name, &cv.linkedinLink, &cv.gitLink, &cv.studies, &cv.experience, &cv.personalCompetencies, &cv.address)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err1)
		}
	}
}

