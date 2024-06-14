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
	adminPanelGroup.Get("/getobj", adminPanelController.GetObj)
	adminPanelGroup.Get("/getlist", adminPanelController.GetList)
	//Todo: добавить setobj - что бы изменять обьект в админке
	//Todo: добавить getTables - отдавать список таблиц

	return adminPanelServer.ServeHTTP(cfg.BindAddr)
}
