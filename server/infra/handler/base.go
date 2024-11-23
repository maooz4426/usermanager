package handler

import (
	"github.com/maooz4426/usermanager/lib/usermanage"
	"github.com/maooz4426/usermanager/usecases"
)

type Handler struct {
	UserUsecase *usecases.UserUsecase
}

func NewHandler(
	userUsecase *usecases.UserUsecase,
) usermanage.ServerInterface {
	return &Handler{
		UserUsecase: userUsecase,
	}
}
