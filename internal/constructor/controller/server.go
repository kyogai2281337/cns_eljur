package controller

import (
	"database/sql"
	"log"

	"github.com/kyogai2281337/cns_eljur/internal/constructor/service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
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

	constructorController := service.NewConstructorController(newServer, cfg.MongoURL)

	controllerGroup := newServer.App.Group("/private/schedule")
	controllerGroup.Use(constructorController.Authentication())
	controllerGroup.Post("/find", constructorController.Find)
	// controllerGroup.Post("/add", constructorController.Add)
	// controllerGroup.Post("/delete", constructorController.Delete)
	// controllerGroup.Post("/update", constructorController.Update)
	// TODO implement
	// controllerGroup.Get("/export", constructorController.Export)

	// controllerGroup.Get("/import", constructorController.Import)

	// controllerGroup.Get("/create", constructorController.Create)
	// controllerGroup.Get("/groups", constructorController.GetAllGroups)
	// controllerGroup.Get("/cabinets", constructorController.GetAllCabinets)
	// controllerGroup.Get("/teachers", constructorController.GetAllTeachers)
	// controllerGroup.Get("/subjects", constructorController.GetAllSubjects)
	return newServer.ServeHTTP(cfg.BindAddr)
}
