package mysql

import (
	"context"
	"log"

	"github.com/gocraft/dbr/v2"
	"github.com/google/uuid"
	"github.com/maooz4426/usermanager/domain/entity"
	"github.com/maooz4426/usermanager/domain/repository"
)

const (
	TableName = "users"
)

type UserRepository struct {
	sess *dbr.Session
}

func NewUserRepository(sess *dbr.Session) repository.IUserRepository {
	return &UserRepository{
		sess: sess,
	}
}

type schemaUsers struct {
	ID       int    `db:"id"`
	UUID     string `db:"uuid"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func entityToSchemaUsers(entity *entity.User) *schemaUsers {
	return &schemaUsers{
		UUID:     entity.UUID.String(),
		Name:     entity.Name,
		Email:    entity.Email,
		Password: entity.Password,
	}
}

func (r *UserRepository) Create(ctx context.Context, entity *entity.User) error {
	schemaUser := entityToSchemaUsers(entity)

	if _, err := r.sess.InsertInto(TableName).Columns("uuid", "name", "email", "password").Record(schemaUser).ExecContext(ctx); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var schemaUser schemaUsers
	_, err := r.sess.Select("uuid", "name", "email", "password").From(TableName).Where("email = ?", email).LoadContext(ctx, &schemaUser)
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(schemaUser.UUID)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		UUID:     uuid,
		Password: schemaUser.Password,
	}
	return user, err
}

func (r *UserRepository) FindByUUID(ctx context.Context, uuidReq uuid.UUID) (*entity.User, error) {

	var schemaUser schemaUsers
	if _, err := r.sess.Select("uuid", "name", "email").From(TableName).Where("uuid = ?", uuidReq).LoadContext(ctx, &schemaUser); err != nil {
		return nil, err
	}

	user := &entity.User{
		UUID:  uuidReq,
		Name:  schemaUser.Name,
		Email: schemaUser.Email,
	}

	log.Print(user)

	return user, nil
}
