package main

import (
	"net/http"

	"github.com/dhaaana/go-http-server/config"
	"github.com/dhaaana/go-http-server/router"
	"github.com/dhaaana/go-http-server/utils"
)

func main() {
	serverPort := config.GetEnvVariables("PORT")

	http.Handle("/", router.Handler)

	utils.LogInfo("Server started on " + serverPort)
	http.ListenAndServe(":"+serverPort, nil)
}
