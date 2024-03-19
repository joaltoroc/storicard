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

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/google/uuid"

	"github/joaltoroc/storicard/internal/transaction"
	"github/joaltoroc/storicard/internal/transaction/dtos"
	"github/joaltoroc/storicard/internal/transaction/entities"
)

type (
	usecase struct {
		repo         transaction.Repository
		storage      transaction.Storage
		notification transaction.Notification
	}

	Month struct {
		Month string
		Count int
	}

	Months []Month
)

func NewUseCase(repo transaction.Repository, storage transaction.Storage, notification transaction.Notification) transaction.UseCase {
	return &usecase{repo, storage, notification}
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

	var body []map[string]string

	total := calculateTotal(transactions)
	body = appendResult(body, "Total balance is", total)

	averageDebit := calculateAverage(transactions, entities.DebitType)
	body = appendResult(body, "Average debit amount", averageDebit)

	averageCredit := calculateAverage(transactions, entities.CreditType)
	body = appendResult(body, "Average credit amount", averageCredit)

	months := groupByMonth(transactions)
	for _, result := range months {
		body = appendResult(body, fmt.Sprintf("Number of transactions in %s", result.Month), float64(result.Count))
	}

	err = u.notification.SendMail(payload.Email, payload.FileName, executionID, body)
	if err != nil {
		return http.StatusInternalServerError, executionID, err
	}

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

func calculateTotal(transactions []entities.Transaction) float64 {
	return linq.From(transactions).Select(func(t interface{}) interface{} {
		return t.(entities.Transaction).Value
	}).SumFloats()
}

func calculateAverage(transactions []entities.Transaction, transactionType entities.TransactionType) float64 {
	filterFunc := func(t interface{}) bool {
		return t.(entities.Transaction).TypeTransaction == transactionType
	}

	return linq.From(transactions).Where(filterFunc).Select(func(t interface{}) interface{} {
		return t.(entities.Transaction).Value
	}).Average()
}

func appendResult(body []map[string]string, text string, value float64) []map[string]string {
	return append(body, map[string]string{
		"text":  text,
		"value": strconv.FormatFloat(value, 'f', -1, 64),
	})
}

func groupByMonth(transactions []entities.Transaction) Months {
	query := linq.From(transactions).GroupBy(func(t interface{}) interface{} {
		return t.(entities.Transaction).Date.Month()
	}, func(t interface{}) interface{} {
		return t.(entities.Transaction).Date.Month()
	}).Select(func(group interface{}) interface{} {
		month := fmt.Sprintf("%s", group.(linq.Group).Key)
		count := len(group.(linq.Group).Group)

		return Month{
			Month: month,
			Count: count,
		}
	})

	var months Months

	query.ToSlice(&months)

	return months
}
