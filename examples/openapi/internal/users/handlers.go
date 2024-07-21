package users

import (
	"database/sql"
	"net/http"

	"github.com/canidam/examples/openapi/internal/api"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
)

type UsersHandler struct {
	DB *sql.DB
}

func (u *UsersHandler) GetUser(ctx echo.Context) error {
	// load user from database and return it to the caller
	return ctx.JSON(http.StatusOK, api.User{
		Email: types.Email("demo@devopsian.net"),
		Name:  "DemoUser",
		Id:    "1",
	})
}

func (u *UsersHandler) PostSignup(ctx echo.Context) error {
	var body api.PostSignupJSONBody
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.Error{Code: http.StatusBadRequest, Message: "invalid request"})
	}
	// save user in database

	return ctx.NoContent(http.StatusOK)
}

func New(db *sql.DB) *UsersHandler {
	return &UsersHandler{DB: db}
}
