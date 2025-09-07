package bootstrap

import (
	"context"
	"log"

	"github.com/AlfredoPrograma/diployable/db"
	"github.com/jackc/pgx/v5"
)

// ConnectDB establishes a connection to the database using the provided connection string.
// It returns a pointer to db.Queries for executing database operations.
// If the connection fails, it logs the error and terminates the application.
func ConnectDB(connString string) *db.Queries {
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Printf("cannot connect to database using connection string: %s\n", connString)
		log.Fatalln(err)
	}

	return db.New(conn)
}
