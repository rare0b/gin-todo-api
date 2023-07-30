package main

import (
	"rare0b/gin-todo-api/db"
	"rare0b/gin-todo-api/router"
)

func main() {
	db := db.Init()

	router.Router(db)
}
