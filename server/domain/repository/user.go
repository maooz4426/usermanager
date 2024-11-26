package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/maooz4426/usermanager/domain/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, entity *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByUUID(ctx context.Context, uuidReq uuid.UUID) (*entity.User, error)
}
