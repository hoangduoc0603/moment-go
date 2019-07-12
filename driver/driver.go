package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// DB ... Struct giữ tham chiếu đến 1 thể hiện của SQL database
type DB struct {
	SQL *sql.DB
}

// DB connection -> lấy về address của connection
var dbConn = &DB{}

// ConnectSQL ... Thiết lập 1 kết nối đế database và trả về lỗi nếu có
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)

	d, err := sql.Open("mysql", dbSource)
	dbConn.SQL = d
	return dbConn, err
}
