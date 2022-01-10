package main

import "awesomeProject/src/controllers"

func main() {
	app := controllers.InitRoutes()
	app.Run(":8050")

}
