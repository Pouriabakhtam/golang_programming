package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi.ca/modules"
	"restapi.ca/utils"
)

func signup(contex *gin.Context) {
	var user modules.Users
	err := contex.ShouldBindJSON(&user)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"Message": "Couldn't bind the user"})
		return
	}
	err = user.Save()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"Message": "Couldn't store user in DB"})
		return
	}
	contex.JSON(http.StatusOK, gin.H{"Message": "User has been created !!!!!DE😎 "})
}

func Login(contex *gin.Context) {
	var user modules.Users
	err := contex.ShouldBindJSON(&user)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"Message": "Binding has been failed"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		contex.JSON(http.StatusUnauthorized, gin.H{"Message": "Validation is failed"})
		return
	}

	tokengenerated, err := utils.GenerateToken(user.Email, int64(user.Id))
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"Message": "Generating Token was failed"})
		return
	}
	contex.JSON(http.StatusOK, gin.H{"Message": "Authentication has been successful", "Token": tokengenerated})
}
