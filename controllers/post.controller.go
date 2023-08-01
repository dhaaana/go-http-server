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

	rows, err := db.Query(`
		SELECT posts.id, posts.title, posts.body, posts.userId, users.name, users.email
		FROM posts
		JOIN users ON posts.userId = users.id
	`)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer rows.Close()

	var posts []models.PostWithAuthor

	for rows.Next() {
		var post models.PostWithAuthor
		var user models.User

		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &user.ID, &user.Name, &user.Email); err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, "Error scanning database result", nil)
			return
		}
		post.Author = user
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error iterating through database result", nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Get all posts", posts)
}

func GetPostsByUserID(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	userID, err := utils.GetUserIDFromToken(r)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Unauthorized: "+err.Error(), nil)
		return
	}

	rows, err := db.Query(`
		SELECT posts.id, posts.title, posts.body, users.id, users.name, users.email
		FROM posts
		JOIN users ON posts.userId = users.id
		WHERE posts.userId = $1
	`, userID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error fetching posts", nil)
		return
	}

	defer rows.Close()

	var posts []models.PostWithAuthor

	for rows.Next() {
		var post models.PostWithAuthor
		var user models.User
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &user.ID, &user.Name, &user.Email); err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, "Error scanning posts", nil)
			return
		}
		posts = append(posts, post)
	}

	utils.JsonResponse(w, http.StatusOK, "Posts retrieved successfully", posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	userID, err := utils.GetUserIDFromToken(r)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Unauthorized: "+err.Error(), nil)
		return
	}

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid request payload: "+err.Error(), nil)
		return
	}

	if post.Title == "" || post.Body == "" || userID == 0 {
		utils.JsonResponse(w, http.StatusBadRequest, "title, body, and userId must not be empty", nil)
		return
	}

	// Insert the post into the database
	insertQuery := "INSERT INTO posts (title, body, userId) VALUES ($1, $2, $3) RETURNING id;"
	var postID int
	err = db.QueryRow(insertQuery, post.Title, post.Body, userID).Scan(&postID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error creating post", nil)
		return
	}

	// Fetch the newly inserted post from the database
	selectQuery := "SELECT id, title, body, userId FROM posts WHERE id = $1;"
	err = db.QueryRow(selectQuery, postID).Scan(&post.ID, &post.Title, &post.Body, &post.UserID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error fetching newly created post", nil)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, "Post created successfully", post)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	postID := r.URL.Path[len("/posts/"):]

	var post models.PostWithAuthor
	var user models.User

	err := db.QueryRow(`
		SELECT posts.id, posts.title, posts.body, users.id, users.name, users.email
		FROM posts
		JOIN users ON posts.userId = users.id
		WHERE posts.id = $1
	`, postID).Scan(&post.ID, &post.Title, &post.Body, &user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.JsonResponse(w, http.StatusNotFound, "Post not found", nil)
		} else {
			utils.JsonResponse(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	post.Author = user

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

	// Update the post in the database
	updateQuery := "UPDATE posts SET title = $1, body = $2, userId = $3 WHERE id = $4;"
	_, err := db.Exec(updateQuery, post.Title, post.Body, post.UserID, postID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error updating post", nil)
		return
	}

	// Fetch the updated post from the database
	selectQuery := "SELECT id, title, body, userId FROM posts WHERE id = $1;"
	err = db.QueryRow(selectQuery, postID).Scan(&post.ID, &post.Title, &post.Body, &post.UserID)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error fetching updated post", nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Post updated successfully", post)
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
