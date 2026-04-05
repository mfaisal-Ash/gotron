package auth

import "errors"

type Repository interface {
	Create(User *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}

type InMemoryRepository struct {
	Users map[string]*User
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		Users: make(map[string]*User),
	}
}

func (r *InMemoryRepository) Create(user *User) error {
	for _, u := range r.users {
		if u.Email == user.Email {
			return errors.New("email already registered")
		}
	}
	r.users = append(r.users, user)
	return nil
}

func (r *InMemoryRepository) FindByEmail(email string) (*User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryRepository) FindByID(id string) (*User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}