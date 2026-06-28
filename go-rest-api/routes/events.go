package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"restapi.ca/modules"
)

func GetEvents(contex *gin.Context) {
	events, err := modules.GetAllEvents()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"Message": "Counldn't fetch events"})
	}
	contex.JSON(http.StatusOK, events)
}
func GetSingleEvent(contex *gin.Context) {
	eventID, err := strconv.ParseInt(contex.Param("id"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"Message": "Id value is incorrect"})
		return
	}
	event, err := modules.GetEventbyID(eventID)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"Message": "Couldn't fetch the sinle event"})
	}
	contex.JSON(http.StatusOK, event)

}

func CreateEvent(contex *gin.Context) {

	var events modules.Events
	err := contex.ShouldBindJSON(&events)
	if err != nil {
		log.Println("bind error:", err)
		contex.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userid := contex.GetInt64("user_id")
	events.ID = userid
	events.UserID = 1243
	contex.JSON(http.StatusCreated, gin.H{"Message": "Event has been created", "event": events})
	err = events.Save()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"Message": "Counldn't create events"})
	}
}
func UpdateEvent(contex *gin.Context) {
	eventID, err := strconv.ParseInt(contex.Param("id"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"Message": "Couldn't get that id"})
	}
	event, err := modules.GetEventbyID(eventID)
	user_id := contex.GetInt64("userId")
	if event.UserID != user_id {
		contex.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized user!!"})
		return
	}
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"Message": "Couldn't get that id"})
	}

	var updatedEvent modules.Events
	err = contex.ShouldBindJSON(&updatedEvent)
	if err != nil {
		log.Println("bind error:", err)
		contex.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.UpdateEvent()
	if err != nil {
		log.Fatal(err)
		contex.JSON(http.StatusInternalServerError, gin.H{"Message": "Update was failed"})
	}
	contex.JSON(http.StatusOK, gin.H{"Message": "Update successful"})

}

func DeleteEvent(contex *gin.Context) {
	eventID, err := strconv.ParseInt(contex.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
		contex.JSON(http.StatusBadRequest, gin.H{"Message": "Couldn't fetch the event ID"})
	}
	event, err := modules.GetEventbyID(eventID)
	user_id := contex.GetInt64("userId")
	if user_id != event.UserID {
		contex.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized user!!"})
		return
	}
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"Message": "Couldn't get the event ID"})
	}
	err = event.DeleteEvent()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"Message": "Delettion was not successful"})
	}
	contex.JSON(http.StatusOK, gin.H{"Message": "Event has been deleted!"})
}
