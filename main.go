package main

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.POST("/", func(c *gin.Context) {
		secretKey := c.PostForm("secretkey")
		ownemail := c.PostForm("ownemail")
		email := c.PostForm("email")
		message := c.PostForm("message")
		auth := smtp.PlainAuth("", ownemail, secretKey, "smtp.gmail.com")

		to := []string{email}
		msg := message
		// send message on email
		err := smtp.SendMail("smtp.gmail.com:587", auth, ownemail, to, []byte(fmt.Sprint(msg)))

		if err != nil {
			c.String(http.StatusBadRequest, "Something went wrong")
		}
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(200, "about.html", nil)
	})
	//run our server on localhost 8080
	r.Run(":8080")
}
