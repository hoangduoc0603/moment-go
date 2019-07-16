package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hoangduoc0603/moment/driver"
	"github.com/hoangduoc0603/moment/models"
	"github.com/hoangduoc0603/moment/repository"
	"github.com/hoangduoc0603/moment/repository/account"
	"net/http"
)

// Bọc lại thể hiện của đối tượng implement AccountRepo để viết thêm phương thức
type Account struct {
	repo repository.AccountRepo
}

// Lấy vào 1 db connection và trả ra Account đã được implement các func liên quan tới account
func NewPostHandler(db *driver.DB) *Account {
	return &Account{
		repo: account.NewSQLAccountRepo(db.SQL),
	}
}

// Tạo tài khoản
func (a *Account) Create(w http.ResponseWriter, r *http.Request) {
	account := models.Account{}
	json.NewDecoder(r.Body).Decode(&account)

	newID, err := a.repo.Create(r.Context(), &account)
	fmt.Println(newID)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	} else {
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
