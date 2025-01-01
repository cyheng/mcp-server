package logger

import (
	"github.com/sirupsen/logrus"
	"mcp-server/easytcp"
	"runtime/debug"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetLevel(logrus.TraceLevel)

}

func RecoverMiddleware(log *logrus.Logger) easytcp.MiddlewareFunc {
	return func(next easytcp.HandlerFunc) easytcp.HandlerFunc {
		return func(c easytcp.Context) {
			defer func() {
				if r := recover(); r != nil {
					log.WithField("sid", c.Session().ID()).Errorf("PANIC | %+v | %s", r, debug.Stack())
				}
			}()
			next(c)
		}
	}
}

func LogMiddleware(next easytcp.HandlerFunc) easytcp.HandlerFunc {
	return func(c easytcp.Context) {
		req := c.Request()
		logger.Infof("rec <<< id:(%d) size:(%d) data: %s", req.ID(), len(req.Data()), req.Data())
		defer func() {
			resp := c.Response()
			logger.Infof("snd >>> id:(%d) size:(%d) data: %s", resp.ID(), len(resp.Data()), resp.Data())
		}()
		next(c)
	}
}
func SetLogger(newlog *logrus.Logger) {
	logger = newlog
}

func Ins() *logrus.Logger {
	return logger
}
