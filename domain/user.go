package domain

/*
 * ビジネスルールの中核となるユーザーエンティティ
 * 「User」モデルはID、姓、名をもつ
 * 「Users」は「Userエンティティ」の集合エンティティで、配列として公開するのではなく、
 * type宣言して、集合型として公開する
 */

// User : ユーザーエンティティ
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// Users : ユーザー集合エンティティ
type Users []User
