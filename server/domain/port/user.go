package port

import (
	"context"
	"github.com/google/uuid"
	"github.com/maooz4426/usermanager/lib/usermanage"
)

type IUserUsecase interface {
	Signup(ctx context.Context, user *usermanage.UserSignupRequest) (uuid.UUID, error)
}
