package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushmitha007/bookstore_users_api/domain/users"
	"github.com/sushmitha007/bookstore_users_api/services"
	"github.com/sushmitha007/bookstore_users_api/util/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// the above logic is done in one line using gin
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)

		restErr := errors.NewBadRequest("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
