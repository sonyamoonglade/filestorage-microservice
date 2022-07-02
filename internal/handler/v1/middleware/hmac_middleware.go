package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/headers"
	"github.com/sonyamoonglade/storage-service/pkg/hmac"
	"github.com/sonyamoonglade/storage-service/pkg/util"
	"go.uber.org/zap"
	"os"
)

func HmacVerificationMiddleware(logger *zap.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		secret := os.Getenv("HASH_SECRET")

		var putHeaders headers.PutHeaders

		putHeaders.XFileName = c.GetHeader(headers.XFileName)
		putHeaders.XDestination = c.GetHeader(headers.XDestination)
		putHeaders.XFileExt = c.GetHeader(headers.XFileExt)
		putHeaders.XHmacSignature = c.GetHeader(headers.XHmacSignature)

		logger.Info(fmt.Sprintf("got headers - %v", putHeaders))

		if ok := util.ValidateHeaders(putHeaders); ok != true {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "invalid request headers",
			})
			logger.Error("cannot bind headers")
			return
		}

		xhmac := putHeaders.XHmacSignature
		logger.Info(fmt.Sprintf("got hmac from header - %s", xhmac))

		xheaderSum := []string{putHeaders.XFileExt, putHeaders.XFileName, putHeaders.XDestination}
		xsignature := hmacservice.GenerateHexSignature(xheaderSum)

		compResult := hmacservice.ValidateHMAC(xhmac, xsignature, secret)
		logger.Info(fmt.Sprintf("got comp result - %t", compResult))
		if !compResult {
			c.AbortWithStatusJSON(403, gin.H{
				"message": "invalid hmac",
			})
			logger.Error("got invalid hmac")
			return
		}

		h := headers.Headers{
			headers.XDestination: putHeaders.XDestination,
			headers.XFileName:    putHeaders.XFileName,
			headers.XFileExt:     putHeaders.XFileExt,
		}

		c.Set(headers.XHeaders, h)
		c.Next()
		return
	}

}
