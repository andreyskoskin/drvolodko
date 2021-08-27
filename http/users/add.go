package users

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/andreyskoskin/drvolodko/domain"
)

type (
	Add struct {
		domain domain.AddUser
	}
)

func NewAdd(d domain.AddUser) *Add {
	return &Add{d}
}

func (c *Add) Handle(e echo.Context) (err error) {
	var request domain.AddUserRequest

	if err = e.Bind(&request); err != nil {
		return err
	}

	var response domain.AddUserResponse
	if response, err = c.domain.AddUser(domain.AddUserRequest{
		Name: request.Name,
	}); err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}
