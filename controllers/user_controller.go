package controllers

import (
	"github.com/Ismadrade/api-book/database"
	"github.com/Ismadrade/api-book/models"
	"github.com/Ismadrade/api-book/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	p.Password = services.SHA256Encoder(p.Password)

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
