package database

import (
	"encoding/json"

	"github.com/ansh-devs/tc-assessment/internal/entity"
)

type Repository struct {
	Users []entity.User
}

// GetUsersListByIds
func (repo *Repository) GetUsersListByIds([]int) (users []entity.User, err error) { return users, nil }

// GetUserById
func (repo *Repository) GetUserById(userID int) (user entity.User, err error) {
	for _, v := range repo.Users {
		if v.ID == userID {
			user = v
			return
		}
	}
	return user, nil
}

// GetUsersByField
func (repo *Repository) GetUsersByField(field, value string) (users []entity.User, err error) {
	return users, nil
}

// NewUserRepository constructor for the UserRepository
func NewUserRepository() UserRepository {
	var users []entity.User

	_ = json.Unmarshal([]byte(MockedUsers), &users)
	// logrus.WithFields(logrus.Fields{"json_data": users}).Info("user_service")
	return &Repository{
		Users: users,
	}
}
