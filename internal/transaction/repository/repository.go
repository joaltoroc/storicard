package repository

import (
	"context"
	"github/joaltoroc/storicard/internal/transaction"
	"github/joaltoroc/storicard/internal/transaction/entities"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) transaction.Repository {
	return &repository{
		db,
	}
}

// InsertData implements transaction.Repository.
func (r *repository) InsertData(ctx context.Context, transactions []entities.Transaction) error {
	result := r.db.Create(transactions)

	return result.Error
}

func (r *repository) GetData(ctx context.Context) ([]entities.Transaction, error) {
	transactions := make([]entities.Transaction, 0)
	result := r.db.Find(&transactions)

	return transactions, result.Error
}
