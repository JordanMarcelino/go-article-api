package config

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func MigrateDB(viper *viper.Viper, logger *logrus.Logger) {
	username := GetEnv("POSTGRES_USER", viper.GetString("database.user"))
	password := GetEnv("POSTGRES_PASSWORD", viper.GetString("database.password"))
	host := GetEnv("DB_HOST", viper.GetString("database.host"))
	port := viper.GetInt("database.port")
	database := viper.GetString("database.name")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Jakarta", username, password, host, port, database)

	migrations, err := migrate.New("file://db/migrations", dsn)
	if err != nil {
		logger.Fatalf("Failed to create migrate instance : %+v", err)
	}

	if err := migrations.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatalf("Failed to migrate database : %+v", err)
	}

	logger.Info("Success migrate database")
}
