package main

// r.GET("/login", ControllerLogin)
// fungsi Controller Login seperti apa yang harus dibuat agar aplikasi dapat berjalan

// Catatan :

// go version go1.15.2
// github.com/gin-gonic/gin v1.6.3

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("running")
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})
	r.Run()
}
