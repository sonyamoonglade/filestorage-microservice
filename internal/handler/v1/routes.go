package handler

import "github.com/sonyamoonglade/storage-service/internal/handler/v1/middleware"

func (h *FileHandler) Routes() {
	v1 := h.v1.Group("/api/v1")
	srv := v1.Group("/service", middleware.HmacVerificationMiddleware(h.logger))
	{
		srv.POST("/put", h.put)
		srv.GET("/all", h.getAll)
		srv.DELETE("/delete", h.delete)
	}

}
