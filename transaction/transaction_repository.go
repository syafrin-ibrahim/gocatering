package transaction

import (
	"gocatering/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTransaction(trs *model.Transaction) error
	UpdateTransaction(id int, trs *model.Transaction) error
	FindTransactionById(id int) (*model.Transaction, error)
	GetAllTransaction() ([]*model.Transaction, error)
	FindTransactionByUserId(id int) ([]*model.Transaction, error)
	FindPaketById(id int) (*model.Paket, error)
	FindRegencyById(id int) (*model.Regency, error)
	FindUserById(id int) (*model.User, error)
}

type TransactionRepository struct {
	conn *gorm.DB
}

func NewTransactionRepository(conn *gorm.DB) *TransactionRepository {
	return &TransactionRepository{conn: conn}
}

func (r *TransactionRepository) GetAllTransaction() ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := r.conn.Preload("Paket").Preload("User").Preload("Regency").Find(&transactions).Error
	if err != nil {
		return transactions, nil
	}
	return transactions, nil
}

func (r *TransactionRepository) FindTransactionByUserId(id int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := r.conn.Preload("Paket").Preload("User").Preload("Regency").Where("user_id=?", id).Find(&transactions).Error
	if err != nil {

		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) CreateTransaction(trs *model.Transaction) error {
	return r.conn.Create(&trs).Error
}

func (r *TransactionRepository) FindTransactionById(id int) (*model.Transaction, error) {

	var trans model.Transaction
	err := r.conn.Where("id=?", id).First(&trans).Error

	if err != nil {

		return nil, err
	}

	return &trans, nil
}
func (r *TransactionRepository) FindPaketById(id int) (*model.Paket, error) {

	var paket model.Paket
	err := r.conn.Where("id=?", id).First(&paket).Error

	if err != nil {

		return nil, err
	}

	return &paket, nil
}

func (r *TransactionRepository) FindRegencyById(id int) (*model.Regency, error) {
	var regency model.Regency
	err := r.conn.Where("id=?", id).First(&regency).Error

	if err != nil {

		return nil, err
	}

	return &regency, nil
}
func (r *TransactionRepository) FindUserById(id int) (*model.User, error) {
	var user model.User
	err := r.conn.Where("id=?", id).First(&user).Error

	if err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *TransactionRepository) UpdateTransaction(id int, trs *model.Transaction) error {
	var trans model.Transaction
	err := r.conn.Find(&trans, id).Error
	if err != nil {
		return err

	}

	return r.conn.Model(&trans).Updates(&trs).Error
}
