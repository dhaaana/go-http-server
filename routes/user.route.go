package routes

import (
	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/controllers"
)

func UserRoutes(r *app.Router) {
	r.Post("/login", controllers.Login)
	r.Post("/register", controllers.Register)
	r.Get("/profile", controllers.GetProfile)
}
