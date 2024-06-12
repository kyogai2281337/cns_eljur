package controller

import (
	"github.com/kyogai2281337/cns_eljur/internal/auth/service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
)

func Start(config *server.Config) error {
	db, err := server.NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	s := server.NewServer(store)
	controller := service.NewAuthController(s)
	s.App.Use(controller.Log())
	s.App.Use(controller.RequestID())

	s.App.Post("/signup", controller.Register)
	s.App.Post("/signin", controller.Login)

	private := s.App.Group("/private")
	private.Use(controller.Authentication())
	private.Get("/profile", controller.User)
	private.Get("/logout", controller.Logout)
	private.Get("/delete", controller.Delete)
	return s.ServeHTTP(config.BindAddr)
}
