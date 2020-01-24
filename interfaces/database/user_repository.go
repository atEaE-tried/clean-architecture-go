package database

import (
	"github.com/atEaE-tried/clean-architecture-go/domain"
)

type UserRepository struct {
	SQLHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return -1, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	id = int(id64)
	return id, nil
}

func (repo *UserRepository) FindById(userID int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", userID)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var firstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	return
}
