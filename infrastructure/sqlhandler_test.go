package infrastructure

import "testing"

// インスタンス生成および生成される値についてチェックします。
func TestNewSQLHandler(t *testing.T) {
	t.Run("インスタンスが生成されていること。", func(t *testing.T) {
		sql := NewSQLHandler()
		if sql == nil {
			t.Error("インスタンスが生成されていません。")
		}
	})
}
