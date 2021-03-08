package main

// r.GET("/login", ControllerLogin)
// fungsi Controller Login seperti apa yang harus dibuat agar aplikasi dapat berjalan

// Catatan :

// go version go1.15.2
// github.com/gin-gonic/gin v1.6.3

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/refinerydev/gogin-login/user"
)

func main() {
	fmt.Println("running")
	r := gin.Default()

	userService := user.UserService(user.User{})
	user1 := userService.GetUser("test1@email.com")

	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": user1})
	})
	r.Run()
}
