package routes

import (
	"github.com/Thanasak1412/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	userId, exists := c.Get("foo")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "user does not exist"})

		return
	}

	if userIdFloat, ok := userId.(float64); ok {
		event.UserId = int64(userIdFloat)
	}

	err := event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	var updatedEvent models.Event

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "details": err.Error()})

		return
	}

	if err = c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "details": err.Error()})

		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	updatedEvent.Id = eventId

	err = updatedEvent.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	c.JSON(http.StatusOK, updatedEvent.Id)
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "details": err.Error()})

		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not found", "details": err.Error()})

		return
	}

	if err = models.DeleteEventById(eventId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event", "details": err.Error()})

		return
	}

	c.JSON(http.StatusNoContent, nil)
}
