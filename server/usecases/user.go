package usecases

import (
	"context"
	"github.com/google/uuid"
	"github.com/maooz4426/usermanager/domain/entity"
	"github.com/maooz4426/usermanager/domain/port"
	"github.com/maooz4426/usermanager/domain/repository"
	"github.com/maooz4426/usermanager/lib/usermanage"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) port.IUserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (s *UserUsecase) Signup(ctx context.Context, req *usermanage.UserSignupRequest) (uuid.UUID, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.UUID{}, err
	}

	user := entity.User{
		UUID:     uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashPassword),
	}

	err = s.userRepository.Create(ctx, &user)
	if err != nil {
		return uuid.UUID{}, err
	}
	return user.UUID, nil
}
