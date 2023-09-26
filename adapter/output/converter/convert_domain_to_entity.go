package converter

import (
	entity "github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/output/model/entity/user"
	domain "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user"
)

func ConvertDomainToEntity(
	domain domain.UserDomain,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.Email,
		Password: domain.Password,
		Name:     domain.Name,
		Age:      domain.Age,
	}
}
