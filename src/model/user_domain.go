package model

import (
	"crud-api/src/configuration/rest_err"
	"crypto/md5"
	"encoding/hex"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	return nil
}

func (ud *userDomain) UpdateUser(s string, domain userDomain) *rest_err.RestErr {
	return nil
}

func (ud *userDomain) FindUser(s string) (*userDomain, *rest_err.RestErr) {
	return nil, nil
}

func (ud *userDomain) DeleteUser(s string) *rest_err.RestErr {
	return nil
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string, userDomain) *rest_err.RestErr
	FindUser(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
	GetPassword() string
	EncryptPassword()
}
