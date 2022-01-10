package main

import (
	"awesomeProject/src/controllers"
	"awesomeProject/src/database"
)

func main() {
	database.Setup()
	app := controllers.InitRoutes()

	app.Run(":8050")

}
