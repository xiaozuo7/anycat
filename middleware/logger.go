package middleware

import (
	"anycat/global/consts"
	"anycat/global/variable"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type logField struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	SpendTime string `json:"spend_time"`
	Hostname  string `json:"hostname"`
	ClientIP  string `json:"client_ip"`
	Uri       string `json:"uri"`
	Method    string `json:"method"`
	Status    int    `json:"status"`
	UserAgent string `json:"user_agent"`
}

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		// Process request
		c.Next()
		// Stop timer
		end := time.Now()
		// Execution time
		spend := end.Sub(start)
		logField := logField{
			StartTime: start.Format(consts.TimeForMate),
			EndTime:   end.Format(consts.TimeForMate),
			SpendTime: spend.String(),
			Hostname:  c.Request.Host,
			ClientIP:  c.ClientIP(),
			Uri:       c.Request.RequestURI,
			Method:    c.Request.Method,
			Status:    c.Writer.Status(),
			UserAgent: c.Request.UserAgent(),
		}
		str, err := json.Marshal(logField)
		if err != nil {
			variable.ZapLog.Error("Request log error", zap.Error(err))
		}
		if len(c.Errors) > 0 {
			variable.ZapLog.Error("Request log errors", zap.String("error", c.Errors.ByType(gin.ErrorTypePrivate).String()))
		}
		if status := c.Writer.Status(); status == 404 {
			variable.ZapLog.Warn("Request log 404", zap.String("error", string(str)))
		} else if status >= 500 {
			variable.ZapLog.Error("Request log >= 500", zap.String("error", string(str)))
		} else {
			variable.ZapLog.Info("Request log", zap.String("info", string(str)))
		}
	}
}
