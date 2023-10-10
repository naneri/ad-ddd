package controllers

import (
	"ad-ddd/app/http/dto"
	"ad-ddd/internal/core/users"
	"encoding/json"
	"net/http"
)

type UserController struct {
	UserRepo users.UserRepo
}

func (controller UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

}
