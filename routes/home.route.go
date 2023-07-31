package routes

import (
	"github.com/dhaaana/go-http-server/app"
	"github.com/dhaaana/go-http-server/controllers"
)

func HomeRoutes(r *app.Router) {
	r.Get("/", controllers.HelloWorld)
}
