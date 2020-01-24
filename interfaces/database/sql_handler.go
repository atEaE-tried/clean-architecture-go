package database

// SQLHandler : インフラ層に対するアダプタ。
type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

// Result : SQLハンドラ実行結果(インフラ層に対するアダプタ)
type Result interface {
	LastInsertId() (int64, error)
	RowAffected() (int64, error)
}

// Row : SQLハンドラ実行結果(インフラ層に対するアダプタ)
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
