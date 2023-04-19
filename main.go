package main

import (
	"Assignment3/controller"
	"Assignment3/repository"
	"net/http"
)

func main() {
	conf := repository.Config{Host: "localhost", Port: "5432", Username: "postgres", Password: "ramka123", DBName: "bookstore", SSLMode: "disable"}
	db, _ := repository.NewpostgresDB(conf)
	postgres := repository.Postgres{db}
	c := controller.Controller{&postgres}
	r := controller.Routes(c)
	http.ListenAndServe(":8080", r)

}
