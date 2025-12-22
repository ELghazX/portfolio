package database

// example with postgres
import (
	"database/sql"

	"github.com/elghazx/portfolio/internal/database/gen"
	// _ "github.com/lib/pq"
)

var Queries *gen.Queries

func Init(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	Queries = gen.New(db)
	return nil
}
