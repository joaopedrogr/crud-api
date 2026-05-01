package repository

import (
	"context"
	"crud-api/src/configuration/logger"
	"crud-api/src/configuration/rest_err"
	"crud-api/src/model"
	"crud-api/src/model/repository/entity"
	"crud-api/src/model/repository/entity/converter"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByEmail"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmail"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByID"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByID"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}
