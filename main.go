package main

import (
	"context"
	"log"
	"project/compiled"
	"project/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := database.DB.ConnectMySQL()
	if err != nil {
		log.Fatalln("Database connection error:", err.Error())
	}

	tenantId := "tenant-1"
	ctx := context.WithValue(context.Background(), database.TenantContextKey{}, tenantId)
	dbProcess(ctx, tenantId)

	tenantId = "tenant-2"
	ctx = context.WithValue(context.Background(), database.TenantContextKey{}, tenantId)
	dbProcess(ctx, tenantId)
}

func dbProcess(ctx context.Context, tenantId string) {
	database.CreateMigrationsTable(database.DB.Conn, tenantId)
	database.RunMigrations(database.DB.Conn, tenantId)

	err := database.DB.Query.InsertUser(ctx, compiled.InsertUserParams{
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
