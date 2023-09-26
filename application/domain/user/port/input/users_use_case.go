package input

import (
	domain "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user"
	http_errors "github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/errors"
)

type UserDomainService interface {
	CreateUserServices(domain.UserDomain) (
		*domain.UserDomain, *http_errors.HttpErr)
	FindUserByIDServices(
		id string,
	) (*domain.UserDomain, *http_errors.HttpErr)
	FindUserByEmailServices(
		email string,
	) (*domain.UserDomain, *http_errors.HttpErr)
}
