package global

import (
	"github.com/tianmai777/blog/pkg"
	"github.com/tianmai777/blog/pkg/logger"
)

var (
	ServerSetting   *pkg.ServerSettingS
	AppSetting      *pkg.AppSettingS
	DatabaseSetting *pkg.DatabaseSettingS
	Log             *logger.Logger
)