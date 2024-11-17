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

	// @Summary Get object
	// @Description Retrieve a specific object by its ID
	// @Tags admin
	// @Accept json
	// @Produce json
	// @Param body body structures.GetObjRequest true "Request body with object ID"
	// @Success 200 {object} structures.GetObjResponse
	// @Failure 400 {object} ErrorResponse
	// @Router /private/getobj [post]
	adminPanelGroup.Post("/getobj", adminPanelController.GetObj)

	// @Summary Get list of objects
	// @Description Retrieve a paginated list of objects from a specific table
	// @Tags admin
	// @Accept json
	// @Produce json
	// @Param body body structures.GetListRequest true "Request body with pagination"
	// @Success 200 {object} structures.GetListResponse
	// @Failure 400 {object} ErrorResponse
	// @Router /private/getlist [post]
	adminPanelGroup.Post("/getlist", adminPanelController.GetList)

	// @Summary Set object
	// @Description Update or create an object in a specific table
	// @Tags admin
	// @Accept json
	// @Produce json
	// @Param body body structures.SetObjRequest true "Request body with object data"
	// @Success 200 {object} structures.SetObjResponse
	// @Failure 400 {object} ErrorResponse
	// @Router /private/setobj [post]
	adminPanelGroup.Post("/setobj", adminPanelController.SetObj)

	// @Summary Get tables
	// @Description Retrieve a list of available tables in the admin panel
	// @Tags admin
	// @Produce json
	// @Success 200 {object} structures.GetTablesResponse
	// @Failure 400 {object} ErrorResponse
	// @Router /private/gettables [get]
	adminPanelGroup.Get("/gettables", adminPanelController.GetTables)

	// @Summary Create object
	// @Description Create a new object in a specific table
	// @Tags admin
	// @Accept json
	// @Produce json
	// @Param body body structures.CreateRequest true "Request body with object data"
	// @Success 201 {object} structures.CreateResponse
	// @Failure 400 {object} ErrorResponse
	// @Router /private/create [post]
	adminPanelGroup.Post("/create", adminPanelController.Create)

	return adminPanelServer.ServeHTTP(cfg.BindAddr)
}
