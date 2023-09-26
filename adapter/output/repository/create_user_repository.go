package repository

import (
	"context"
	"os"

	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/logger"

	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/adapter/output/converter"
	domain "github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user"
	"github.com/williamMDsilva/hexagonal-arch-poc-go.git/application/domain/user/port/output"
	http_errors "github.com/williamMDsilva/hexagonal-arch-poc-go.git/configuration/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(
	database *mongo.Database,
) output.UserPort {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func (ur *userRepository) CreateUser(
	userDomain domain.UserDomain,
) (*domain.UserDomain, *http_errors.HttpErr) {
	logger.Info("Init createUser repository",
		zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, http_errors.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"))

	return converter.ConvertEntityToDomain(*value), nil
}
