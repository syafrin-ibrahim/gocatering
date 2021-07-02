package paket

import (
	"gocatering/model"
	"log"
	"os"
)

type Service interface {
	CreatePaket(a *model.Paket) error
	GetPaketByID(id int) (*model.Paket, error)
	GetAllPaket() ([]*model.Paket, error)
	UpdatePaket(id int, a *model.Paket) error
	DeletePaket(id int) (*model.Paket, error)
	CreateImage(a *model.Image) error
	DeleteImage(id int) (*model.Image, error)
	DetailPaket(id int) (*model.Paket, error)
}

type PaketService struct {
	repository Repository
}

func NewPaketService(r Repository) *PaketService {
	return &PaketService{repository: r}
}

func (s *PaketService) CreatePaket(paket *model.Paket) error {
	return s.repository.CreatePaket(paket)
}

func (s *PaketService) GetPaketByID(id int) (*model.Paket, error) {
	return s.repository.GetPaketByID(id)
}

func (s *PaketService) DetailPaket(id int) (*model.Paket, error) {
	return s.repository.DetailPaket(id)
}

func (s *PaketService) GetAllPaket() ([]*model.Paket, error) {
	return s.repository.GetAllPaket()
}

func (s *PaketService) UpdatePaket(id int, a *model.Paket) error {

	return s.repository.UpdatePaket(id, a)
}

func (s *PaketService) DeletePaket(id int) (*model.Paket, error) {
	paket, err := s.repository.GetPaketByID(id)
	if err != nil {
		return paket, err
	}

	deletedPaket, err := s.repository.DeletePaket(paket)
	if err != nil {
		return deletedPaket, err
	}

	return deletedPaket, nil
}

func (s *PaketService) CreateImage(a *model.Image) error {

	return s.repository.CreateImage(a)

}

func (s *PaketService) DeleteImage(id int) (*model.Image, error) {
	image, err := s.repository.GetImageByID(id)
	if err != nil {
		return image, err
	}

	imageResource := image.FileName

	if FileExists(imageResource) {
		e := os.Remove(imageResource)
		if e != nil {
			log.Fatal(e)
		}
	}

	deletedImage, err := s.repository.DeleteImage(image)
	if err != nil {
		return deletedImage, err
	}

	return deletedImage, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
