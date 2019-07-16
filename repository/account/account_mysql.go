package account

import (
	"context"
	"database/sql"
	"github.com/hoangduoc0603/moment/models"
	"github.com/hoangduoc0603/moment/repository"
)

/*
	Để viết phương thức mở rộng cho connection thì connection và các phương thức phải đều định nghĩa ở file này
	=> Khai báo thêm struct bọc connection lại.
**/
type mysqlAccountRepo struct {
	Conn *sql.DB
}

// Lấy vào 1 thể hiện của 1 connection kết nối đến DB và trả ra connection mà đã implement các phương thức của AccountRepo interface
func NewSQLAccountRepo(conn *sql.DB) repository.AccountRepo {
	return &mysqlAccountRepo{
		Conn: conn,
	}
}

func (m *mysqlAccountRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Account, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Account, 0)

	for rows.Next() {
		data := new(models.Account)

		err := rows.Scan(
			&data.Email,
			&data.Password,
		)

		if err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, nil
}

func (m *mysqlAccountRepo) Create(ctx context.Context, account *models.Account) (int64, error) {
	query := "Insert Accounts SET email=?, password=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, account.Email, account.Password)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlAccountRepo) GetByEmail(ctx context.Context, email string) (*models.Account, error) {
	query := "Select email, password From Accounts where email=?"

	accounts, err := m.fetch(ctx, query, email)
	if err != nil {
		return nil, err
	}

	payload := &models.Account{}
	if len(accounts) > 0 {
		payload = accounts[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}
