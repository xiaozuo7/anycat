package v1

import (
	"anycat/model"
	"anycat/service"
	"anycat/util/response"

	"github.com/gin-gonic/gin"
)

func TransHandler(c *gin.Context) {
	var formData model.TransReq

	err := c.ShouldBindJSON(&formData)
	if err != nil {
		response.Fail(c, 400, err.Error(), "")
		return
	}
	res, err := service.Trans(formData)
	if err != nil {
		response.Fail(c, 400, err.Error(), "")
		return
	}
	response.Success(c, "success", res)

}
