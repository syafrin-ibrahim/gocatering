package transaction

import (
	"gocatering/model"
)

type Service interface {
	GetAllTransaction() ([]*model.Transaction, error)
	CreateTransaction(trs *model.Transaction) error
	UpdateTransaction(id int, t *model.Transaction) error
	FindTransactionById(id int) (*model.Transaction, error)
	FindPaketById(id int) (*model.Paket, error)
	FindRegencyById(id int) (*model.Regency, error)
	FindUserById(id int) (*model.User, error)
	FindTransactionByUserId(id int) ([]*model.Transaction, error)
}

type TransactionService struct {
	repository Repository
}

func NewTransactionService(r Repository) *TransactionService {
	return &TransactionService{repository: r}
}

func (s *TransactionService) CreateTransaction(trs *model.Transaction) error {

	return s.repository.CreateTransaction(trs)
}

func (s *TransactionService) GetAllTransaction() ([]*model.Transaction, error) {

	return s.repository.GetAllTransaction()
}

func (s *TransactionService) FindTransactionByUserId(id int) ([]*model.Transaction, error) {
	return s.repository.FindTransactionByUserId(id)
}

func (s *TransactionService) FindTransactionById(id int) (*model.Transaction, error) {
	return s.repository.FindTransactionById(id)
}
func (s *TransactionService) FindPaketById(id int) (*model.Paket, error) {
	return s.repository.FindPaketById(id)
}

func (s *TransactionService) FindRegencyById(id int) (*model.Regency, error) {
	return s.repository.FindRegencyById(id)
}

func (s *TransactionService) FindUserById(id int) (*model.User, error) {
	return s.repository.FindUserById(id)
}

func (s *TransactionService) UpdateTransaction(id int, t *model.Transaction) error {
	return s.repository.UpdateTransaction(id, t)
}
