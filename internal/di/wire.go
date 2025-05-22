//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"go-chi-wire/internal/handler"
	"go-chi-wire/internal/router"
	"go-chi-wire/internal/service"
	"net/http"
)

func InitializeRouter() http.Handler {
	wire.Build(
		service.NewUserService,
		handler.NewUserHandler,
		router.NewUserRouter,
	)

	return nil
}
