package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/models"
	"github.com/dhaaana/go-http-server/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if user.Password == "" || user.Email == "" || user.Name == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "name, password, and email must not be empty", nil)
		return
	}

	sqlStmt := `
		INSERT INTO users (password, email, name) 
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	var userID int
	err := db.QueryRow(sqlStmt, user.Password, user.Email, user.Name).Scan(&userID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error registering user: "+err.Error(), nil)
		return
	}

	user.ID = userID

	// Omit the password field in the JSON response
	user.Password = ""

	utils.JsonResponse(w, http.StatusCreated, "User registered successfully", user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if user.Password == "" || user.Email == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Password and email must not be empty", nil)
		return
	}

	sqlStmt := "SELECT COUNT(*) FROM users WHERE email = $1 AND password = $2"
	var count int
	err := db.QueryRow(sqlStmt, user.Email, user.Password).Scan(&count)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error logging in", nil)
		return
	}

	if count == 0 {
		utils.JsonResponse(w, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	sqlStmt = "SELECT id, email, name FROM users WHERE email = $1 AND password = $2"
	err = db.QueryRow(sqlStmt, user.Email, user.Password).Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error logging in", nil)
		return
	}

	token, err := utils.GenerateJWTToken(user.ID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error generating token", nil)
		return
	}

	// Omit the password field in the JSON response
	user.Password = ""

	utils.JsonResponse(w, http.StatusOK, "Login successful", map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	userID, err := utils.GetUserIDFromToken(r)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Unauthorized: "+err.Error(), nil)
		return
	}

	if userID == 0 {
		utils.JsonResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	sqlStmt := "SELECT id, email, name FROM users WHERE id = $1"
	var user models.User
	err = db.QueryRow(sqlStmt, userID).Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.JsonResponse(w, http.StatusNotFound, "User not found", nil)
			return
		}
		utils.JsonResponse(w, http.StatusInternalServerError, "Error retrieving user data", nil)
	}

	// Omit the password field in the JSON response
	user.Password = ""

	utils.JsonResponse(w, http.StatusOK, "User profile retrieved", user)
}
