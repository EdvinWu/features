package app

import (
	"features/config"
	"features/internal/db/postgres"
	"features/internal/domain/feature/handler"
	"features/internal/domain/feature/repository"
	"features/internal/domain/feature/service"
	"features/internal/server"
	"features/internal/util"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
)

type App struct {
	conf   *config.Config
	logger *logrus.Entry
}

func NewApp(conf *config.Config, logger *logrus.Entry) App {
	return App{conf: conf, logger: logger}
}

func (a *App) Run() {
	db := a.setupPostgres()
	repo := repository.NewFeature(db)
	serv := service.NewFeature(repo, a.logger)
	feature := handler.NewFeature(serv)
	addr := fmt.Sprintf(":%d", a.conf.Server.Port)
	a.logger.Infof("Starting HTTP server on [%s]", addr)
	echo := server.Echo(addr, feature)
	if err := graceful.ListenAndServe(echo.Server, 0); err != nil {
		a.logger.WithError(err).Fatal("Failed to start server")
	}
}

func (a App) setupPostgres() *sqlx.DB {
	db, err := postgres.Connect(&a.conf.Postgres, postgres.URL(&a.conf.Postgres))
	util.PanicIfError(err, "failed to connect to db")

	err = postgres.Migrate(db, a.conf.Postgres.MigrationPath, a.conf.Postgres.Database)
	util.PanicIfError(err, "failed to run postgres migration")

	return db
}
