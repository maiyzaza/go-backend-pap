package main

import (
	"PattayaAvenueProperty/app"
	"PattayaAvenueProperty/config"
	"PattayaAvenueProperty/db"

	"PattayaAvenueProperty/web"
	"fmt"
)

func init() {
	config.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	app := app.New()
	web.StartServer(app)
	fmt.Println("App is running main")
}
