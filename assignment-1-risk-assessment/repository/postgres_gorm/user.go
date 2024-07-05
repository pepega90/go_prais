package postgres_gorm

import (
	"assignment_1/entity"
	"assignment_1/service"
	"context"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db *gorm.DB) service.IUserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	res := u.db.Create(user)
	if res.Error != nil {
		log.Fatalf("Error create user: %v", res.Error)
		return entity.User{}, res.Error
	}
	return entity.User{ID: user.ID}, nil
}

func (u *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	err := u.db.Table("users").
		Select("users.id, users.name, users.email, submissions.risk_score, submissions.risk_category, submissions.risk_definition, users.created_at, users.updated_at").
		Joins("left join submissions on submissions.user_id = users.id").Order("submissions.id desc").First(&user, id).Error
	if err != nil {
		log.Fatalf("Error get user with id = %v", id)
		return entity.User{}, nil
	}
	return user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	updatedUser, _ := u.GetUserByID(ctx, id)
	updatedUser.Name = user.Name
	updatedUser.Email = user.Email
	err := u.db.Save(updatedUser).Error
	if err != nil {
		log.Fatalf("Error updating user with id = %v", id)
		return entity.User{}, nil
	}
	return updatedUser, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	err := u.db.Delete(&entity.User{}, id).Error
	if err != nil {
		log.Fatalf("error deleting user with id = %v", id)
		return err
	}
	return nil
}

func (u *userRepository) GetUserCount(ctx context.Context) (int, error) {
	var count int64
	err := u.db.Model(&entity.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (u *userRepository) GetAllUsers(ctx context.Context, limit, offset int) ([]entity.User, error) {
	var listUser []entity.User
	err := u.db.Table("users").
		Select("users.id, users.name, users.email, submissions.risk_score, submissions.risk_category, submissions.risk_definition, users.created_at, users.updated_at").
		Joins("left join submissions on submissions.user_id = users.id").Limit(limit).Offset(offset).Find(&listUser).Error
	if err != nil {
		log.Fatalf("error get all users: %v", err.Error())
		return nil, err
	}
	return listUser, nil
}
