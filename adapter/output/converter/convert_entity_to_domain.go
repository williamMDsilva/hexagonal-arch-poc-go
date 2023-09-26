package converter

import (
	entity "github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/output/model/entity/user"
	domain "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) *domain.UserDomain {
	domainConverted := &domain.UserDomain{
		Email:    entity.Email,
		Password: entity.Password,
		Name:     entity.Name,
		Age:      entity.Age,
	}

	domainConverted.Id = entity.ID.Hex()
	return domainConverted
}
