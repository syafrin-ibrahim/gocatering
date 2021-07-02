package category

import (
	"gocatering/model"
)

type Service interface {
	CreateCategory(*model.Category) error
	GetCategoryByID(id int) (*model.Category, error)
	GetAllCategory() ([]*model.Category, error)
	UpdateCategory(id int, a *model.Category) error
	DeleteCategory(id int) (*model.Category, error)
}

type CategoryService struct {
	repo Repository
}

func NewCategoryService(r Repository) *CategoryService {
	return &CategoryService{repo: r}
}
func (s *CategoryService) CreateCategory(c *model.Category) error {
	return s.repo.CreateCategory(c)
}

func (s *CategoryService) GetCategoryByID(id int) (*model.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService) GetAllCategory() ([]*model.Category, error) {
	return s.repo.GetAllCategory()
}

func (s *CategoryService) UpdateCategory(id int, a *model.Category) error {

	return s.repo.UpdateCategory(id, a)
}

func (s *CategoryService) DeleteCategory(id int) (*model.Category, error) {
	Category, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return Category, err
	}

	deletedCategory, err := s.repo.DeleteCategory(Category)
	if err != nil {
		return deletedCategory, err
	}

	return deletedCategory, nil
}
