package port

import (
	"context"
	"github.com/maooz4426/usermanager/domain/entity"

	"github.com/google/uuid"
	"github.com/maooz4426/usermanager/lib/usermanage"
)

type IUserUsecase interface {
	Signup(ctx context.Context, user *usermanage.UserSignupRequest) (uuid.UUID, error)
	SignIn(ctx context.Context, req *usermanage.UserSignInRequest) (uuid.UUID, error)
	Get(ctx context.Context, uuid uuid.UUID) (*entity.User, error)
}
