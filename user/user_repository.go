package user

import (
	"gocatering/model"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindUserById(id int) (*model.User, error)

	UpdateUser(id int, user *model.User) error
}

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) RegisterUser(user *model.User) error {
	return r.conn.Create(&user).Error
}

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.conn.Where("email=?", email).Find(&user).Error

	if err != nil {

		return nil, err
	}

	return &user, nil
}
func (r *UserRepository) FindUserById(id int) (*model.User, error) {
	var user model.User
	err := r.conn.Where("id=?", id).Find(&user).Error

	if err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(id int, user *model.User) error {
	var fUser model.User
	err := r.conn.Find(&fUser, id).Error
	if err != nil {
		return err

	}

	return r.conn.Model(&fUser).Updates(&user).Error
}
