package controller

import (
	"github.com/kyogai2281337/cns_eljur/internal/auth/service"
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
	adminpanelGroup.Get("/profile", controller.User)
	adminpanelGroup.Get("/logout", controller.Logout)
	adminpanelGroup.Get("/delete", controller.Delete)

	return server.ServeHTTP(cfg.BindAddr)
