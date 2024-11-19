package mysql

import (
	"context"
	"github.com/gocraft/dbr/v2"
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
	ID       int    `db:id`
	UUID     string `db:uuid`
	Name     string `db:name`
	Email    string `db:email`
	Password string `db:password`
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
