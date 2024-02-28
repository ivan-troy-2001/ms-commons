package logger

import (
	"github.com/pterm/pterm"
	"reflect"
)

type Empty struct{}

var (
	Log *pterm.Logger
)

func init() {
	Log = pterm.DefaultLogger.WithLevel(pterm.LogLevelDebug).WithCaller(false)
	Log.Trace(reflect.TypeOf(Empty{}).PkgPath() + " initialized")
}
