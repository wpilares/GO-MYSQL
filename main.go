package main

import (
	"sqlgo/db"
)

func main() {
	db.Connect()
	db.Health()

	//db.CreateTable(models.UserSchema, "users")

	//Restore Table
	//db.TruncateTable("users")

	//user := models.CreateUser("Kevin", "Kevin123", "kevin@gmail.com")
	//fmt.Println(user)

	//users := models.ListUsers()
	//fmt.Println(users)

	//user := models.GetUser(1)
	//fmt.Println(user)

	//Update User
	//user := models.GetUser(1)
	//fmt.Println(user)
	//user.Username="Ruben"
	//user.Save()

	//Delete User
	//user := models.GetUser(1)
	//fmt.Println(user)
	//user.Delete()

	db.Disconnect()
}
