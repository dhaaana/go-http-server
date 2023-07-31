package controllers

import (
	"net/http"

	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/models"
	"github.com/dhaaana/go-http-server/utils"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	rows, err := db.Query("SELECT id, title, body, userId FROM posts")
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.UserID); err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, "Error scanning database result", nil)
			return
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error iterating through database result", nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Get all posts", posts)
}

func CreateSamplePost(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	_, err := db.Exec("INSERT INTO posts (title, body, userId) VALUES ($1, $2, $3)", "Sample post title", "Sample post body", 1)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Post created successfully", nil)
}
