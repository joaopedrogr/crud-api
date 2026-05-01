package service

import (
	"crud-api/src/configuration/rest_err"
	"crud-api/src/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserDomainService(db *mongo.Database) *userDomainService {
	return &userDomainService{
		db: db,
	}
}

type userDomainService struct {
	db *mongo.Database
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
