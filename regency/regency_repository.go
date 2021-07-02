package regency

import (
	"gocatering/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateRegency(regency *model.Regency) error
	GetRegencyByID(id int) (*model.Regency, error)
	GetAllRegency() ([]*model.Regency, error)
	UpdateRegency(id int, regency *model.Regency) error
	DeleteRegency(regency *model.Regency) (*model.Regency, error)
}

type RegencyRepository struct {
	conn *gorm.DB
}

func NewRegencyRepository(conn *gorm.DB) *RegencyRepository {
	return &RegencyRepository{conn: conn}
}

//create regency
func (r *RegencyRepository) CreateRegency(regency *model.Regency) error {
	return r.conn.Create(&regency).Error

}

func (r RegencyRepository) UpdateRegency(id int, regency *model.Regency) error {
	var foundRegency model.Regency
	err := r.conn.Find(&foundRegency, id).Error
	if err != nil {
		return err

	}

	return r.conn.Model(&foundRegency).Updates(&regency).Error

	//return r.conn.Save(&regency).Error
}

func (r *RegencyRepository) GetRegencyByID(id int) (*model.Regency, error) {
	var regency model.Regency
	err := r.conn.Where("id=?", id).First(&regency).Error

	if err != nil {

		return nil, err
	}

	return &regency, nil

}

func (r *RegencyRepository) GetAllRegency() ([]*model.Regency, error) {
	var regencies []*model.Regency
	err := r.conn.Find(&regencies).Error
	if err != nil {
		return regencies, nil
	}
	return regencies, nil
}

func (r *RegencyRepository) DeleteRegency(regency *model.Regency) (*model.Regency, error) {
	err := r.conn.Delete(&regency).Error

	if err != nil {
		return regency, err
	}

	return regency, nil
}
