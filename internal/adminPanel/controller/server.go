package controller

import (
	"github.com/kyogai2281337/cns_eljur/internal/adminPanel/service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
)

func Start(cfg *server.Config) error {
	db, err := server.NewDB(cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	server := server.NewServer(store)
	controller := service.NewAdminPanelController(server)

	adminpanelGroup := server.App.Group("/admin")
	adminpanelGroup.Use(controller.Authentication())
	adminpanelGroup.Get("/GetObj", controller.User)
	adminpanelGroup.Get("/GetList", controller.User)

	return server.ServeHTTP(cfg.BindAddr)
