package controller

import (
	"database/sql"
	"log"

	"github.com/kyogai2281337/cns_eljur/internal/constructor/service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
)

// Start initializes the server and starts the controller.
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
	serverController := service.NewConstructorController(newServer)
	constructorGroup := newServer.App.Group("/private/constructor")

	constructorGroup.Post("/create", serverController.Create)

	constructorGroup.Post("/get", serverController.Get)

	constructorGroup.Get("/getlist", serverController.GetList)
	return newServer.ServeHTTP(cfg.BindAddr)
}
