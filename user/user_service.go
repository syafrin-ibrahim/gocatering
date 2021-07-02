package user

import "gocatering/model"

type Service interface {
	RegisterUser(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	//UpdateUser(id int, a *model.User) error
	//CreateUser(user *model.User) error
}

type UserService struct {
	repository Repository
}

func NewUserService(r Repository) *UserService {
	return &UserService{repository: r}
}

func (s *UserService) RegisterUser(user *model.User) error {
	return s.repository.RegisterUser(user)
}

func (s *UserService) FindUserByEmail(email string) (*model.User, error) {
	return s.repository.FindUserByEmail(email)

}

// func (s *UserService) UpdateUser(id int, user *model.User) error {
// 	return s.repository.UpdateUser(id, user)
// }
