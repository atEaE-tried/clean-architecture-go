package usecase

import "github.com/atEaE-tried/clean-architecture-go/interfaces/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u domain.User) (err error) {
	_, err := i.UserRepository.Store(u)
	return err
}

func (i *UserInteractor) Users() (users domain.Users, err error) {
	users, err = i.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (i *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = i.UserRepository.FindById(id)
	return
}
