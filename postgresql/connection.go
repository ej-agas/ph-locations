package postgresql

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

func NewConnection(config Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Password, config.DatabaseName)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, fmt.Errorf("error connecting to postgresql: %s", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("error pinging postgresql: %s", err)
	}

	return db, nil
}
