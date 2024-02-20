package v1

import (
	"anycat/model"
	"anycat/service"
	"anycat/util/response"

	"github.com/gin-gonic/gin"
)

func Base64Encode(c *gin.Context) {
	var formData model.Base64Encode

	_ = c.ShouldBindJSON(&formData)
	res := service.Base64Encode(formData.Content)
	response.Success(c, "success", res)

}

func Base64Decode(c *gin.Context) {
	var formData model.Base64Decode

	_ = c.ShouldBindJSON(&formData)
	res, err := service.Base64Decode(formData.Content)
	if err != nil {
		response.Fail(c, 400, err.Error(), "")
		return
	}
	response.Success(c, "success", res)

}
