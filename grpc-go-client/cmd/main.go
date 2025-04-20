package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/tanalam2411/grpc-demo/internal/adapter/bank"
	"github.com/tanalam2411/grpc-demo/internal/adapter/hello"
	"github.com/tanalam2411/grpc-demo/internal/adapter/resiliency"
	dbank "github.com/tanalam2411/grpc-demo/internal/application/domain/bank"
	dresl "github.com/tanalam2411/grpc-demo/internal/application/domain/resiliency"

	grpcr "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	opts = append(opts,
		grpc.WithUnaryInterceptor(
			grpcr.UnaryClientInterceptor(
				grpcr.WithCodes(codes.Unknown, codes.Internal),
				grpcr.WithMax(4),
				grpcr.WithBackoff(grpcr.BackoffExponential(2*time.Second)),
			),
		),
	)

	opts = append(opts, 
		grpc.WithStreamInterceptor(
			grpcr.StreamClientInterceptor(
				grpcr.WithCodes(codes.Unknown, codes.Internal),
				grpcr.WithMax(4),
				grpcr.WithBackoff(grpcr.BackoffLinear(3*time.Second)),
			),
		),
	)

	conn, err := grpc.Dial("localhost:9090", opts...)

	if err != nil {
		log.Fatalln("Can not connect to gRPC server: ", err)
	}

	defer conn.Close()

	// helloAdapter, err := hello.NewHelloAdapter(conn)

	// if err != nil {
	// 	log.Fatalln("Can not create HelloAdapter: ", err)
	// }

	// bankAdapter, err := bank.NewBankAdapter(conn)

	// if err != nil{
	// 	log.Fatalln("Can not create BankAdapter: ", err)
	// }

	resiliencyAdapter, err := resiliency.NewResiliencyAdapter(conn)

	if err != nil {
		log.Fatalln("Can not create ResiliencyAdapter: ", err)
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
	// runTransferMultiple(bankAdapter, "7835697004", "7835697001", 200)

	// runUnaryResiliencyWithTiimeout(resiliencyAdapter, 0, 3, []uint32{dresl.OK}, 5 * time.Second)
	// runUnaryResiliencyWithTiimeout(resiliencyAdapter, 2, 8, []uint32{dresl.OK}, 5 * time.Second)
	// runServerStreamingResiliencyWithTimeout(resiliencyAdapter, 0, 3, []uint32{dresl.OK}, 15 * time.Second)
	// runClientStreamingResiliencyWithTimeout(resiliencyAdapter, 0, 3, []uint32{dresl.OK}, 10, 60 * time.Second)
	// runClientStreamingResiliencyWithTimeout(resiliencyAdapter, 0, 3, []uint32{dresl.OK}, 10, 3 * time.Second)
	// runBiDirectionalResiliencyWithTimeout(resiliencyAdapter, 0, 3, []uint32{dresl.OK}, 10, 6 * time.Second)

	// runUnaryResiliency(resiliencyAdapter, 0, 3, []uint32{dresl.UNKNOWN, dresl.OK})
	// runServerStreamingResiliency(resiliencyAdapter, 0, 9, []uint32{dresl.UNKNOWN})
	// runClientStreamingResiliency(resiliencyAdapter, 0, 3, []uint32{dresl.UNKNOWN}, 10)
	// runBiDirectionalResiliency(resiliencyAdapter, 0, 3, []uint32{dresl.UNKNOWN}, 10)



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

	if err != nil {
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

		if i%3 == 0 {
			ttype = dbank.TransactionTypeOut
		}

		t := dbank.Transaction{
			Amount:          float64(rand.Intn(500) + 10),
			TransactionType: ttype,
			Notes:           fmt.Sprint("Dummy transaction %v", i),
		}

		tx = append(tx, t)
	}

	adapter.SummarizeTransactions(context.Background(), acct, tx)
}

func runTransferMultiple(adapter *bank.BankAdapter, fromtAcct string, toAcct string, numDummyTransactions int) {

	var trf []dbank.TransferTransaction

	for i := 1; i <= numDummyTransactions; i++ {
		tr := dbank.TransferTransaction{
			FromAccountNumber: fromtAcct,
			ToAccountNumber:   toAcct,
			Currency:          "USD",
			Amount:            float64(rand.Intn(200) + 5),
		}

		trf = append(trf, tr)
	}

	adapter.TransferMultiple(context.Background(), trf)
}

func runUnaryResiliencyWithTiimeout(adapter *resiliency.ResiliencyAdapter, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32, timeout time.Duration) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	res, err := adapter.UnaryResiliency(ctx, minDelaySecond, maxDelaySecond, statusCodes)

	if err != nil {
		log.Fatalln("Failed to call UnaryResiliency: ", err)
	}

	log.Println(res.DummyString)
}

func runServerStreamingResiliencyWithTimeout(adapter *resiliency.ResiliencyAdapter, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32, timeout time.Duration) {

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	adapter.ServerStreamingResiliency(ctx, minDelaySecond, maxDelaySecond, statusCodes)

}

func runClientStreamingResiliencyWithTimeout(adapter *resiliency.ResiliencyAdapter, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32, count int, timeout time.Duration) {

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	adapter.ClientStreamingResiliency(ctx, minDelaySecond, maxDelaySecond, statusCodes, count)
}

func runBiDirectionalResiliencyWithTimeout(adapter *resiliency.ResiliencyAdapter, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32, count int, timeout time.Duration) {

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	adapter.BiDirectionalResiliency(ctx, minDelaySecond, maxDelaySecond, statusCodes, count)
}

// ------------------

func runUnaryResiliency(adapter *resiliency.ResiliencyAdapter, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32) {

	res, err := adapter.UnaryResiliency(context.Background(), minDelaySecond, maxDelaySecond, statusCodes)

	if err != nil {
		log.Fatalln("Failed to call UnaryResiliency: ", err)
	}

	log.Println(res.DummyString)
}

func runServerStreamingResiliency(adapter *resiliency.ResiliencyAdapter, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32) {

	adapter.ServerStreamingResiliency(context.Background(), minDelaySecond, maxDelaySecond, statusCodes)
}

func runClientStreamingResiliency(adapter *resiliency.ResiliencyAdapter, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32, count int) {

	adapter.ClientStreamingResiliency(context.Background(), minDelaySecond, maxDelaySecond, statusCodes, count)
}

func runBiDirectionalResiliency(adapter *resiliency.ResiliencyAdapter, minDelaySecond int32,
	maxDelaySecond int32, statusCodes []uint32, count int) {

	adapter.BiDirectionalResiliency(context.Background(), minDelaySecond, maxDelaySecond, statusCodes, count)
}