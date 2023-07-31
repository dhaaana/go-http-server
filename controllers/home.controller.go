package controllers

import (
	"net/http"

	"github.com/dhaaana/go-http-server/utils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, http.StatusOK, "Hello World", nil)
}
