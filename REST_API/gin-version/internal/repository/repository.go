package repository

import "rest-api-gin/internal/models"

type UserRepository struct {
	users  []models.User
	nextId int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []models.User{},
		nextId: 1,
	}

}
func (r *UserRepository) Create(user models.User) models.User {
	user.ID = r.nextId
	r.nextId++
	r.users = append(r.users, user)
	return user

}
func (r *UserRepository) FindByName(name string) (models.User, bool) {
	for _, u := range r.users {
		if u.Name == name {
			return u, true

		}
	}
	return models.User{}, false
}
