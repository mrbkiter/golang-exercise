package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", `host=localhost port=5432 sslmode=disable 
							dbname=testgo user=dbapplication_user password=dbapplication_user`)
	if err != nil {
		panic(err)
	}
	//db.Begin()
	candidate := Candidate{Name: "c1"}
	fmt.Println(db.HasTable(&candidate))
	fmt.Println(db.HasTable("candidate"))

	db.Model(&candidate).Save(&candidate)
	//db.Commit()
	defer db.Close()
}

type Candidate struct {
	gorm.Model
	ID    int32
	Name  string
	Email string
}

func (Candidate) TableName() string {
	return "candidate"
}
