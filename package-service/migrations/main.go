package main

import (
	"github.com/ancore09/package-service/internal/config"
	"github.com/ancore09/package-service/internal/pkg/db"
	"github.com/rs/zerolog/log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	if err := config.ReadConfigYML("../config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	conn, err := db.ConnectDb(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("sql open")
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("ping failed")
	}

	query := sq.Insert("packages").Columns("name").PlaceholderFormat(sq.Dollar).Values("lib3").RunWith(conn)
	_, err = query.Exec()
	if err != nil {
		log.Fatal().Err(err).Msg("insert failed")
	}
}
