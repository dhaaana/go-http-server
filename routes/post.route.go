package routes

import (
	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/controllers"
)

func PostRoutes(r *app.Router) {
	r.Get("/posts", controllers.GetAllPosts)
	r.Get("/posts/sample", controllers.CreateSamplePost)
}
