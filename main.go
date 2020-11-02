package main

import (
	"fmt"
	"go_certainty_factor/router"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server dijalankan pada port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
