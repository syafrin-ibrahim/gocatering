package category

import (
	"gocatering/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCategory(*model.Category) error
	GetCategoryByID(id int) (*model.Category, error)
	GetAllCategory() ([]*model.Category, error)
	UpdateCategory(id int, Category *model.Category) error
	DeleteCategory(*model.Category) (*model.Category, error)
}

type CategoryRepository struct {
	conn *gorm.DB
}

func NewCategoryRepository(c *gorm.DB) *CategoryRepository {
	return &CategoryRepository{conn: c}
}

func (r *CategoryRepository) CreateCategory(category *model.Category) error {
	return r.conn.Create(&category).Error
}

func (r CategoryRepository) UpdateCategory(id int, category *model.Category) error {
	var foundCategory model.Category
	err := r.conn.Find(&foundCategory, id).Error
	if err != nil {
		return err

	}

	return r.conn.Model(&foundCategory).Updates(&category).Error

	//return r.conn.Save(&Category).Error
}

func (r *CategoryRepository) GetCategoryByID(id int) (*model.Category, error) {
	var category model.Category
	err := r.conn.Where("id=?", id).First(&category).Error

	if err != nil {

		return nil, err
	}

	return &category, nil

}

func (r *CategoryRepository) GetAllCategory() ([]*model.Category, error) {
	var categories []*model.Category
	err := r.conn.Find(&categories).Error
	if err != nil {
		return categories, nil
	}
	return categories, nil
}

func (r *CategoryRepository) DeleteCategory(category *model.Category) (*model.Category, error) {
	err := r.conn.Delete(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
