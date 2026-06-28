package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi.ca/modules"
	"restapi.ca/utils"
)

func Authorization(contex *gin.Context) { 
	token := contex.Request.Header.Get("Authorization")
	if token == "" {
		contex.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not Authorized"})
		return 
	}

	User_id, err := utils.VerifyingToken(token)
	if err != nil {
		contex.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "User is not authorized"})
		return 
	}
	contex.Set("user_id",User_id)
	contex.Next()
}