package middleware

import (
	"anycat/global/variable"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func defaultHandleRecovery(c *gin.Context, _ interface{}) {
	c.AbortWithStatus(http.StatusInternalServerError)
}

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					variable.ZapLog.Errorf("[Recovery] panic recovered:\n%s\n%s", string(httpRequest), err)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				variable.ZapLog.Errorf("[Recovery] panic recovered:\n%s\n%s", string(httpRequest), err)
				defaultHandleRecovery(c, err)
			}
		}()
		c.Next()
	}
}
