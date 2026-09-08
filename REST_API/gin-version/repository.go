package main

type UserRepository struct {
	users  []User
	nextId int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  []User{},
		nextId: 1,
	}

}
func (r *UserRepository) Create(user User) User {
	user.ID = r.nextId
	r.nextId++
	r.users = append(r.users, user)
	return user

}
func (r *UserRepository) FindByName(name string) (User, bool) {
	for _, u := range r.users {
		if u.Name == name {
			return u, true

		}
	}
	return User{}, false
}
