package fungsi

import (
	"fmt"

	"../connection"
)

type whatsapps struct {
	ID   int    `json:"id"`
	Info string `json:"info"`
	Name string `json:"name"`
	User string `json:"user"`
}

func Select(table string) {

	db, err := connection.Connect()
	if err != nil {
		panic(err.Error())
		// fmt.Println("[ERROR] ", err.Error)
	}
	fmt.Println("success access")
	defer db.Close()

	result, err := db.Query("SELECT id, info, name, user FROM " + table)
	if err != nil {
		// fmt.Println("[ERROR] ", err.Error)
		panic(err.Error())
	}
	fmt.Println("success query")
	defer result.Close()

	for result.Next() {
		var wa whatsapps
		err = result.Scan(&wa.ID, &wa.Info, &wa.Name, &wa.User)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("Data : %v, %v, %s, %s \n", wa.ID, wa.Info, wa.Name, wa.User)

	}
}
