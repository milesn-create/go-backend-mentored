package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"rest-api-gin/internal/models"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByName(name string) (models.User, bool, error)
	FindByID(id int) (models.User, bool, error)
}
type InMemoryUserRepository struct {
	users  []models.User
	nextId int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  []models.User{},
		nextId: 1,
	}

}
func (r *InMemoryUserRepository) Create(user models.User) (models.User, error) {
	user.ID = r.nextId
	r.nextId++
	r.users = append(r.users, user)
	return user, nil

}
func (r *InMemoryUserRepository) FindByName(name string) (models.User, bool, error) {
	for _, u := range r.users {
		if u.Name == name {
			return u, true, nil

		}
	}
	return models.User{}, false, nil
}
func (r *InMemoryUserRepository) FindByID(id int) (models.User, bool, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, true, nil
		}
	}
	return models.User{}, false, nil
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(DB *sql.DB) *PostgresUserRepository {

	return &PostgresUserRepository{db: DB}

}

func (r *PostgresUserRepository) Create(user models.User) (models.User, error) {

	err := r.db.QueryRow("INSERT INTO users(name,age) values($1,$2) RETURNING id", user.Name, user.Age).Scan(&user.ID)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user : %w", err)

	}
	return user, nil

}

func (r *PostgresUserRepository) FindByName(name string) (models.User, bool, error) {
	var u models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE name = $1", name).Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, false, nil
		}
		return models.User{}, false, fmt.Errorf("failed to select by name: %w", err)
	}

	return u, true, nil

}
func (r *PostgresUserRepository) FindByID(id int) (models.User, bool, error) {
	var u models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, false, nil
		}
		return models.User{}, false, fmt.Errorf("failed to select by id: %w", err)
	}
	return u, true, nil
}
