package storage

import (
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client
var ctx context.Context

func init() {
	ctx = context.Background()

	var err error
	minioClient, err = minio.New(os.Getenv("STORAGE_ENDPOINT"), &minio.Options{
		Creds: credentials.NewStaticV4(
			os.Getenv("STORAGE_ACCESS_KEY_ID"),
			os.Getenv("STORAGE_SECRET_ACCESS_KEY"),
			"",
		),
		Secure: os.Getenv("ENV") != "dev",
	})

	if err != nil {
		log.Fatalln(err)
	}

	bucketName := os.Getenv("STORAGE_BUCKET_NAME")

	if found, errBucketExists := minioClient.BucketExists(ctx, bucketName); errBucketExists == nil && !found {
		createBucket(bucketName)
	}
}

func createBucket(bucketName string) {
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: os.Getenv("STORAGE_BUCKET_LOCATION")})

	if err != nil {
		log.Println("Error creating bucket:", err)

		return
	}

	log.Printf("Bucket %s created", bucketName)
}

// AddToBucket a given file
func AddToBucket(fileName string, reader io.Reader, objectSize int64, contentType string) error {
	if _, err := minioClient.PutObject(
		ctx,
		os.Getenv("STORAGE_BUCKET_NAME"),
		fileName,
		reader,
		objectSize,
		minio.PutObjectOptions{ContentType: contentType},
	); err != nil {
		return err
	}

	return nil
}

// GetFromBucket a given file and returns a encoded base64 string of it
func GetFromBucket(fileName string) (string, error) {
	pictureReader, err := minioClient.GetObject(
		ctx,
		os.Getenv("STORAGE_BUCKET_NAME"),
		fileName,
		minio.GetObjectOptions{},
	)

	if err != nil {
		return "", err
	}

	bytePicture, err := ioutil.ReadAll(pictureReader)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytePicture), nil
}
