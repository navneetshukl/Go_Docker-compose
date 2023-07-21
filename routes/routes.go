package routes

import (
	"docker-compose/database"
	"docker-compose/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

var users []models.User

func Home(c *gin.Context) {

	users=database.FindAllUser()

	c.HTML(200,"home.page.tmpl",gin.H{
		"Users":users,
	})

}

func Display(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")

	fmt.Println(name, " ", email, " ", phone)

	database.InsertUser(name,email,phone)
	

	users=database.FindAllUser()

	c.HTML(200,"home.page.tmpl",gin.H{
		"Users":users,
	})
}
