package routers

import (
	"github.com/gin-gonic/gin"
	"upload/routers/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)

	return r
}
