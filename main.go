package main

import (
	"awesomeProject/src/models"
	"fmt"
)

func main() {
	var user models.User
	user.Id = "5"
	fmt.Println(user.Id)
}
