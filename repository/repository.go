// Các interface thao tác với db cho các models

package repository

import (
	"context"
	"github.com/hoangduoc0603/moment/models"
)

type AccountRepo interface {
	Create(ctx context.Context, account *models.Account) (int64, error)
	GetByEmail(ctx context.Context, email string) (*models.Account, error)
}
