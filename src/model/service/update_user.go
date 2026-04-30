package service

import (
	"crud-api/src/configuration/rest_err"
	"crud-api/src/model"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
