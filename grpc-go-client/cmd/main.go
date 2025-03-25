package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/tanalam2411/grpc-demo/internal/adapter/bank"
	"github.com/tanalam2411/grpc-demo/internal/adapter/hello"
	dbank "github.com/tanalam2411/grpc-demo/internal/application/domain/bank"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:9090", opts...)

	if err != nil{
		log.Fatalln("Can not connect to gRPC server: ", err)
	}

	defer conn.Close()
 
	// helloAdapter, err := hello.NewHelloAdapter(conn)

	// if err != nil {
	// 	log.Fatalln("Can not create HelloAdapter: ", err)
	// }

	bankAdapter, err := bank.NewBankAdapter(conn)

	if err != nil{
		log.Fatalln("Can not create BankAdapter: ", err)
	}


	// runSayHello(helloAdapter, "Bruce  W.")
	// runSayManyHellos(helloAdapter, "Stream msg ...")
	// runSayHelloToEveryOne(helloAdapter, []string{"Tan", "San", "Nab", "Man"})
	// runSayHelloContinuous(helloAdapter, []string{"Tan", "San", "Nab", "Man"})

	// runGetCurrentBalance(bankAdapter, "7835697001")
	// runGetCurrentBalance(bankAdapter, "7835697001-000")

	// runFetchExchangeRates(bankAdapter, "USD", "TDR")
	// runFetchExchangeRates(bankAdapter, "USD", "IDR")
	// runSummarizeTransactions(bankAdapter, "7835697002", 10 )
	// runTransferMultiple(bankAdapter, "7835697004", "7835697001", 10)
	runTransferMultiple(bankAdapter, "7835697004", "7835697001", 200)


}


func runSayHello(adapter *hello.HelloAdapter, name string) {
	greet, err := adapter.SayHello(context.Background(), name)

	if err != nil {
		log.Fatalln("Can not call SayHello: ", err)
	}

	log.Println(greet.Greet)
}

func runSayManyHellos(adapter *hello.HelloAdapter, name string) {
	adapter.SayManyHellos(context.Background(), name)
}

func runSayHelloToEveryOne(adapter *hello.HelloAdapter, names []string) {
	adapter.SayHelloToEveryOne(context.Background(), names)
}

func runSayHelloContinuous(adapter *hello.HelloAdapter, names []string) {
	adapter.SayHelloContinuous(context.Background(), names)
}


func runGetCurrentBalance(adapter *bank.BankAdapter, acct string) {
	bal, err := adapter.GetCurrentBalance(context.Background(), acct)

	if err != nil{
		log.Fatalln("Failed to call GetCurrentBalance: ", err)
	}

	log.Println(bal)
}

 func runFetchExchangeRates(adapter *bank.BankAdapter, fromCur string, toCur string) {
	adapter.FetchExchangeRates(context.Background(), fromCur, toCur)
 }


 func runSummarizeTransactions(adapter *bank.BankAdapter, acct string, numDummyTransactions int) {
	var tx []dbank.Transaction

	for i := 1; i <= numDummyTransactions; i++ {
		ttype := dbank.TransactionTypeIn

		if i%3 == 0{
			ttype = dbank.TransactionTypeOut
		}

		t := dbank.Transaction{
			Amount: float64(rand.Intn(500) + 10),
			TransactionType: ttype,
			Notes: fmt.Sprint("Dummy transaction %v", i),
		}

		tx = append(tx, t)
	}

	adapter.SummarizeTransactions(context.Background(), acct, tx)
 }


 func runTransferMultiple(adapter *bank.BankAdapter, fromtAcct string, toAcct string, numDummyTransactions int) {

	var trf []dbank.TransferTransaction

	for i := 1; i <= numDummyTransactions; i++{
		tr := dbank.TransferTransaction{
			FromAccountNumber: fromtAcct,
			ToAccountNumber: toAcct,
			Currency: "USD",
			Amount: float64(rand.Intn(200) + 5),
		}

		trf = append(trf, tr)
	}

	adapter.TransferMultiple(context.Background(), trf)
 }