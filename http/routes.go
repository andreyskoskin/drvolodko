package http

import (
	"fmt"
	"net/http"

	"github.com/andreyskoskin/drvolodko/domain"
	"github.com/andreyskoskin/drvolodko/http/users"
	"github.com/labstack/echo"
)

type (
	Domain interface {
		domain.ListUsers
		domain.GetUser
		domain.AddUser
		domain.KillUser
	}

	LogInternalErr func(...interface{})
)

func Routes(d Domain, log LogInternalErr) *echo.Echo {
	var e = echo.New()

	e.GET("/users", users.NewList(d).Handle)
	e.POST("/users", users.NewAdd(d).Handle)
	e.GET("/users/:id", users.NewGet(d).Handle)
	e.DELETE("/users/:id", users.NewKill(d).Handle)

	e.HTTPErrorHandler = func(err error, context echo.Context) {
		if err == domain.ErrNotFound {
			_ = context.String(http.StatusNotFound, "Not Found")
		} else if httpErr, ok := err.(*echo.HTTPError); ok {
			_ = context.String(httpErr.Code, fmt.Sprint(httpErr.Message))
		} else {
			log(err)
			_ = context.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	return e
}
