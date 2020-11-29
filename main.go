package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"godb/dbutil"
	"godb/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	log.Println("Start")
	conf := dbutil.NewDbConf(
		"postgres",
		"admin",
		"postgres",
		5432,
		"postgres",
	).Mapped("localhost", 5432)
	db, err := gorm.Open(postgres.Open(conf.String()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&model.Person{}); err != nil {
		panic(err)
	}

	p1 := model.Person{FirstName: "John", LastName: "Doe"}
 	p2 := model.Person{FirstName: "Jane", LastName: "Smith"}

 	db.Create(&p1)
 	db.Commit()
 	var p3 model.Person
 	db.Find(&p3)

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	log.Println("End")
}
