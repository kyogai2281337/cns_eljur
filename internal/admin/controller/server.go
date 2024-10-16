package controller

import (
	"log"

	"github.com/kyogai2281337/cns_eljur/internal/admin/service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
)

func Start(cfg *server.Config) error {
	db, err := server.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Printf("Failed to establish a DB connection: %s", err)
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Failed to close DB connection: %s", err)
		}
	}()

	store := sqlstore.New(db)
	adminPanelServer := server.NewServer(store)
	adminPanelController := service.NewAdminPanelController(adminPanelServer)

	adminPanelGroup := adminPanelServer.App.Group("/private")
	adminPanelGroup.Use(adminPanelController.Authentication())
	adminPanelGroup.Post("/getobj", adminPanelController.GetObj)
	adminPanelGroup.Post("/getlist", adminPanelController.GetList)
	adminPanelGroup.Post("/setobj", adminPanelController.SetObj)
	adminPanelGroup.Get("/gettables", adminPanelController.GetTables)

	adminPanelGroup.Post("/create", adminPanelController.Create)

	return adminPanelServer.ServeHTTP(cfg.BindAddr)
}
