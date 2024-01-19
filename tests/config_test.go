package tests

import (
	"fmt"
	"github.com/jordanmarcelino/go-article-api/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViperConfig(t *testing.T) {
	appName := viperConfig.GetString("app.name")
	appVersion := viperConfig.GetString("app.version")

	webPrefork := viperConfig.GetBool("web.prefork")
	webPort := viperConfig.GetInt("web.port")

	logLevel := viperConfig.GetInt("log.level")

	dbUser := viperConfig.GetString("database.user")
	dbPassword := viperConfig.GetString("database.password")
	dbHost := viperConfig.GetString("database.host")
	dbName := viperConfig.GetString("database.name")
	dbPort := viperConfig.GetInt("database.port")
	dbPoolIdle := viperConfig.GetInt("database.pool.idle")
	dbPoolMax := viperConfig.GetInt("database.pool.max")
	dbPoolLifetime := viperConfig.GetInt("database.pool.lifetime")

	assert.EqualValues(t, "golang-article-api", appName)
	assert.EqualValues(t, "1.0.0", appVersion)
	assert.True(t, webPrefork)
	assert.EqualValues(t, 5050, webPort)
	assert.EqualValues(t, 5, logLevel)
	assert.EqualValues(t, "root", dbUser)
	assert.EqualValues(t, "", dbPassword)
	assert.EqualValues(t, "localhost", dbHost)
	assert.EqualValues(t, "golang_article_api", dbName)
	assert.EqualValues(t, 3306, dbPort)
	assert.EqualValues(t, 10, dbPoolIdle)
	assert.EqualValues(t, 50, dbPoolMax)
	assert.EqualValues(t, 600, dbPoolLifetime)
}

func TestLogrus(t *testing.T) {
	assert.NotNil(t, log)

	log.Info("Testing log")
}

func TestHashing(t *testing.T) {
	password := "kebab123"
	hashPassword, _ := util.HashPassword(password)
	fmt.Println(hashPassword)

	fmt.Println(util.CheckPasswordHash(password, hashPassword))
	fmt.Println(util.CheckPasswordHash(hashPassword, hashPassword))
}
