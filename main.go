package main

// r.GET("/login", ControllerLogin)
// fungsi Controller Login seperti apa yang harus dibuat agar aplikasi dapat berjalan

// Catatan :

// go version go1.15.2
// github.com/gin-gonic/gin v1.6.3

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/refinerydev/gogin-login/auth"
	"github.com/refinerydev/gogin-login/helper"
	"github.com/refinerydev/gogin-login/user"
)

func main() {
	fmt.Println("running")
	r := gin.Default()

	// user1 := userService.GetUser("test1@email.com")

	r.POST("/login", func(c *gin.Context) {
		var req helper.LoginRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		userService := user.UserService(user.User{})
		authService := auth.AuthService(userService)

		authenticatedUser, err := authService.UserAuth(req)
		authToken, _ := authService.GenerateToken(req.Email)

		userData := helper.UserResponseFormatter(authenticatedUser, authToken)

		responseFormatter := helper.ResponseFormatter(
			http.StatusOK,
			"success",
			"authenticated",
			userData,
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, responseFormatter)
		}
	})

	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
