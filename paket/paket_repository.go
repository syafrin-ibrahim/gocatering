package paket

import (
	"gocatering/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreatePaket(paket *model.Paket) error
	GetPaketByID(id int) (*model.Paket, error)
	GetAllPaket() ([]*model.Paket, error)
	UpdatePaket(id int, paket *model.Paket) error
	DeletePaket(paket *model.Paket) (*model.Paket, error)
	CreateImage(img *model.Image) error
	GetImageByID(id int) (*model.Image, error)
	DeleteImage(image *model.Image) (*model.Image, error)
}

type PaketRepository struct {
	conn *gorm.DB
}

func NewPaketRepository(conn *gorm.DB) *PaketRepository {
	return &PaketRepository{conn: conn}
}
func (r *PaketRepository) CreatePaket(paket *model.Paket) error {
	return r.conn.Create(&paket).Error
}

func (r *PaketRepository) UpdatePaket(id int, paket *model.Paket) error {
	var foundPaket model.Paket
	err := r.conn.Find(&foundPaket, id).Error
	if err != nil {
		return err

	}

	return r.conn.Model(&foundPaket).Updates(&paket).Error

	//return r.conn.Save(&Paket).Error
}

func (r *PaketRepository) GetPaketByID(id int) (*model.Paket, error) {
	var paket model.Paket
	err := r.conn.Where("id=?", id).First(&paket).Error

	if err != nil {

		return nil, err
	}

	return &paket, nil

}

func (r *PaketRepository) GetAllPaket() ([]*model.Paket, error) {
	var pakets []*model.Paket
	err := r.conn.Find(&pakets).Error
	if err != nil {
		return pakets, nil
	}
	return pakets, nil
}

func (r *PaketRepository) GetImageByID(id int) (*model.Image, error) {
	var image model.Image
	err := r.conn.Where("id=?", id).First(&image).Error

	if err != nil {

		return nil, err
	}

	return &image, nil
}

func (r *PaketRepository) DeletePaket(paket *model.Paket) (*model.Paket, error) {
	err := r.conn.Delete(&paket).Error

	if err != nil {
		return paket, err
	}

	return paket, nil
}

func (r *PaketRepository) CreateImage(img *model.Image) error {
	return r.conn.Create(&img).Error
}

func (r *PaketRepository) DeleteImage(image *model.Image) (*model.Image, error) {
	err := r.conn.Delete(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}
