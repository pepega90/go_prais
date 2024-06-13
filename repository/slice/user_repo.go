package slice

import (
	"errors"
	"go_prais/model"
	"time"
)

type UserRepository struct {
	db     []*model.User
	nextId int
}

func NewSliceRepository() *UserRepository {
	return &UserRepository{
		db:     []*model.User{},
		nextId: 1,
	}
}

func (u *UserRepository) GetAllUsers() []*model.User {
	return u.db
}

func (u *UserRepository) CreateUser(req model.User) *model.User {
	req.Id = u.nextId
	u.nextId++
	req.CreatedAt = time.Now().UTC()
	req.UpdatedAt = time.Now().UTC()
	u.db = append(u.db, &req)
	return &req
}

func (u *UserRepository) GetUser(id int) (*model.User, error) {
	for _, v := range u.db {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, errors.New("not found")
}

func (u *UserRepository) UpdateUser(id int, req model.User) (*model.User, error) {
	for i, v := range u.db {
		if v.Id == id {
			req.Id = id
			req.UpdatedAt = time.Now()
			u.db[i] = &req
			break
		}
	}
	return &req, nil
}

func (u *UserRepository) DeleteUser(id int) error {
	for i, user := range u.db {
		if user.Id == id {
			u.db = append(u.db[:i], u.db[i+1:]...)
			return nil
		}
	}
	return errors.New("error while deleting user")
}
