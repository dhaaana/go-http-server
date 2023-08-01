package routes

import (
	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/controllers"
)

func PostRoutes(r *app.Router) {
	r.Get("/posts", controllers.GetAllPosts)
	r.Get("/posts/user", controllers.GetPostsByUserID)
	r.Post("/posts", controllers.CreatePost)
	r.Get("/posts/:id", controllers.GetPostByID)
	r.Put("/posts/:id", controllers.UpdatePost)
	r.Delete("/posts/:id", controllers.DeletePost)
}
