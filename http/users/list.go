package users

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/andreyskoskin/drvolodko/domain"
)

type (
	List struct {
		domain domain.ListUsers
	}
)

func NewList(d domain.ListUsers) *List {
	return &List{d}
}

func (c *List) Handle(e echo.Context) (err error) {
	var response domain.ListUsersResponse
	if response, err = c.domain.ListUsers(); err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}
