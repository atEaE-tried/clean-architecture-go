package usecase

import (
	"github.com/atEaE-tried/clean-architecture-go/interfaces/domain"
)

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
