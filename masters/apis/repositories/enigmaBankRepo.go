package repositories

import "github.com/inact25/movieplusplus/masters/apis/models"

type BankRepositories interface {
	OAuth(userName, userPass string) ([]*models.UserModels, error)
	GetAmount(userName string) ([]*models.AmountModels, error)
	GetTransaction(userName string) ([]*models.Transaction, error)

	AddAmount(amount, userName string) (string, error)
	AddTransaction(amount, userName string) (string, error)
}
