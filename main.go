package main

import (
	"net/http"

	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/config"
	"github.com/dhaaana/go-http-server/routes"
	"github.com/dhaaana/go-http-server/utils"
)

func main() {
	serverPort := config.GetEnvVariables("PORT")

	app.InitDB()
	r := app.NewRouter()

	routes.HomeRoutes(r)

	http.Handle("/", r)

	utils.LogInfo("Server started on " + serverPort)
	http.ListenAndServe(":"+serverPort, nil)
}
