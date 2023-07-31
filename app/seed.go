package app

import (
	"log"

	"github.com/dhaaana/go-http-server/models"
)

func SeedData() error {
	samplePosts := []models.Post{
		{Title: "Post 1", Body: "Body of Post 1", UserID: 1},
		{Title: "Post 2", Body: "Body of Post 2", UserID: 1},
		// Add more sample posts as needed
	}

	for _, post := range samplePosts {
		_, err := db.Exec("INSERT INTO posts (title, body, userId) VALUES ($1, $2, $3)", post.Title, post.Body, post.UserID)
		if err != nil {
			return err
		}
	}

	log.Println("Database seeding completed successfully")

	return nil
}
