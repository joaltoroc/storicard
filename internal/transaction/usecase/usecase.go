package usecase

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github/joaltoroc/storicard/internal/transaction"
	"github/joaltoroc/storicard/internal/transaction/dtos"
	"github/joaltoroc/storicard/internal/transaction/entities"
)

type (
	usecase struct {
		repo    transaction.Repository
		storage transaction.Storage
	}
)

func NewUseCase(repo transaction.Repository, storage transaction.Storage) transaction.UseCase {
	return &usecase{repo, storage}
}

// ProcessFile implements transaction.UseCase.
func (u *usecase) ProcessFile(ctx context.Context, payload dtos.Payload) (httpCode int, err error) {
	dataFile, err := u.storage.GetFile(payload.FileName)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	transactions := make([]entities.Transaction, 0)

	r := bytes.NewBuffer(dataFile)
	csvReader := csv.NewReader(r)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if rec[0] == "ID" {
			continue
		}

		id, err := strconv.ParseInt(rec[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		dateStr := strings.Split(rec[1], "/")
		str := fmt.Sprintf("%d-%02s-%02s", time.Now().Year(), dateStr[0], dateStr[1])
		date, err := time.Parse("2006-01-02", str)
		if err != nil {
			log.Fatal(err)
		}

		value, err := strconv.ParseFloat(rec[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		transactions = append(transactions, entities.Transaction{
			ID:    id,
			Date:  date,
			Value: value,
		})
	}

	err = u.repo.InsertData(context.Background(), transactions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(transactions)

	return http.StatusOK, nil
}

// GetData implements transaction.UseCase.
func (u *usecase) GetData(ctx context.Context) ([]entities.Transaction, error) {
	return u.repo.GetData(ctx)
}
