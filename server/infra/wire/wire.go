//go:build wireinject

package wire

import (
	"github.com/gocraft/dbr/v2"
	"github.com/google/wire"
	"github.com/maooz4426/usermanager/infra/handler"
	"github.com/maooz4426/usermanager/infra/mysql"
	"github.com/maooz4426/usermanager/lib/usermanage"
	"github.com/maooz4426/usermanager/usecases"
)

func Wire(sess *dbr.Session) usermanage.ServerInterface {
	wire.Build(
		mysql.NewUserRepository,
		usecases.NewUserUsecase,
		handler.NewHandler,
	)

	return nil
}
