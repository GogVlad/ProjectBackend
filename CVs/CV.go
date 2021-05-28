package CVs

import "database/sql"

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "projectcv"
	DB_USER = "root"
	DB_PASS = ""
)

type CV struct {
	id                    int    `json:"id"`
	name                  string `json:"name"`
	linkedinLink          string `json:"linkedinLink"`
	gitLink               string `json:"gitLink"`
	studies               string `json:"studies"`
	experience            string `json:"experience"`
	personalCompetencies string `json:"personal_competencies"`
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