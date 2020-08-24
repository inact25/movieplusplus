package models

type UserModels struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	UserPass string `json:"user_pass"`
}

type AmountModels struct {
	Amount string `json:"amount"`
}

type Transaction struct {
	TransactionID     string `json:"transaction_id"`
	TransactionAmount string `json:"transaction_amount"`
	TransactionTo     string `json:"transaction_to"`
}
