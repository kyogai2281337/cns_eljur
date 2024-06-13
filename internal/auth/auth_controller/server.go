package auth_controller

import (
	"database/sql"
	"github.com/kyogai2281337/cns_eljur/internal/auth/auth_service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
	"log"
)

// Start initializes the server and starts the authentication auth_controller.
func Start(cfg *server.Config) error {
	db, err := server.NewDB(cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error: %s", err)
		}
	}(db)

	store := sqlstore.New(db)
	newServer := server.NewServer(store)
	authController := auth_service.NewAuthController(newServer)

	newServer.App.Use(authController.RequestID())
	newServer.App.Use(authController.Log())

	newServer.App.Post("/signup", authController.Register)
	newServer.App.Post("/signin", authController.Login)

	privateGroup := newServer.App.Group("/private")
	privateGroup.Use(authController.Authentication())
	privateGroup.Get("/profile", authController.User)
	privateGroup.Get("/logout", authController.Logout)
	privateGroup.Get("/delete", authController.Delete)

	return newServer.ServeHTTP(cfg.BindAddr)
}
