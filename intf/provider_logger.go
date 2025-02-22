package intf

import (
	"github.com/hdget/common/types"
	"log"
)

type LoggerProvider interface {
	types.Provider
	GetStdLogger() *log.Logger
	Log(keyvals ...interface{}) error
	Trace(msg string, keyvals ...interface{})
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Warn(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	Fatal(msg string, keyvals ...interface{})
	Panic(msg string, keyvals ...interface{})
}
