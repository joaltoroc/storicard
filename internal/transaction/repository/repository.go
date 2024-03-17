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

func NewRepository(c *gorm.DB) transaction.Repository {
	return &repository{
		db: c,
	}
}

// InsertData implements transaction.Repository.
func (r *repository) InsertData(ctx context.Context, transactions []entities.Transaction) error {
	result := r.db.Create(transactions)

	return result.Error
}
