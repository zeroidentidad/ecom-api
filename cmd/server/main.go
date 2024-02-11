package main

import (
	"log"
	"os"

	"github.com/zeroidentidad/ecom-api/internal/db" // moduleName/internal/pkgName
	"github.com/zeroidentidad/ecom-api/internal/ecommerce"
	transportHttp "github.com/zeroidentidad/ecom-api/internal/transport/http" // moduleName/internal/pkgName
)

func main() {
	if err := Run(); err != nil {
		log.Println(err.Error())
	}
}

// simulate defined ENVs on GAE or AWS instance
func init() {
	os.Setenv("DB_HOST", "postgresql-zeroidentidad.alwaysdata.net")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "zeroidentidad_codifin")
	os.Setenv("DB_NAME", "zeroidentidad_test")
	os.Setenv("DB_PASSWORD", "codifin.xd")
	os.Setenv("DB_SSLMODE", "disable")
}

func Run() error {
	log.Println("starting up our api")

	db, err := db.NewDatabase()
	if err != nil {
		log.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		log.Println("Failed to migrate the database")
		return err
	}

	ecommerceService := ecommerce.NewService(db)
	httpHandler := transportHttp.NewHandler(ecommerceService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}
