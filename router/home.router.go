package router

import (
	"net/http"

	"github.com/dhaaana/go-http-server/controllers"
	"github.com/dhaaana/go-http-server/middleware"
)

var Handler = http.NewServeMux()

func init() {
	Handler.HandleFunc("/", middleware.Logging(controllers.Home))
}
