package adminPanel_controller

import (
	"database/sql"
	"github.com/kyogai2281337/cns_eljur/internal/adminPanel/adminPanel_service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
	"log"
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
	adminPanelController := adminPanel_service.NewAdminPanelController(adminPanelServer)

	adminPanelGroup := adminPanelServer.App.Group("/admin")
	adminPanelGroup.Use(adminPanelController.Authentication())
	adminPanelGroup.Get("/profile", adminPanelController.User)
	adminPanelGroup.Get("/logout", adminPanelController.Logout)
	adminPanelGroup.Get("/delete", adminPanelController.Delete)

	return adminPanelServer.ServeHTTP(cfg.BindAddr)
}
