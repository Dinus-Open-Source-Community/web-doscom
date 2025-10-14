package main

import (
	"web_doscom/internal/database"
	"web_doscom/internal/env"
)

func main() {
	// load env
	env.LoadEnv()

	// connect database
	database.ConnectDB()

}
