package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sonyamoonglade/storage-service/internal/app_errors"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/dto"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/headers"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/middleware"
	"github.com/sonyamoonglade/storage-service/internal/service"
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

	var hmap map[string]string

	if err := headers.Decode(xheaders, &hmap); err != nil {
		app_errors.Internal(c)
		h.logger.Error(err.Error())
		return
	}

	putDto.Destination = hmap[headers.XDestination]
	putDto.FilenameWithExt = hmap[headers.XFileName] + "." + hmap[headers.XFileExt]

	ok, err := h.service.Put(c.Request.Context(), c.Request.Body, putDto)
	if err != nil {
		app_errors.Internal(c)
		h.logger.Error(err.Error())
		return
	}

	if !ok {
		app_errors.InternalMsg(c, "operation cannot be completed")
		h.logger.Error(err.Error())
		return
	}

	c.JSON(201, gin.H{
		"ok": ok,
	})
	h.logger.Info("operation executed successfully")
	return
}

func (h *FileHandler) delete(c *gin.Context) {

}

func (h *FileHandler) getAll(c *gin.Context) {

}
