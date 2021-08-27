package users

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/andreyskoskin/drvolodko/domain"
)

type (
	Get struct {
		domain domain.GetUser
	}
)

func NewGet(d domain.GetUser) *Get {
	return &Get{d}
}

func (c *Get) Handle(e echo.Context) (err error) {
	var (
		userID   = e.Param("id")
		response domain.GetUserResponse
	)

	if response, err = c.domain.GetUser(domain.UserID(userID)); err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}
