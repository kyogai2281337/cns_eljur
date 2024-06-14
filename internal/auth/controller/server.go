package controller

import (
	"database/sql"
<<<<<<< HEAD
	"log"

=======
>>>>>>> 856b863359fac258d673aad6420eef74e1e0c7bb
	"github.com/kyogai2281337/cns_eljur/internal/auth/service"
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
	authController := service.NewAuthController(newServer)

	newServer.App.Use(authController.RequestID())
	newServer.App.Use(authController.Log())

	newServer.App.Post("/signup", authController.Register)
	newServer.App.Post("/signin", authController.Login)

	newServer.App.Get("/private/profile", authController.User)
	newServer.App.Get("/private/logout", authController.Logout)
	newServer.App.Get("/private/delete", authController.Delete)

	return newServer.ServeHTTP(cfg.BindAddr)
}
