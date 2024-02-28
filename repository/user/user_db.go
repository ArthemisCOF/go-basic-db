package userrepository

import (
	"go-fiber/datamodel"
	"go-fiber/repository"
)

func RegiterUser(user *datamodel.User) error {
	err := repository.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
