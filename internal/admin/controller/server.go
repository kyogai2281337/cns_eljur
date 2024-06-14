package controller

import (
	"database/sql"
	"log"

	"github.com/kyogai2281337/cns_eljur/internal/admin/service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
)

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
	adminPanelServer := server.NewServer(store)
	adminPanelController := service.NewAdminPanelController(adminPanelServer)

	adminPanelGroup := adminPanelServer.App.Group("/private/admin")
	adminPanelGroup.Use(adminPanelController.Authentication())
	adminPanelGroup.Get("/GetObj", adminPanelController.GetObj)
	adminPanelGroup.Get("/GetList", adminPanelController.GetList)

	return adminPanelServer.ServeHTTP(cfg.BindAddr)
}
