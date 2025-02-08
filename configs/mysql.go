package configs

import (
	"database/sql"

	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySql(databaseDSN string) *sql.DB {
	log.Infoln("Connecting to Database")

	db, err := sql.Open("mysql", databaseDSN)
	if err != nil || db.Ping() != nil {
		for i := 1; i <= 20; i++ {
			log.Infof("Retrying Database Connection #%d\n", i)
			db, err = sql.Open("mysql", databaseDSN)
			if err == nil && db.Ping() == nil {
				break
			}
			time.Sleep(5 * time.Second)
		}

		if err != nil || db.Ping() != nil {
			panic(db.Ping())
		}
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	log.Infoln("Database Connected!")
	return db
}
