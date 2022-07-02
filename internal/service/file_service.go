package service

import (
	"context"
	"github.com/sonyamoonglade/s3-yandex-go/s3yandex"
	"github.com/sonyamoonglade/storage-service/internal/handler/v1/dto"
	"io"
)

type File interface {
	Put(ctx context.Context, body io.ReadCloser, dto dto.PutFileDto) (bool, error)
	Delete() (bool, error)
}

type fileService struct {
	client *s3yandex.YandexS3Client
}

func NewFileService(client *s3yandex.YandexS3Client) *fileService {
	return &fileService{client: client}
}

func (f *fileService) Put(ctx context.Context, body io.ReadCloser, dto dto.PutFileDto) (bool, error) {

	fileBytes, err := io.ReadAll(body)
	if err != nil {
		return false, err
	}

	err = f.client.PutFileWithBytes(ctx, &s3yandex.PutFileWithBytesInput{
		ContentType: s3yandex.ImagePNG,
		FileName:    dto.FilenameWithExt,
		Destination: dto.Destination,
		FileBytes:   &fileBytes,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f *fileService) Delete() (bool, error) {
	panic("implement me")
}
