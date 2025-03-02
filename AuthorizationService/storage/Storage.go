package storage

import "authorization-service/models"

type UserStorage interface {
	CreateUser(models.User)
	FindUserByLogin(string) (models.User, bool)
	FindUserById(string) (models.User, bool)
	GetAllUsersJSON() (string, bool)
	DeleteUser(id int)
}
