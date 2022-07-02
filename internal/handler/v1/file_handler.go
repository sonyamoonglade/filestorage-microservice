package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/dto"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/headers"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/middleware"
	"github.com/sonyamoonglade/storage-service/internal/service"
	"github.com/sonyamoonglade/storage-service/pkg/util"
	"go.uber.org/zap"
)

type FileHandler struct {
	v1      *gin.Engine
	service service.File
	logger  *zap.Logger
}

func NewFileHandler(v1 *gin.Engine, service service.File, logger *zap.Logger) *FileHandler {
	return &FileHandler{v1: v1, service: service, logger: logger}
}

func (h *FileHandler) InitRoutes() {

	srv := h.v1.Group("/service", middleware.HmacVerificationMiddleware(h.logger))
	{
		srv.POST("/put", h.put)
		srv.GET("/all", h.getAll)
		srv.DELETE("/delete", h.delete)
	}

}

func (h *FileHandler) put(c *gin.Context) {

	var putDto dto.PutFileDto

	xheaders, _ := c.Get(headers.XHeaders)

	var xheaderMap map[string]string

	if err := util.GetHeaderMap(xheaders, &xheaderMap); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "internal service error",
		})
		h.logger.Error(err.Error())
		return
	}

	putDto.Destination = xheaderMap[headers.XDestination]
	putDto.FilenameWithExt = xheaderMap[headers.XFileName] + "." + xheaderMap[headers.XFileExt]

	ok, err := h.service.Put(c.Request.Context(), c.Request.Body, putDto)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "internal service error",
		})
		h.logger.Error(err.Error())
		return
	}

	if !ok {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "operation cannot be completed",
		})
		h.logger.Error(err.Error())
		return
	}

	c.JSON(200, gin.H{
		"ok": ok,
	})
	h.logger.Info("operation executed successfully")
	return
}

func (h *FileHandler) delete(c *gin.Context) {

}

func (h *FileHandler) getAll(c *gin.Context) {

}
