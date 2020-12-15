package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
var (
	users = map[int64]*User{}
	currentId int64 =1
	router = gin.Default()
)
type httpError struct {
	Message string`json:"message"`
	Code int`json:"code"`
}

type User struct{
	Id int64`json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}
func main(){
	mapUrls()
	router.Run(":8080")
}
func mapUrls(){
	router.POST("/users",Create)
	router.GET("/users/:id", Get)

}
func Get(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err!= nil{
		httpErr := httpError{Message:"invalid user id", Code:http.StatusBadRequest}
		c.JSON(httpErr.Code , httpErr)
		return

	}
	user := users[userId]
	if user == nil {
		httpErr := httpError{Message:fmt.Sprintf("user %d not found", userId), Code:http.StatusNotFound}
		c.JSON(httpErr.Code , httpErr)
		return
	}
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusOK, user)


}



func Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		httpErr := httpError{Message:"invalid json body", Code:http.StatusBadRequest}
			c.JSON(httpErr.Code , httpErr)
		return

	}

	user.Id = currentId
	currentId ++
	users[user.Id] = &user
	c.JSON(http.StatusCreated, user)

}
