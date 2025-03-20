package application

import (
	"log"

	"github.com/tanalam2411/grpc-demo/internal/port"
)

type BankService struct{
	db port.BankDatabasePort
}

func NewBankService(dbPort port.BankDatabasePort) *BankService {
	return &BankService{
		db: dbPort,
	}
}


func (s *BankService) FindCurrentBalance(acct string) float64{
	bankAccount, err := s.db.GetBankAccountByAccountNumber(acct)

	if err != nil{
		log.Println("Error on FindCurrentBalance: ", err)
	}
	
	return bankAccount.CurrentBalance
}