package main

import (
	"fmt"
	"log"

	"github.com/mdzkm/todo-list-app/server/internal/database"
	"github.com/mdzkm/todo-list-app/server/internal/router"
)

//	@title			ToDo List App
//	@version		1.0
//	@description	This is a simple todo list app.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://github.com/mdzkm
//	@contact.email	mdzkm17@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/tasks

func main() {
	database.Init()
	r := router.Router()
	port := ":8080"
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(r.Listen(port))
}
