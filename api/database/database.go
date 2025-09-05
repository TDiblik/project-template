package database

import (
	"database/sql"
	"sync"
	"time"

	"github.com/TDiblik/project-template/api/utils"
	_ "github.com/jackc/pgx/v5/stdlib" // Register pgx with database/sql
	"github.com/jmoiron/sqlx"
)

var (
	db    *sqlx.DB
	once  sync.Once
	dbErr error
)

// CreateConnection initializes or returns the global DB pool.
func CreateConnection() (*sqlx.DB, error) {
	once.Do(func() {
		d, err := sqlx.Connect("pgx", utils.EnvData.DB_CONNECTION_STRING)
		if err != nil {
			dbErr = err
			return
		}

		// Pool tuning (adjust based on workload & DB limits)
		d.SetMaxOpenConns(50)                  // Max total connections
		d.SetMaxIdleConns(10)                  // Keep some idle for reuse
		d.SetConnMaxIdleTime(5 * time.Minute)  // Kill idle conns after 5 min
		d.SetConnMaxLifetime(30 * time.Minute) // Recycle conns after 30 min

		db = d
	})

	return db, dbErr
}

func ExecuteTransaction(db *sqlx.DB, fn func(*sql.Tx) error) (err error) {
	utils.Log("starting a transaction")
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if utils.EnvData.Debug {
				utils.Log("transaction panic, rolling back: ", p)
			}
			if rbErr := tx.Rollback(); rbErr != nil {
				utils.Log("failed to rollback after panic: ", rbErr)
			}
			panic(p) // rethrow panic after rollback
		} else if err != nil {
			utils.Log("transaction failed, rolling back: ", err)
			if rbErr := tx.Rollback(); rbErr != nil {
				utils.Log("failed to rollback after error: ", rbErr)
			}
		} else {
			utils.Log("committing transaction")
			if cmErr := tx.Commit(); cmErr != nil {
				utils.Log("failed to commit transaction: ", cmErr)
				err = cmErr
			}
		}
	}()

	err = fn(tx)
	return err
}
