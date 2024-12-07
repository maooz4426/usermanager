package usecases

import (
	"context"

	"github.com/google/uuid"

	"github.com/maooz4426/usermanager/domain/entity"
	"github.com/maooz4426/usermanager/domain/repository"
	"github.com/maooz4426/usermanager/lib/usermanage"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
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

func (s *UserUsecase) SignIn(ctx context.Context, req *usermanage.UserSignInRequest) (uuid.UUID, error) {
	user, err := s.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return uuid.UUID{}, err
	}

	return user.UUID, nil
}

func (s *UserUsecase) Get(ctx context.Context, uuid uuid.UUID) (*entity.User, error) {
	user, err := s.userRepository.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
