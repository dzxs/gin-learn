package logging

import (
	"fmt"
	"github.com/dzxs/blog/pkg/setting"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(TimeFormat), setting.AppSetting.LogFileExt)
}