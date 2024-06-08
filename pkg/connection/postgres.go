package connection

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/ssofiica/test-task-gazprom/config"
)

const maxPingCount = 10

func InitPostgres(cfg *config.Project) *sql.DB {
	dbConfig := cfg.Postgres

	dataConnection := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.User, dbConfig.Password, dbConfig.SslMode)

	db, err := sql.Open("postgres", dataConnection)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to connect to PostgreSQL %s at address %s:%d\n: ",
			cfg.Postgres.Database, cfg.Postgres.Host, cfg.Postgres.Port)
		log.Fatal(errorMsg, err.Error())
	}

	// maximum number of open connections
	db.SetMaxOpenConns(int(dbConfig.MaxOpenConns))
	// maximum amount of time the connection can be reused
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Second)
	// maximum number of connections in the pool of idle connections
	db.SetMaxIdleConns(int(dbConfig.MaxIdleConns))
	// maximum time during which the connection can be idle
	db.SetConnMaxIdleTime(time.Duration(dbConfig.ConnMaxIdleTime) * time.Second)

	currentStep := 0
	for i := 0; i < maxPingCount; i++ {
		err = db.Ping()
		currentStep++
		if err == nil {
			break
		}
		if currentStep == maxPingCount {
			log.Fatal("Unable to connect to PostgreSQL: ", err.Error())
		}

	}

	fmt.Println("Postgres connected successfully")
	return db
}
