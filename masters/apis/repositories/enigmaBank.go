package repositories

import (
	"database/sql"
	"github.com/inact25/movieplusplus/masters/apis/models"
	"log"
)

type BankRepoImpl struct {
	db *sql.DB
}

func (b BankRepoImpl) OAuth(userName, userPass string) ([]*models.UserModels, error) {
	var dataOAuth []*models.UserModels
	data := models.UserModels{}
	query := "select userID, userName from user where userName = ? and userPass = ?;"
	row := b.db.QueryRow(query, userName, userPass)
	if err := row.Scan(&data.UserId, &data.UserName); err != nil {
		return nil, err
	}
	dataOAuth = append(dataOAuth, &data)
	return dataOAuth, nil
}

func (b BankRepoImpl) GetAmount(userName string) ([]*models.AmountModels, error) {
	var dataAmount []*models.AmountModels
	query := "select amount from amount where userName = ?"
	data, err := b.db.Query(query, userName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		amount := models.AmountModels{}
		err := data.Scan(&amount.Amount)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataAmount = append(dataAmount, &amount)
	}
	return dataAmount, nil
}

func (b BankRepoImpl) GetTransaction(userName string) ([]*models.Transaction, error) {
	var dataTransaction []*models.Transaction
	query := "select transactionID, transactionAmount, transactionTo from transaction where userName = ?;"
	data, err := b.db.Query(query, userName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		transaction := models.Transaction{}
		err := data.Scan(&transaction.TransactionID, &transaction.TransactionAmount, &transaction.TransactionTo)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataTransaction = append(dataTransaction, &transaction)
	}
	return dataTransaction, nil
}

func (b BankRepoImpl) AddAmount(amount, userName string) (string, error) {
	panic("implement me")
}

func (b BankRepoImpl) AddTransaction(amount, userName string) (string, error) {
	panic("implement me")
}

func InitMovieRepoImpl(db *sql.DB) BankRepositories {
	return &BankRepoImpl{db}
}
