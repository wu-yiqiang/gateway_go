package global

import (
  "github.com/spf13/viper"
  "gateway_go/config"
  "go.uber.org/zap"
  "gorm.io/gorm"
)

type Application struct {
  ConfigViper *viper.Viper
  Config config.Configuration
  Log *zap.Logger
  DB *gorm.DB
}

var App = new(Application)