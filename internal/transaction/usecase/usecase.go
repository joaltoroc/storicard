package usecase

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

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
func (u *usecase) ProcessFile(ctx context.Context, payload dtos.Payload) (int, string, error) {
	executionID := uuid.New().String()

	dataFile, err := u.storage.GetFile(payload.FileName)
	if err != nil {
		return http.StatusInternalServerError, executionID, err
	}

	transactions, err := extractDataCSV(dataFile, executionID)
	if err != nil {
		return http.StatusInternalServerError, executionID, err
	}

	err = u.repo.InsertData(context.Background(), transactions)
	if err != nil {
		return http.StatusInternalServerError, executionID, err
	}

	fmt.Println(transactions)

	return http.StatusOK, executionID, nil
}

// GetData implements transaction.UseCase.
func (u *usecase) GetData(ctx context.Context) ([]entities.Transaction, error) {
	return u.repo.GetData(ctx)
}

// GetDataByID implements transaction.UseCase.
func (u *usecase) GetDataByID(ctx context.Context, executionID string) ([]entities.Transaction, error) {
	return u.repo.GetDataByID(ctx, executionID)
}

func extractDataCSV(dataFile []byte, executionID string) ([]entities.Transaction, error) {
	transactions := make([]entities.Transaction, 0)
	now := time.Now()

	r := bytes.NewBuffer(dataFile)
	csvReader := csv.NewReader(r)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if rec[0] == "ID" {
			continue
		}

		id, err := strconv.ParseInt(rec[0], 10, 64)
		if err != nil {
			return nil, err
		}

		dateStr := strings.Split(rec[1], "/")
		str := fmt.Sprintf("%d-%02s-%02s", time.Now().Year(), dateStr[0], dateStr[1])
		date, err := time.Parse("2006-01-02", str)
		if err != nil {
			return nil, err
		}

		value, err := strconv.ParseFloat(rec[2], 64)
		if err != nil {
			return nil, err
		}

		typeTransaction := entities.CreditType
		if value < 0 {
			typeTransaction = entities.DebitType
		}

		transactions = append(transactions, entities.Transaction{
			FileID:          id,
			ExecutionID:     executionID,
			Date:            date,
			Value:           value,
			CreatedAt:       now,
			TypeTransaction: typeTransaction,
		})
	}

	return transactions, nil
}
