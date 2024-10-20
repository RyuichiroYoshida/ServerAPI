package controllers

import (
	"github.com/gin-gonic/gin"
	"golangAPI/models"
	"net/http"
)

type RegisterInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		Username: input.UserName,
		Password: input.Password,
	}

	user, err := user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user.PrepareOutput(),
	})
}
