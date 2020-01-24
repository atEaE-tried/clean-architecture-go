package infrastructure

import (
	"database/sql"

	"github.com/atEaE-tried/clean-architecture-go/interfaces/database"
	_ "github.com/lib/pq"
)

/*
 * DB接続は外部ライブラリを使い、外部接続を行っている。
 * 外層にあたるこの層のルールを内側層へ持ち越さないのように、ここでルールは完結させる。
 */

// DBドライバ
var driver = "postgres"

// SQLHandler : SQLへのハンドリング提供するインフラHelper
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler : SQLHanlderインスタンスを生成します。
func NewSQLHandler() *SQLHandler {
	db, err := sql.Open(driver, "postgres://test:test@db:5432/sample_db_test")
	if err != nil {
		panic(err.Error)
	}

	return &SQLHandler{Conn: db}
}

func (hndl *SQLHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SQLResult{}
	result, err := hndl.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (hndl *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := hndl.Conn.Query(statement, args...)
	if err != nil {
		return new(SQLRow), err
	}
	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}

// SQLResult : SQL実行結果
type SQLResult struct {
	Result sql.Result
}

// LastInsertId : 最後に挿入したデータのIDを取得します。
func (rlt SQLResult) LastInsertId() (int64, error) {
	return rlt.Result.LastInsertId()
}

// RowsAffected : SQLの実行により影響のあった行数を返却します。
func (rlt SQLResult) RowsAffected() (int64, error) {
	return rlt.Result.RowsAffected()
}

// SQLRow : SQL実行結果のレコード
type SQLRow struct {
	Rows *sql.Rows
}

// Scan : SQLを実行しスキャンを実行します。
func (row *SQLRow) Scan(dest ...interface{}) error {
	return row.Rows.Scan(dest...)
}

// Next : 次の要素が存在するか確認します。
func (row *SQLRow) Next() bool {
	return row.Rows.Next()
}

// Close : 接続を閉じます。
func (row *SQLRow) Close() error {
	return row.Rows.Close()
}
