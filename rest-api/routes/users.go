package routes

import (
	"github.com/Thanasak1412/go-rest-api/models"
	"github.com/Thanasak1412/go-rest-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func signUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse sign-up body", "details": err.Error()})

		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse sign-up body", "details": err.Error()})

		return
	}

	existingUser, err := models.GetUserByEmail(user.Email)

	if err != nil && !strings.Contains(err.Error(), "no rows") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to sign-up", "details": err.Error()})

		return
	}

	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists", "details": user.Email})

		return
	}

	if err = user.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to sign-up", "details": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func singIn(c *gin.Context) {
	var rq models.User

	if err := c.ShouldBindJSON(&rq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse sign-in body", "details": err.Error()})

		return
	}

	validate := validator.New()
	if err := validate.Struct(rq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse sing-in body", "details": err.Error()})

		return
	}

	existingUser, err := models.GetUserByEmail(rq.Email)

	if err != nil || existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to sing-in", "details": err})

		return
	}

	if err = utils.VerifyPassword(existingUser.Password, rq.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid credentials"})

		return
	}

	token, err := utils.GenerateJWT(existingUser.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to sign-in token"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully sign-in", "data": gin.H{"token": token}})
}
