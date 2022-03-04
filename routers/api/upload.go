package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"upload/pkg/app"
	"upload/pkg/errcode"
	"upload/service"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(file, fileHeader)
	if err != nil {
		log.Println("svc.UploadFile err:", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"name": fileInfo.Name,
	})
}
