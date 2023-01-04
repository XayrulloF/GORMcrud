package main

import (
	"encoding/json"
	"fmt"
	"gorm/dbtools"
	"gorm/model"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	DataSourceName string `json:"dataSourceName"`
}

func main() {
	file, err := os.Open("configuration/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	conf := new(Configuration)
	json.NewDecoder(file).Decode(conf)
	dbtools.DBInitializer(conf.DataSourceName)

	//Create table ...................
	//dbtools.CreateTable(&model.Students{})
	//................................

	//insert..........................
	// student:=model.Students{
	// 	Id: 3,
	// 	Name: "Xyon",
	// 	Age: 45,
	// }
	// dbtools.Save(student)
	// fmt.Println("Inserted an info")
	//.................................

	//select all.............................
	//var students []model.Students
	//dbtools.Select(students)
	//.......................................

	//single update .........................
	// student := model.Students{
	// 	Id: 1,
	// }
	// data := map[string]interface{}{
	// 	"Name": "Malik",
	// 	"Age":  1222,
	// }
	// row := dbtools.SingleUpdate(&student, data)
	// fmt.Println("Rows affected: ", row)
	//.........................................

	//update rows where name is kim............
	// name := "kim"
	// nameWhereClause := "Name='" + name + "'"
	// data := map[string]interface{}{
	// 	"Name": "David",
	// 	"Age":  50,
	// }
	// rows := dbtools.MultipleUpdate(model.Students{}, nameWhereClause, data)
	// fmt.Println("Rows affected: ", rows)
	//..........................................

	//update rows where age is 50............
	// age := 50
	// ageWhereClause := "Age=" + strconv.Itoa(age)
	// data := map[string]interface{}{
	// 	"Name": "Bory",
	// 	"Age":  1111,
	// }
	// rows := dbtools.MultipleUpdate(model.Students{}, ageWhereClause, data)
	// fmt.Println("Rows affected: ", rows)
	//..........................................

	//Single delete.............................
	// student := model.Students{
	// 	Id: 5,
	// }
	// row := dbtools.SingleDelete(&student)
	// fmt.Println("Rows affected: ", row)
	//..........................................

	//Multi delete by name.............................
	// name := "Chon"
	// nameWhereClause := "Name='" + name + "'"
	// rows := dbtools.MultipleDelete(model.Students{}, nameWhereClause)
	// fmt.Println("Rows affected ", rows)
	//.................................................

	//Multi delete by age..............................
	age := 1
	ageWhereClause := "Age=" + strconv.Itoa(age)
	row := dbtools.MultipleDelete(model.Students{}, ageWhereClause)
	fmt.Println(row)
	//..................................................
}
