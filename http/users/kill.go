package users

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/andreyskoskin/drvolodko/domain"
)

type (
	Kill struct {
		domain domain.KillUser
	}
)

func NewKill(d domain.KillUser) *Kill {
	return &Kill{d}
}

func (c *Kill) Handle(e echo.Context) (err error) {
	var userID = e.Param("id")

	if err = c.domain.KillUser(domain.UserID(userID)); err != nil {
		return err
	}

	return e.String(http.StatusOK, "OK")
}
