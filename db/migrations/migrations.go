package migrations

import (
	"github.com/go-pg/migrations/v8"
	"github.com/labstack/gommon/log"
	"ml-mutant-test/config"
	database "ml-mutant-test/db"
)

func execute(cmd string) error {
	var err error
	log.Info("Executing migrations")
	dbInstance := database.ConnInstance()

	createMigrationTable := `CREATE TABLE IF NOT EXISTS gopg_migrations (
		id SERIAL,
		version int,
		created_at timestamp
	);`
	_, err = dbInstance.Exec(createMigrationTable)
	if err != nil {
		return err
	}

	oldVersion, newVersion, err := migrations.Run(dbInstance, cmd)

	if err != nil {
		return err
	}

	if newVersion != oldVersion {
		log.Infof("Migrated from version %d to %d", oldVersion, newVersion)
	} else {
		log.Infof("Version is %d", oldVersion)
	}

	return nil
}

func Run() error {
	return execute("up")
}

func Rollback() error {
	return execute("down")
}

func Reset() error {
	return execute("reset")
}

func StartConfiguration() {
	switch config.Environments().MigrationsCommand {
	case "run":
		if err := Run(); err != nil {
			log.Errorf("Unable to run migrations due to %#v", err)
		}
	case "rollback":
		if err := Rollback(); err != nil {
			log.Errorf("Unable to rollback migrations due to %#v", err)
		}
	case "reset":
		if err := Reset(); err != nil {
			log.Errorf("Unable to reset migrations due to %#v", err)
		}
	}
}
