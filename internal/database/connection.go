package database

// example with postgres
import (
	"database/sql"

	"github.com/elghazx/portfolio/internal/database/generated"
	// _ "github.com/lib/pq"
)

var Queries *generated.Queries

func Init(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	Queries = generated.New(db)
	return nil
}
