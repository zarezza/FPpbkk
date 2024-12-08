package main

import (
	"final-project/config"
	bookcontroller "final-project/controllers/BookController"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := config.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = config.MigrateAndSeed(db)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", bookcontroller.Index)
	http.HandleFunc("/book", bookcontroller.Index)
	http.HandleFunc("/book/index", bookcontroller.Index)
	http.HandleFunc("/book/add", bookcontroller.Add)
	http.HandleFunc("/book/edit", bookcontroller.Edit)
	http.HandleFunc("/book/delete", bookcontroller.Delete)

	fmt.Println("server started at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
