package main

import (
	"context"

	"github.com/naufalfmm/csv-transaction/models/dto"
	"github.com/naufalfmm/csv-transaction/persistents"
	"github.com/naufalfmm/csv-transaction/resources/config"
	"github.com/naufalfmm/csv-transaction/resources/db"
	"github.com/naufalfmm/csv-transaction/resources/log"
	"github.com/naufalfmm/csv-transaction/usecases"
)

func main() {
	// file, err := os.Open("data.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// balance := 0
	// failedAmountTotal := 0

	// csvReader := csv.NewReader(file)
	// for {
	// 	rec, err := csvReader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	amount, err := strconv.Atoi(rec[3])
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if rec[2] == "DEBIT" {
	// 		amount = amount * -1
	// 	}

	// 	if rec[4] == "FAILED" {
	// 		failedAmountTotal += amount
	// 		continue
	// 	}

	// 	if rec[4] == "PENDING" {
	// 		continue
	// 	}

	// 	balance += amount

	// }

	// fmt.Println(balance, failedAmountTotal)

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	lg, err := log.NewLogger(conf)
	if err != nil {
		panic(err)
	}

	db, err := db.NewMysql(conf, lg)
	if err != nil {
		panic(err)
	}

	persists, err := persistents.Init(db, lg, conf)
	if err != nil {
		panic(err)
	}

	uscs, err := usecases.Init(persists, db)
	if err != nil {
		panic(err)
	}

	req := dto.RecordTransactionRequest{
		Filename: conf.CsvFilename,
	}

	if err := uscs.Transactions.Record(context.Background(), req); err != nil {
		panic(err)
	}
}
