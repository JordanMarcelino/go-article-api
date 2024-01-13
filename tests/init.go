package tests

import (
	"github.com/jordanmarcelino/go-article-api/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var viperConfig *viper.Viper

var log *logrus.Logger

func init() {
	viperConfig = config.NewViper()
	log = config.NewLogger(viperConfig)
}
