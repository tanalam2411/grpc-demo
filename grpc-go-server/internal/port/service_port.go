package port

import (
	"time"

	"github.com/google/uuid"
	"github.com/tanalam2411/grpc-demo/internal/application/domain/bank"
)

type HelloServicePort interface {
	GenerateHello(name string) string
}

type BankServicePort interface {
	FindCurrentBalance(acct string) float64
	CreateExchangeRate(r bank.ExchangeRate) (uuid.UUID, error)
	FindExchangeRate(fromCur string, toCur string, ts time.Time) float64
	CreateTransaction(acct string, t bank.Transaction) (uuid.UUID, error)
	CalculateTransactionSummary(tcur *bank.TransactionSummary, trans bank.Transaction) error
	Transfer(tt bank.TransferTransaction) (uuid.UUID, bool, error)
}
