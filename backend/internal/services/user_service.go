package services

import (
	"maps"
	"net/http"
	"strconv"
	"voiting-system/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Users = make(map[string]models.User)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Users)
}

func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
        //TODO::throw exception 400 as incoming request body was not able to map
		return
	}

	for _, value := range Users {
		if value.Email == newUser.Email {
			c.IndentedJSON(http.StatusAlreadyReported, value)
			return
		}
	}

	userId := getRandUserId()
	newUser.UserId = userId
    
    //TODO::VSD Accept password as well
    newUser.SetPassword()
	Users[newUser.Email] = newUser
	c.IndentedJSON(http.StatusCreated, maps.Values(Users))
}

func getRandUserId() string {
	userId := uuid.New().ID()
	return "ACC" + strconv.FormatUint(uint64(userId), 10)
}
