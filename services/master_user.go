package services

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
	masterModel "github.com/kalleriakronos24/mygoapp2nd/models/master"
)

func (m *module) RetrieveUser(username string) (user masterModel.User, err error) {
	if user, err = m.db.userModel.GetOneByUsername(username); err != nil {
		return masterModel.User{}, fmt.Errorf("cannot find user with username %s", username)
	}
	return
}

func (m *module) UpdateUser(id uuid.UUID, user dto.UserUpdate) (err error) {
	if err = m.db.userModel.UpdateUser(masterModel.User{
		ID:    id,
		Email: user.Email,
		Bio:   user.Bio,
	}); err != nil {
		return errors.New("cannot update user")
	}
	return
}
