package main

import (
	"fmt"
	"sqlgo/db"
	"sqlgo/models"
)

func main() {
	db.Connect()
	db.Health()
	//fmt.Println(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	user := models.CreateUser("Kevin", "Kevin123", "kevin@gmail.com")
	fmt.Println(user)
	//db.TruncateTable("users")
	//users := models.ListUsers()
	//fmt.Println(users)
	//db.TruncateTable("users")
	//fmt.Println(models.ListUsers())

	db.Disconnect()
}
