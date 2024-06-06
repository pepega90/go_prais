package model

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository struct {
	DB []*User
}

func New() *UserRepository {
	return &UserRepository{
		DB: []*User{},
	}
}

func (u *UserRepository) GetAll() []*User {
	return u.DB
}

func (u *UserRepository) CreateUser(req *User) {
	u.DB = append(u.DB, req)
}

func (u *UserRepository) GetUserByID(id int) *User {
	for _, v := range u.DB {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func (u *UserRepository) UpdateUser(req *User) {
	for i, v := range u.DB {
		if v.Id == req.Id {
			req.UpdatedAt = time.Now()
			u.DB[i] = req
			break
		}
	}
}

func (u *UserRepository) DeleteUser(id int) bool {
	for i, user := range u.DB {
		if user.Id == id {
			u.DB = append(u.DB[:i], u.DB[i+1:]...)
			return true
		}
	}
	return false
}
