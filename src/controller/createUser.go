package controller

import (
	"crud-api/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Rota incorreta")
	c.JSON(err.Code, err)
}
