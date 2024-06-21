package middleware

import (
	"anycat/global/variable"
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

type logField struct {
	Uri         string `json:"uri"`
	Lantency    string `json:"lantency"`
	Status      int    `json:"status"`
	Method      string `json:"method"`
	Hostname    string `json:"hostname"`
	ClientIP    string `json:"clientIP"`
	UserAgent   string `json:"userAgent"`
	ContentType string `json:"contentType"`
	ReqBody     string `json:"reqBody"`
	RespBody    string `json:"respBody"`
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqBody, _ := ctx.GetRawData()
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = bodyLogWriter
		start := time.Now()

		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)
		respBody := ""
		if bodyLogWriter.body.Len() > 0 {
			respBody = bodyLogWriter.body.String()
		}
		logField := logField{
			Uri:         ctx.Request.URL.Path,
			Lantency:    latency.String(),
			Status:      ctx.Writer.Status(),
			Method:      ctx.Request.Method,
			Hostname:    ctx.Request.Host,
			ClientIP:    ctx.ClientIP(),
			UserAgent:   ctx.Request.UserAgent(),
			ContentType: ctx.ContentType(),
			ReqBody:     string(reqBody),
			RespBody:    respBody,
		}
		str, err := json.Marshal(logField)
		if err != nil {
			variable.ZapLog.Errorw("Request log json.Marshal failed", zap.Error(err))
			return
		}
		if status := ctx.Writer.Status(); status == 404 {
			variable.ZapLog.Warnw("Request log 404", zap.String("info", string(str)))
		} else if status >= 500 {
			variable.ZapLog.Errorw("Request log >= 500", zap.String("info", string(str)))
		} else {
			variable.ZapLog.Infow("Request log 200", zap.String("info", string(str)))
		}
	}
}
