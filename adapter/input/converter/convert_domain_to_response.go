package converter

import (
	response "github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/input/model/response/user"
	domain "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user"
)

func ConvertDomainToResponse(
	userDomain *domain.UserDomain,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.Id,
		Email: userDomain.Email,
		Name:  userDomain.Name,
		Age:   userDomain.Age,
	}
}
