package service

import (
	domain "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user"
	http_errors "github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/errors"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (*domain.UserDomain, *http_errors.HttpErr) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserById"))

	return ud.repository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (*domain.UserDomain, *http_errors.HttpErr) {
	logger.Info("Init findUserByEmail services.",
		zap.String("journey", "findUserById"))

	return ud.repository.FindUserByEmail(email)
}
