package routes

import (
	"net/http"

	handlers "github.com/likhithkp/WorkMate/api/handlers/auth"
)

func UserRouter(mux *http.ServeMux) {
	mux.HandleFunc("/signup", handlers.SignUp)
}
