package handler

import (
	"github.com/maooz4426/usermanager/domain/port"
	"github.com/maooz4426/usermanager/lib/usermanage"
)

type Handler struct {
	UserUsecase port.IUserUsecase
}

func NewHandler(
	userUsecase port.IUserUsecase,
) usermanage.ServerInterface {
	return &Handler{
		UserUsecase: userUsecase,
	}
}
