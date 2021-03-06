package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vsfastrack/TaskManager/common"
	"github.com/vsfastrack/TaskManager/data"
	"github.com/vsfastrack/TaskManager/models"
)

//Register users
func Register(w http.ResponseWriter, r http.Request) {
	var dataResource UserResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", 500)
		return
	}

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")

	//Repo for user
	repo := &data.UserRepository{c}

	user := &dataResource.Data

	repo.CreateUser(user)

	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

//Login user
func Login(w http.ResponseWriter, r http.Request) {
	var dataResource LoginResource
	var token string
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login info",
			500,
		)
		return
	}

	loginModel := &dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")

	repo := &data.UserRepository{c}

	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(w, err, "Invalid login credentials", 401)
		return
	} else {
		token, err = common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(w, err, "Eror while generating the access token", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User:  user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
