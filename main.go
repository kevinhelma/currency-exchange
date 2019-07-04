package main

import (
	"currency-exchange/shared"
	"currency-exchange/shared/config"
	"currency-exchange/shared/data"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"net/http"
)

var (
	runMigration  bool
	runSeeder     bool
	db            *gorm.DB
	configuration config.Configuration
	router        *gin.Engine
)

func init() {
	flag.BoolVar(&runMigration, "migrate", false, "run data migration")
	flag.BoolVar(&runSeeder, "seed", false, "run data seeder after data migration")
	flag.Parse()

	glog.V(2).Info("Migration status : %s", runMigration)
	glog.V(2).Info("Initilizing server...")

	//Setup Configuration
	cfg, err := config.New("")
	if err != nil {
		glog.Fatalf("Failed to load configuration: %s", err)
		panic(fmt.Errorf("Fatal error on load configuration : %s ", err))
	}
	configuration = *cfg

	//Setup Router
	routerInstance := shared.NewRouter(configuration)
	router = routerInstance.SetupRouter()

	if runMigration == true {
		dbMigration, err := data.NewDbMigration(&configuration)
		if err != nil {
			glog.Fatalf("Failed instantiate dbmigration : %s", err)
		}

		var count = 0

		if migrate(dbMigration, runSeeder) != nil {
			if count == 10 {
				return
			}
			count++
			migrate(dbMigration, runSeeder)
		}
	}
}

func migrate(db *data.DbMigration, seed bool) error {
	success, err := db.Migrate(seed)

	if err != nil {
		glog.Fatalf("Failed migrate: %s", err)
	}

	glog.V(2).Infof("database migration : %s", success)
	return err
}

func main() {
	glog.V(2).Infof("Server run on mode: %s", configuration.Server.Mode)
	gin.SetMode(configuration.Server.Mode)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(fmt.Errorf("Fatal error failed to start server : %s", err))
	}
}
