package application

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/tanalam2411/grpc-demo/internal/adapter/database"
	"github.com/tanalam2411/grpc-demo/internal/application/domain/bank"
	"github.com/tanalam2411/grpc-demo/internal/port"
)

type BankService struct {
	db port.BankDatabasePort
}

func NewBankService(dbPort port.BankDatabasePort) *BankService {
	return &BankService{
		db: dbPort,
	}
}

func (s *BankService) FindCurrentBalance(acct string) float64 {
	bankAccount, err := s.db.GetBankAccountByAccountNumber(acct)

	if err != nil {
		log.Println("Error on FindCurrentBalance: ", err)
	}

	return bankAccount.CurrentBalance
}

func (s *BankService) CreateExchangeRate(r bank.ExchangeRate) (uuid.UUID, error) {
	newUuid := uuid.New()
	now := time.Now()

	exchangeRateOrm := database.BankExchangeRateOrm{
		ExchangeRateUuid:   newUuid,
		FromCurrency:       r.FromCurrency,
		ToCurrency:         r.ToCurrency,
		Rate:               r.Rate,
		ValidFromTimestamp: r.ValidFromTimestamp,
		ValidToTimestamp:   r.ValidToTimestamp,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	return s.db.CreateExchangeRate(exchangeRateOrm)
}

func (s *BankService) FindExchangeRate(fromCur string, toCur string, ts time.Time) float64 {
	exchangeRate, err := s.db.GetExchangeRateAtTimestamp(fromCur, toCur, ts)

	if err != nil {
		return 0
	}

	return float64(exchangeRate.Rate)
}



func (s *BankService) CreateTransaction(acct string, t bank.Transaction) (uuid.UUID, error){
	newUuid := uuid.New()
	now := time.Now()

	bankAccountOrm, err := s.db.GetBankAccountByAccountNumber(acct)

	if err != nil {
		log.Printf("Can't create transaction for %v : %v\n", acct, err)
		return uuid.Nil, err
	}

	transactionOrm := database.BankTransactionOrm{
		TransactionUuid: newUuid,
		AccountUuid: bankAccountOrm.AccountUuid,
		TransactionTimestamp: now,
		Amount: t.Amount,
		TransactionType: t.TransactionType,
		Notes: t.Notes,
		CreatedAt: now,
		UpdatedAt: now,
	}

	savedUuid, err := s.db.CreateTransaction(bankAccountOrm, transactionOrm)

	return savedUuid, err
}


func (s *BankService) CalculateTransactionSummary(tcur *bank.TransactionSummary, trans bank.Transaction) error {

	switch trans.TransactionType {
	case bank.TransactionTypeIn:
		tcur.SumIn += trans.Amount
	case bank.TransactionTypeOut:
		tcur.SumOut += trans.Amount
	default:
		return fmt.Errorf("unknown transaction type %v", trans.TransactionType)
	}

	tcur.SumTotal = tcur.SumIn - tcur.SumOut

	return nil
}