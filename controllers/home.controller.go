package controllers

import (
	"net/http"

	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/models"
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

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	// Query the database to get all posts
	rows, err := db.Query("SELECT id, title, body, userId FROM posts")
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error querying the database", nil)
		return
	}
	defer rows.Close()

	// Prepare a slice to store the retrieved posts
	var posts []models.Post

	// Loop through the result set and build the posts slice
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.UserID); err != nil {
			utils.JsonResponse(w, http.StatusInternalServerError, "Error scanning database result", nil)
			return
		}
		posts = append(posts, post)
	}

	// Check for any errors encountered while iterating through the result set
	if err := rows.Err(); err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error iterating through database result", nil)
		return
	}

	// Return the posts slice as a JSON response
	utils.JsonResponse(w, http.StatusOK, "Get all posts", posts)
}

func CreateSamplePost(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()

	// Insert a sample post
	_, err := db.Exec("INSERT INTO posts (title, body, userId) VALUES (?, ?, ?)", "Sample post title", "Sample post body", 1)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error inserting into the database", nil)
		return
	}

	// Return a success message
	utils.JsonResponse(w, http.StatusOK, "Post created successfully", nil)
}
