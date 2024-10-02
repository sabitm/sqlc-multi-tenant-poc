package main

import (
	"context"
	"log"
	"project/compiled"
	"project/database"

	// _ "modernc.org/sqlite"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.WithValue(context.Background(), database.TenantContextKey{}, "tenant1")

	_, err := database.DB.ConnectMySQL()
	if err != nil {
		log.Fatalln("Database connection error:", err.Error())
	}

	err = database.DB.Query.InsertUser(ctx, compiled.InsertUserParams{
		Email: "tech@samba.com",
		Name:  "Samba",
		Role:  "admin",
	})
	if err != nil {
		log.Fatalln("Error inserting record:", err.Error())
	}

	res, err := database.DB.Query.GetAllUsers(ctx)
	if err != nil {
		log.Fatalln("Error getting record:", err.Error())
	}

	log.Println("Users:", res)
}
