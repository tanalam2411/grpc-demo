package database

import (
	"time"

	"github.com/google/uuid"
)


type BankAccountOrm struct {
	AccountUuid 		uuid.UUID `gorm:"primaryKey"`
	AccountNumber 		string
	AccountName 		string
	Currency            string
	CurrentBalance 		float64
	created_at 			time.Time
	updated_at 			time.Time
	Transactions 		[]BankTransactionOrm `gorm:"foreignKey:AccountUuid"`
}


func (BankAccountOrm) TableName() string {
	return "bank_accounts"
}

type BankTransactionOrm struct {
	TransactionUuid			uuid.UUID `gorm:"primaryKey"`
	AccountUuid 			uuid.UUID
	TransactionTimestamp	time.Time
	Amount					float64
	TransactionType			string
	Notes					string
	CreatedAt 				time.Time
	UpdatedAt				time.Time
}

func (BankTransactionOrm) TableName() string {
	return "bank_transactions"
}