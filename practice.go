package main

import (
	"encoding/json"
	"time"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type User struct {
// 	Name string `json:"name"`
// }

// type MyType struct {
// 	A string
// 	FooBar string
// }

func GetDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "ucchy"
	PASS := ",8M6)L(6cZ/YwDQa9"
	DBNAME := "go_app"

	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME
	return DBMS, CONNECT 
}

type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Person struct {
	Model
	Name string
	Age int
}

var person Person
func main() {
	//practice json
	// user := User{"test"}
	// b, _ := json.Marshal(user)
	// fmt.Printf("%s\n", string(b))
	// var mt MyType
	// json.Unmarshal([]byte(`{"A":"aaa", "FooBar":"baz"}`), &mt)
	// fmt.Printf("%#v\n", mt)

	db := GetDBConn()
	// db.AutoMigrate(&Person{})
	// var person = Person{Name: "Gopher", Age: 10}
	// db.Create(&person)
	db.First(&person)
	responseBody, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(responseBody)
}