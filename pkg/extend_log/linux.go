//go:build linux

package extend_log

import (
	"log/syslog"

	"github.com/allanpk716/ChineseSubFinder/pkg/settings"
	"github.com/sirupsen/logrus"
)

type ExtendLog struct {
}

func (e *ExtendLog) AddHook(log *logrus.Logger, extendLog settings.ExtendLog) {

	if extendLog.SysLog.Enable == true {
		pri := syslog.LOG_DEBUG
		if extendLog.SysLog.Tag == 1 {
			pri = syslog.LOG_INFO
		}
		hook, err := lSyslog.NewSyslogHook(
			extendLog.SysLog.Network,
			extendLog.SysLog.Address,
			pri,
			extendLog.SysLog.Tag)
		if err == nil {
			log.Hooks.Add(hook)
		}
	}
}
