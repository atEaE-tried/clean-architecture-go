package domain

import "testing"

// インスタンス生成および生成される値についてチェックします。
func TestInstance(t *testing.T) {
	t.Run("インスタンス生成チェック", func(t *testing.T) {
		// setup
		id := 1
		fname := "test"
		lname := "user"

		// check
		user := &User{
			ID:        id,
			FirstName: fname,
			LastName:  lname,
		}

		// assert
		if user == nil {
			t.Error("インスタンスが生成されていません。")
		}
	})
}
