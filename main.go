package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/andreyskoskin/drvolodko/domainimpl/users"
	"github.com/andreyskoskin/drvolodko/http"
)

// memory:
type memDomain struct {
	users.MemoryUsers
}

// pg:
/*
type pgDomain struct {
	*users.PgUsers
}
*/

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() (err error) {
	var cfg *Config
	if cfg, err = LoadConfig("config.toml"); err != nil {
		return err
	}

	// use mem
	var e = http.Routes(&memDomain{}, log.Println)

	// use postgres
	/*
		var connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)
		var db *sql.DB
		if db, err = sql.Open("postgres", connStr); err != nil {
			return err
		}
		var pg = pgDomain{
			PgUsers: users.NewPgUsers(db),
		}
		var e = http.Routes(pg, log.Println)
	*/

	return e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.HTTP.Port))
}
