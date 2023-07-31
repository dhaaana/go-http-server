package controllers

import (
	"net/http"

	"github.com/dhaaana/go-http-server/utils"
)

func GetHomeData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.JsonResponse(w, http.StatusNotFound, "Not Found", nil)
		return
	}
	utils.JsonResponse(w, http.StatusOK, "Welcome to the home", []string{})
}

func GetAlternateHomeData(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, http.StatusOK, "Welcome to the alternate home", [0]string{})

}

func PostHomeData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	utils.JsonResponse(w, http.StatusOK, "Post success", map[string]string{"name": name})
}
