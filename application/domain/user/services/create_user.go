package service

import (
	domain "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user/port/input"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user/port/output"
	http_errors "github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/errors"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/logger"
	"go.uber.org/zap"
)

func NewUserDomainService(userRepository output.UserPort) input.UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	repository output.UserPort
}

func (ud *userDomainService) CreateUserServices(
	userDomain domain.UserDomain,
) (*domain.UserDomain, *http_errors.HttpErr) {

	logger.Info("Init createUser model.",
		zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.Email)
	if user != nil {
		return nil, http_errors.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.Id),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
