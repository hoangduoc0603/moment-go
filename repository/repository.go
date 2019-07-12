// Các interface thao tác với db cho các models

package repository

import (
	"context"
)

type AccountRepo interface {
	Create(ctx context.Context, uname, pass string)
}
