package dbversion

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"go.uber.org/zap"
)

// WaitUntilLatestVersion blocks until the database is migrated to the latest version.
// This function can exit the program when failing to query postgres for the version.
func WaitUntilLatestVersion(logger *zap.Logger, db *sql.DB) {
	dbMigrateDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatal("could not open driver to postgres connection", zap.Error(err))
	}
	for {
		version, dirty, err := dbMigrateDriver.Version()
		if err != nil && err != migrate.ErrNilVersion {
			logger.Fatal("could not obtain db version", zap.Error(err))
		}
		if dirty {
			logger.Info(fmt.Sprintf("db migration is dirty at version %d, waiting...", version))
			time.Sleep(1 * time.Second)
			continue
		}
		if version != LatestVersion {
			logger.Info(fmt.Sprintf("db is at version %d, require version %d", version, LatestVersion))
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
}
