package auth_controller

import (
	"github.com/kyogai2281337/cns_eljur/internal/auth/auth_service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
)

// Start initializes the server and starts the authentication auth_controller.
func Start(cfg *server.Config) error {
	db, err := server.NewDB(cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	server := server.NewServer(store)
	controller := auth_service.NewAuthController(server)

	server.App.Use(controller.RequestID())
	server.App.Use(controller.Log())

	server.App.Post("/signup", controller.Register)
	server.App.Post("/signin", controller.Login)

	privateGroup := server.App.Group("/private")
	privateGroup.Use(controller.Authentication())
	privateGroup.Get("/profile", controller.User)
	privateGroup.Get("/logout", controller.Logout)
	privateGroup.Get("/delete", controller.Delete)

	return server.ServeHTTP(cfg.BindAddr)
}
