package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sonyamoonglade/s3-yandex-go/s3yandex"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1"
	"github.com/sonyamoonglade/storage-service/internal/service"
	"go.uber.org/zap"
	"os"
)

const (
	// name of the bucket in s3
	bucket = "zharpizza-bucket"
)

func main() {

	logger, _ := zap.NewProduction()

	v1 := gin.Default()
	if err := godotenv.Load(".env"); err != nil {
		logger.Error("could not load environment file")
		os.Exit(1)
	}
	s3CredProvider := s3yandex.NewEnvCredentialsProvider()
	logger.Info("initialized credential provider")

	client := s3yandex.NewYandexS3Client(s3CredProvider, s3yandex.YandexS3Config{
		Owner:  os.Getenv("OWNER"),
		Bucket: bucket,
		Debug:  false,
	})
	logger.Info("initialized s3 client")

	fileService := service.NewFileService(client)
	fileHandler := handler.NewFileHandler(v1, fileService, logger)
	fileHandler.Routes()
	
	logger.Info("initialized file composition")

	if err := v1.Run(":5001"); err != nil {
		logger.Error(fmt.Sprintf("Cannot listen to :5000. %s", err.Error()))
		os.Exit(1)
	}

}
