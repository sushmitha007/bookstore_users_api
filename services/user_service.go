package services

import (
	"strings"

	"github.com/sushmitha007/bookstore_users_api/domain/users"
	"github.com/sushmitha007/bookstore_users_api/util/errors"
)

func CreateUser(users users.User) (*users.User, *errors.RestErr) {
	users.Email = strings.TrimSpace(strings.ToLower(users.Email))
	if users.Email == "" {
		return nil, errors.NewBadRequest("Invalid email Id")
	}
	return &users, nil
}
