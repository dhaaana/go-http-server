package main

import (
	"net/http"
	"os"

	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/config"
	"github.com/dhaaana/go-http-server/routes"
	"github.com/dhaaana/go-http-server/utils"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		utils.LogError("Error loading .env file:", err)
		os.Exit(1)
	}

	serverPort, err := config.GetEnvVariables("PORT")
	if err != nil {
		utils.LogError("Error getting server port:", err)
		os.Exit(1)
	}

	app.InitDB()

	r := app.NewRouter()
	routes.HomeRoutes(r)
	routes.PostRoutes(r)
	http.Handle("/", r)

	utils.LogInfo("Server started on " + serverPort)
	http.ListenAndServe(":"+serverPort, nil)
}
