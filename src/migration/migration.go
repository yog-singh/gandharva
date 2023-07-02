package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/yog-singh/gandharva/src/config"
)

func RunMigration() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)
	db, _ := sql.Open("postgres", dsn)
	driver, _ := postgres.WithInstance(db, &postgres.Config{})

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	migrationsPath := fmt.Sprintf("file://%s/conf/migrations", path)
	fmt.Println(migrationsPath)

	migration, _ := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	migration.Up()
}
