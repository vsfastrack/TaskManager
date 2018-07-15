package controllers

import "github.com/vsfastrack/TaskManager/models"

type (
	//UserResource resource for Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	//LoginResource resource for Post -/users/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	//AuthUserResource resource for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	//LoginModel Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//AuthUserModel for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	TaskResource struct {
		Data models.Task `json:"data"`
	}

	TasksResource struct {
		Data []models.Task `json:"data"`
	}
)
