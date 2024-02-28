package userservice

import (
	"go-fiber/datamodel"
	userrepository "go-fiber/repository/user"
)

func RegiterUser(user *datamodel.User) error {
	err := userrepository.RegiterUser(user)
	if err != nil {
		return err

	}
	return nil
}
