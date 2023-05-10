package repository

import "github.com/itsmekingtiger/error-on-generate-interface-stubs/domain"

type BaseRepo[T any, ID any] interface {
	Create(T) (ID, error)
	ReteriveById(ID) (T, error)
	DeleteById(ID) error
}

type IUserRepo interface {
	BaseRepo[domain.IUser, int]

	GetFollows(domain.IUser) ([]domain.IUser, error)
	Follow(domain.IUser, domain.IUser) error
}

var _ IUserRepo = (*UserRepo)(nil)

// Run Go: Generate Interface Stubs command with:
//
//	u *UserRepo repository.IUserRepo
//
// And also run Code Action Implement IUserRepo
type UserRepo struct{}
