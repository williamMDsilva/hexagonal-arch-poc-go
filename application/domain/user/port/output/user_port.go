package output

import (
	domain "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user"
	http_errors "github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/errors"
)

type UserPort interface {
	CreateUser(
		userDomain domain.UserDomain,
	) (*domain.UserDomain, *http_errors.HttpErr)

	FindUserByEmail(
		email string,
	) (*domain.UserDomain, *http_errors.HttpErr)
	FindUserByID(
		id string,
	) (*domain.UserDomain, *http_errors.HttpErr)
}
