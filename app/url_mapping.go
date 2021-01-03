package app

import (
	"github.com/sushmitha007/bookstore_users_api/controller/ping"
	"github.com/sushmitha007/bookstore_users_api/controller/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/user/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)

}
