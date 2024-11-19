package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/usermanager/lib/usermanage"
	"net/http"
)

// (POST /users/signup)
func (h *Handler) UserSignUp(ctx echo.Context) error {
	var req usermanage.UserSignupRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	//user := entity.User{
	//	Name:     req.Name,
	//	Email:    req.Email,
	//	Password: req.Password,
	//}

	uuid, err := h.UserUsecase.Signup(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"uuid": uuid})
}

// (POST /users/signup)
func (h *Handler) UserSignIn(ctx echo.Context) error {
	return nil
}

// (GET /users/{uuid})
func (h *Handler) GetUser(ctx echo.Context, uuid usermanage.UserUUID) error {
	return nil
}
