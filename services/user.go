package services

import (
	"errors"
	"fmt"

	"github.com/kalleriakronos24/mygoapp2nd/dto"
	"github.com/kalleriakronos24/mygoapp2nd/models"
)

func (m *module) RetrieveUser(username string) (user models.User, err error) {
	if user, err = m.db.userOrmer.GetOneByUsername(username); err != nil {
		return models.User{}, fmt.Errorf("cannot find user with username %s", username)
	}
	return
}

func (m *module) UpdateUser(id uint, user dto.UserUpdate) (err error) {
	if err = m.db.userOrmer.UpdateUser(models.User{
		ID:    id,
		Email: user.Email,
		Bio:   user.Bio,
	}); err != nil {
		return errors.New("cannot update user")
	}
	return
}
