package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnJson(c *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	c.JSON(httpCode, gin.H{
		"status": dataCode,
		"msg":    msg,
		"data":   data,
	})
}

// Success 返回成功 200
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, 200, msg, data)
}

// Fail 返回失败 400
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

// ErrorSystem 服务器错误 500
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, 500, msg, data)
	c.Abort()
}
