package ldb

import (
	"database/sql"
	"log"
	"os"

	"github.com/deeper-x/krowki/conf"
	"github.com/gchaincl/dotsql"
	"github.com/joho/godotenv"

	// blank import
	_ "github.com/lib/pq"
)

// ResultSet query container
type ResultSet interface {
	AllMoored(idPortinformer string) []map[string]string
}

type connector struct {
	db *sql.DB
}

type dotMapper struct {
	resource *dotsql.DotSql
}

// NewConnection is the connection starter
func NewConnection(db *sql.DB) ResultSet {
	return connector{db: db}
}

// Connect run dsn from env file and start
func Connect() (*sql.DB, error) {
	err := godotenv.Load(conf.EnvFile)

	if err != nil {
		log.Fatal(err)
	}

	dbdsn := os.Getenv("DB_DSN")

	db, err := sql.Open("postgres", dbdsn)

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}