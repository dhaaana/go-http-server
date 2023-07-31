package controllers

import (
	"database/sql"
	"encoding/json"
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if post.Title == "" || post.Body == "" || post.UserID == 0 {
		utils.JsonResponse(w, http.StatusBadRequest, "title, body, and userId must not be empty", nil)
		return
	}

	_, err := db.Exec("INSERT INTO posts (title, body, userId) VALUES ($1, $2, $3)", post.Title, post.Body, post.UserID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, "Post created successfully", nil)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	postID := r.URL.Path[len("/posts/"):]

	var post models.Post
	err := db.QueryRow("SELECT id, title, body, userId FROM posts WHERE id = $1", postID).Scan(&post.ID, &post.Title, &post.Body, &post.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.JsonResponse(w, http.StatusNotFound, "Post not found", nil)
		} else {
			utils.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Get post by ID", post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	postID := r.URL.Path[len("/posts/"):]

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if post.Title == "" || post.Body == "" || post.UserID == 0 {
		utils.JsonResponse(w, http.StatusBadRequest, "title, body, and userId must not be empty", nil)
		return
	}

	_, err := db.Exec("UPDATE posts SET title = $1, body = $2, userId = $3 WHERE id = $4", post.Title, post.Body, post.UserID, postID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Post updated successfully", nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	postID := r.URL.Path[len("/posts/"):]

	_, err := db.Exec("DELETE FROM posts WHERE id = $1", postID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Post deleted successfully", nil)
}
