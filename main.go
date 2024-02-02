

package main

import (
	"toDoList/db"
	"toDoList/model"
	"toDoList/router"
)

func main() {
	// create database
	// sql: CREATE DATABASE todo;
	// link db
	if err := db.InitMySQL(); err != nil {
		panic(err)
	}
	defer db.Close()
	// create table
	model.InitModel()
	// create router
	r := router.SetupRouter()
	// run
	r.Run(":3000")
}
