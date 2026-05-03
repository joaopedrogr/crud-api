package view

import (
	"crud-api/src/controller/model/response"
	"crud-api/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   int(userDomain.GetAge()),
	}
}
