package driver

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/murasame29/go-httpserver-template/cmd/config"
	"github.com/pandoratoolbox/bun/extra/bunslog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func NewDB() *sql.DB {
	var sqlDB *sql.DB
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Config.DB.User,
		config.Config.DB.Password,
		config.Config.DB.Host,
		config.Config.DB.Port,
		config.Config.DB.DBName,
		config.Config.DB.SSLMode,
	)
	slog.Debug("connecting database.", "dsn", dsn)

	pgxConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		slog.Error("failed to parse config.", "error", err)
		panic(err)
	}

	sqlDB = stdlib.OpenDB(*pgxConfig)

	return sqlDB
}

func NewBun(sqlDB *sql.DB) *bun.DB {
	var err error
	for i := range 5 {
		slog.Info("trying open DB.", "attempts", i+1)
		db := bun.NewDB(sqlDB, pgdialect.New())
		if err = db.Ping(); err == nil {

			db.AddQueryHook(bunslog.NewQueryHook(
				bunslog.WithQueryLogLevel(slog.LevelDebug),
				bunslog.WithSlowQueryLogLevel(slog.LevelWarn),
				bunslog.WithErrorQueryLogLevel(slog.LevelError),
				bunslog.WithSlowQueryThreshold(3*time.Second),
			))

			return db
		}
	}

	slog.Error("failed to open DB.", "error", err)
	return nil
}
