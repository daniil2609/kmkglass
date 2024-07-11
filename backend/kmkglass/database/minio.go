package database

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var BucketName = "kmkglass-photo-bucket"

func InitMinio() {
	var err error

	// Инициализация MinIO клиента
	MinioClient, err = minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ROOT_USER"), os.Getenv("MINIO_ROOT_PASSWORD"), ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Убедитесь, что ведро существует
	exists, err := MinioClient.BucketExists(context.Background(), BucketName)
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
		err = MinioClient.MakeBucket(context.Background(), BucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
