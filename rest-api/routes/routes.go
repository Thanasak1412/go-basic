package routes

import (
	"github.com/Thanasak1412/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Events
	server.GET("/events", utils.VerifyJWT, getEvents)
	server.GET("/events/:id", utils.VerifyJWT, getEvent)
	server.POST("/events", utils.VerifyJWT, createEvent)
	server.PUT("/events/:id", utils.VerifyJWT, updateEvent)
	server.DELETE("/events/:id", utils.VerifyJWT, deleteEvent)

	// Users
	server.POST("/sign-up", signUp)
	server.POST("/sign-in", singIn)
}
