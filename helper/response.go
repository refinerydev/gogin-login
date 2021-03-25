package helper

import "github.com/refinerydev/gogin-login/user"

type ResponseUser struct {
	Email     string `json:"email"`
	AuthToken string `json:"auth_token"`
}

func UserResponseFormatter(user user.User, auth_token string) ResponseUser {
	formatter := ResponseUser{
		Email:     user.Email,
		AuthToken: auth_token,
	}

	return formatter
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func ResponseFormatter(code int, status string, message interface{}, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}
