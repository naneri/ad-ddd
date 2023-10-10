package http

import (
	"ad-ddd/app/http/controllers"
	"net/http"
)

func main() {

	uc := controllers.UserController{}

	http.HandleFunc("/users/create", uc.Create)

	panic(http.ListenAndServe(":3000", nil))
}
