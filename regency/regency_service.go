package regency

import (
	"gocatering/model"
)

type Service interface {
	CreateRegency(a *model.Regency) error
	GetRegencyByID(id int) (*model.Regency, error)
	GetAllRegency() ([]*model.Regency, error)
	UpdateRegency(id int, a *model.Regency) error
	DeleteRegency(id int) (*model.Regency, error)
}
type RegencyService struct {
	repository Repository
}

func NewRegencyService(r Repository) *RegencyService {
	return &RegencyService{repository: r}
}

func (s *RegencyService) CreateRegency(a *model.Regency) error {

	return s.repository.CreateRegency(a)
}

func (s *RegencyService) GetRegencyByID(id int) (*model.Regency, error) {
	return s.repository.GetRegencyByID(id)
}

func (s *RegencyService) GetAllRegency() ([]*model.Regency, error) {
	return s.repository.GetAllRegency()
}

func (s *RegencyService) UpdateRegency(id int, a *model.Regency) error {

	return s.repository.UpdateRegency(id, a)
}

func (s *RegencyService) DeleteRegency(id int) (*model.Regency, error) {
	regency, err := s.repository.GetRegencyByID(id)
	if err != nil {
		return regency, err
	}

	deletedRegency, err := s.repository.DeleteRegency(regency)
	if err != nil {
		return deletedRegency, err
	}

	return deletedRegency, nil
}
