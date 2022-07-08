package tests

import (
	bytes2 "bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/headers"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/middleware"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMissingHeader(t *testing.T) {
	t.Log("Testing when 1 header is missing")
	r := SetupTestRouter()
	bytes := bytes2.NewBuffer([]byte("i am good file kekw:D"))
	req := httptest.NewRequest(http.MethodPost, "/service/put", bytes)
	req.Header.Add(headers.XDestination, "static/images")
	req.Header.Add(headers.XFileName, "5")
	req.Header.Add(headers.XFileExt, "png")

	// explicitly not adding signature. Expect 400
	//req.Header.Add(headers.XHmacSignature, "png")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	rmap := gin.H{
		"message": "invalid request headers",
	}
	response, _ := json.Marshal(rmap)
	assert.Equal(t, string(response), w.Body.String())

}

func TestInvalidHMACSignature(t *testing.T) {
	t.Log("Testing when 1 header is missing")
	r := SetupTestRouter()
	bytes := bytes2.NewBuffer([]byte("i am good file kekw:D"))
	req := httptest.NewRequest(http.MethodPost, "/service/put", bytes)
	req.Header.Add(headers.XDestination, "static/images")
	req.Header.Add(headers.XFileName, "5")
	req.Header.Add(headers.XFileExt, "png")

	// explicitly adding invalid signature. Expect 403
	req.Header.Add(headers.XHmacSignature, "asdfkljadsfhajsfhaskfhadskjfhsdakfhasfdkjashf")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	rmap := gin.H{
		"message": "invalid hmac",
	}
	response, _ := json.Marshal(rmap)
	assert.Equal(t, string(response), w.Body.String())

}

func SetupTestRouter() *gin.Engine {
	r := gin.New()
	logger, _ := zap.NewProduction()
	srv := r.Group("/service")
	srv.POST("/put", middleware.HmacVerificationMiddleware(logger))

	return r
}
