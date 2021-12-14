package main

import (
	"fmt"
	"gorm-migration/database"
)

func main() {

	// connect to the database
	database.ConnectDB()

	fmt.Println("gorm migration")
}
