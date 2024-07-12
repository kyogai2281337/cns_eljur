package controller

import (
	"database/sql"
	"log"

	"github.com/kyogai2281337/cns_eljur/internal/constructor/service"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"

	mongoDB "github.com/kyogai2281337/cns_eljur/pkg/mongo"
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

	mongoClient, ctx, cancel := mongoDB.ConnectMongoDB(cfg.MongoURL)
	defer cancel()
	defer mongoClient.Disconnect(ctx)

	// mdb := mongoClient.Database("school")
	// cabinetsCollection := mdb.Collection("cabinets")
	// subjectsCollection := mdb.Collection("subjects")
	// specializationsCollection := mdb.Collection("specializations")
	// groupsCollection := mdb.Collection("groups")
	// teachersCollection := mdb.Collection("teachers")
	// lecturesCollection := mdb.Collection("lectures")

	store := sqlstore.New(db)
	newServer := server.NewServer(store)

	constructorController := service.NewConstructorController(newServer, mongoClient)

	controllerGroup := newServer.App.Group("/private/schedule")
	controllerGroup.Use(constructorController.Authentication())
	controllerGroup.Post("/find", constructorController.Find)

	return newServer.ServeHTTP(cfg.BindAddr)
}
