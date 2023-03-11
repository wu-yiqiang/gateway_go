package global

import (
  "github.com/spf13/viper"
  "gateway_go/config"
  "go.uber.org/zap"
)

type Application struct {
  ConfigViper *viper.Viper
  Config config.Configuration
  Log *zap.Logger
}

var App = new(Application)