migrations/up:
	sql-migrate up -config="./config/db.config.yml" -env="development"

gen/wire:
	wire gen ./infra/wire/wire.go

# make gen/migration FILE_NAME=
gen/migration:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d%H%M%S"))
	touch ./migrations/$(DATE)-$(FILE_NAME).sql

golint:
	golangci-lint run --fix --config "./config/.golangci.yml"